package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/cobra"
	"math/big"
	"mimik/internal/data"
	"os"
	"strings"
)

// ValidateFlags validates the clusters, runMimik, and days flags.
func ValidateFlags(cmd *cobra.Command) data.LaunchFlags {
	validator := data.LaunchFlags{}

	clusters, _ := cmd.Flags().GetInt("clusters")
	runMimik, _ := cmd.Flags().GetBool("faux")
	days, _ := cmd.Flags().GetInt("days")
	data_path, _ := cmd.Flags().GetString("data-path")

	// trim the data path
	data_path = strings.Trim(data_path, " ")

	// Setting it to true by default
	validator.Validated = true

	// Check the path
	_, path_err := os.Stat(data_path)

	if path_err != nil {
		if os.IsNotExist(path_err) {
			fmt.Printf("[ERR] Provided data path \" %s \" doesnot exist. Exiting\n", data_path)
		} else {
			fmt.Printf("[ERR] Error checking the path \" %s \"", data_path)
		}
		validator.Validated = false
	}

	// Validate clusters flag
	if clusters >= MinClusters && clusters <= MaxClusters {
		validator.Clusters = uint16(clusters)
	} else {
		fmt.Printf("%sInvalid clusters: %d | Valid values: Min Clusters - %d, Max Clusters - %d\n", LogErrorPrefix, clusters, MinClusters, MaxClusters)
		validator.Validated = false
	}

	// Validate runMimik flag
	validator.RunMimik = runMimik

	// Validate days flag
	if days >= MinDays && days <= MaxDays {
		validator.Days = days
	} else {
		fmt.Printf("%sInvalid days: %d | Valid values: Min Days - %d, Max Days - %d\n", LogErrorPrefix, days, MinDays, MaxDays)
		validator.Validated = false
	}

	return validator
}

func ExtractRequiredLocations(flags data.LaunchFlags) data.SelectedContent {
	var limitReached bool = false
	var locationsCovered uint16 = 0

	var selectedData = data.SelectedContent{
		ContinentCountryLocationMap: make(map[string]map[string]map[string]map[string]data.DeploymentMetadata),
		Validated:                   false, // Initially not validated
	}

	for continent, countries_map := range data.ContinentCountryLocationMap {
		for country, locations_list := range countries_map {
			locationsCovered = locationsCovered + 1

			if _, ok := selectedData.ContinentCountryLocationMap[continent]; !ok {
				selectedData.ContinentCountryLocationMap[continent] = make(map[string]map[string]map[string]data.DeploymentMetadata)
			}

			var continent_map = selectedData.ContinentCountryLocationMap[continent]
			if _, ok := continent_map[country]; !ok {
				continent_map[country] = make(map[string]map[string]data.DeploymentMetadata)
			}
			var country_map = continent_map[country]

			for location_index := range locations_list {
				if _, ok := country_map[locations_list[location_index]]; !ok {
					country_map[locations_list[location_index]] = make(map[string]data.DeploymentMetadata)
				}
				deployments, err := getRandomInRange_UINT8(MinDeployments, MaxDeployments)
				if err != nil {
					deployments = MinDeployments
				}
				var metadata_arr = make([]map[string]uint8, deployments)

				for deployment := uint8(0); deployment < deployments; deployment++ {
					var metadata_map = make(map[string]uint8)
					containers, err := getRandomInRange_UINT8(MinContainers, MaxContainers)
					if err != nil {
						containers = MinContainers
					}
					replicas, err := getRandomInRange_UINT8(MinReplicas, MaxReplicas)
					if err != nil {
						replicas = MinReplicas
					}
					metadata_map = map[string]uint8{
						"containers": containers,
						"replicas":   replicas,
					}
					metadata_arr[deployment] = metadata_map
				}
				country_map[locations_list[location_index]] = map[string]data.DeploymentMetadata{
					"deployments": {
						NumDeployments: uint8(deployments),
						Metadata:       metadata_arr,
					},
				}
			}
			if locationsCovered >= flags.Clusters {
				limitReached = true
				break
			}
		}
		if limitReached {
			selectedData.Validated = true
			break
		}
	}

	fmt.Println(selectedData)

	return selectedData
}

func getRandomInRange_UINT8(min uint8, max uint8) (uint8, error) {
	if min == max {
		return min, nil
	}

	if min > max {
		return 0, fmt.Errorf("Min: %d cannot be greater than Max: %d\n", min, max)
	}

	// Create range
	rng := big.NewInt(int64(max - min + 1))

	// Generate a random number within the specified range
	n, err := rand.Int(rand.Reader, rng)
	if err != nil {
		return 0, err
	}

	// Add min to the generated random number to get a number in the desired range
	return uint8(n.Uint64()) + min, nil
}
