package mimik_core

import (
	"fmt"
	"mimik/internal/data"
	"mimik/internal/utils"
	"path/filepath"
	"sync"
)

func CreateData(flags data.LaunchFlags) {
	fmt.Println("Creating Namespace files")

	var selectedData data.SelectedContent = utils.ExtractRequiredLocations(flags)

	if selectedData.Validated == true {
		var wg sync.WaitGroup
		var total_records uint64 = 0
		var lock sync.Mutex
		for clusterGroup, clusterGroupData := range selectedData.ContinentCountryLocationMap {
			for clusterName, clusterData := range clusterGroupData {
				var fileName string = fmt.Sprintf("%s.txt", clusterName)
				var dirPath string = filepath.Join(utils.DefaultDirName, fileName)
				var filePath string = filepath.Join(flags.DataPath, dirPath)
				fmt.Println(filePath)
				wg.Add(1)
				go func(filePath string,
					clusterGroup string,
					clusterName string,
					clusterData map[string]map[string]data.DeploymentsMetadata) {
					num_records, err := utils.CreateClusterDataFiles(filePath,
						clusterGroup,
						clusterName,
						clusterData,
						flags.Days,
						&wg)

					if err != nil {
						fmt.Println("error")
					} else {
						lock.Lock()
						total_records += num_records
						lock.Unlock()
					}
				}(filePath, clusterGroup, clusterName, clusterData)
			}
		}
		wg.Wait()
		fmt.Printf("Number of records created : %d", total_records)
	}
}
