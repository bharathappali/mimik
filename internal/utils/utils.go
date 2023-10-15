package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/cobra"
	"math/big"
	"mimik/internal/data"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
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
	} else {
		var datadir string = filepath.Join(data_path, DefaultDirName)

		// Create the mimik data directory if it doesn't exist
		if _, err := os.Stat(datadir); os.IsNotExist(err) {
			err := os.MkdirAll(datadir, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				os.Exit(1)
			}
		}
		err := deleteTxtFiles(datadir)
		if err != nil {
			fmt.Println("Error Deleting previous files")
		}
		validator.DataPath = data_path
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
		ContinentCountryLocationMap: make(map[string]map[string]map[string]map[string]data.DeploymentsMetadata),
		Validated:                   false, // Initially not validated
	}

	for continent, countries_map := range data.ContinentCountryLocationMap {
		for country, locations_list := range countries_map {
			var deploymentsNamesMap map[string]bool = make(map[string]bool)
			locationsCovered = locationsCovered + 1

			if _, ok := selectedData.ContinentCountryLocationMap[continent]; !ok {
				selectedData.ContinentCountryLocationMap[continent] = make(map[string]map[string]map[string]data.DeploymentsMetadata)
			}

			var continent_map = selectedData.ContinentCountryLocationMap[continent]
			if _, ok := continent_map[country]; !ok {
				continent_map[country] = make(map[string]map[string]data.DeploymentsMetadata)
			}
			var country_map = continent_map[country]

			for location_index := range locations_list {
				if _, ok := country_map[locations_list[location_index]]; !ok {
					country_map[locations_list[location_index]] = make(map[string]data.DeploymentsMetadata)
				}
				deployments, err := getRandomInRange_UINT8(MinDeployments, MaxDeployments)
				if err != nil {
					deployments = MinDeployments
				}
				var metadata_hl_map = make(map[string]map[string]uint8, deployments)

				for deployment := uint8(0); deployment < deployments; deployment++ {
					var deploymentName string = generateRandomCombination()
					// Find and fill the deployment name
					for {
						if _, found := deploymentsNamesMap[deploymentName]; found {
							deploymentName = generateRandomCombination()
						} else {
							deploymentsNamesMap[deploymentName] = true
							break // Break as we fount a unique deployment name
						}
					}
					var metadata_map = make(map[string]uint8)
					containers, err := getRandomInRange_UINT8(MinContainers, MaxContainers)
					if err != nil {
						containers = MinContainers
					}
					replicas, err := getRandomInRange_UINT8(MinReplicas, MaxReplicas)
					if err != nil {
						replicas = MinReplicas
					}
					mode_index, mode_err := getRandomInRange_UINT8(0, uint8(len(ResourceUsageMode)-1))
					if mode_err != nil {
						mode_index = 0
					}
					metadata_map = map[string]uint8{
						"containers":    containers,
						"replicas":      replicas,
						"resource_mode": mode_index,
					}
					metadata_hl_map[deploymentName] = metadata_map
				}

				country_map[locations_list[location_index]] = map[string]data.DeploymentsMetadata{
					"deployments": {
						NumDeployments: uint8(deployments),
						Metadata:       metadata_hl_map,
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

	return selectedData
}

func CreateClusterDataFiles(
	file_name string,
	cluster_group string,
	cluster_name string,
	namespace_map map[string]map[string]data.DeploymentsMetadata,
	days int,
	wg *sync.WaitGroup) (uint64, error) {

	// Create or open the file for writing
	file, err := os.Create(file_name)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s", file_name)
			os.Exit(1)
		}
	}(file)

	container_counters := make(map[string]float64)

	var records_counter uint64 = 0

	var started bool = true

	// Get current UTC time
	currentTime := time.Now().UTC()
	// Calculate the start time by subtracting numDays from the current time
	startTime := currentTime.AddDate(0, 0, -days)

	// Iterate over the time range for every 60 seconds
	for t := startTime; t.Before(currentTime); t = t.Add(60 * 60 * time.Second) {
		// Calculate the milliseconds for the current time
		millis := t.UnixNano() / int64(time.Millisecond)

		// Iterate over the keys of the namespace map
		for namespace, namespace_map := range namespace_map {
			deploymentsMetadata := namespace_map["deployments"]
			// Loop over number of deployments and each
			for deploymentName, deploymentMap := range deploymentsMetadata.Metadata {
				var replicas uint8 = deploymentMap["replicas"]
				var containers uint8 = deploymentMap["containers"]
				var resource_mode_index uint8 = deploymentMap["resource_mode"]
				if resource_mode_index >= uint8(len(ResourceUsageMode)) {
					resource_mode_index = 0
				}
				for podNum := uint8(0); podNum < replicas; podNum++ {
					var podName string = fmt.Sprintf("%s_pod_%d", deploymentName, podNum)
					for containerNum := uint8(0); containerNum < containers; containerNum++ {
						var containerName string = fmt.Sprintf("%s_container_%d", deploymentName, containerNum)

						var counter_key string = fmt.Sprintf("%s_pod_%d_container_%d", deploymentName, podNum, containerNum)
						for query, query_type := range Queries {
							line := ""
							if started {
								line = ""
							} else {
								started = false
								line += "\n"
							}
							line += fmt.Sprintf("# HELP %s Resource Query", query)
							line += fmt.Sprintf("\n# TYPE %s %s", query, query_type)

							var value float64 = 0.0
							var uint_val uint64 = 0
							var is_float bool = true

							switch query {
							case Query_Memory_Usage:
								is_float = false
								uint_val = generateMemoryUsage(resource_mode_index)
							case Query_CPU_Usage:
								cpu_usage, _ := generateCPUUsage(resource_mode_index)
								var cpu_seconds float64 = 60 * 60 * cpu_usage
								current_count, key_exist := container_counters[counter_key]
								if key_exist {
									container_counters[counter_key] = current_count + cpu_seconds
								} else {
									container_counters[counter_key] = 100.00 + cpu_seconds
								}
								value = container_counters[counter_key]
							case Query_Memory_Request:
								value = getMemoryRequest(resource_mode_index)
							case Query_Memory_limit:
								value = getMemoryLimit(resource_mode_index)
							case Query_CPU_Limit:
								value = getCPULimit(resource_mode_index)
							case Query_CPU_Request:
								value = getCPURequest(resource_mode_index)
							}

							if is_float {
								line += fmt.Sprintf("\n%s{cluster_group=\"%s\",cluster_name=\"%s\",namespace=\"%s\",deployment_name=\"%s\",pod_name=\"%s\",container_name=\"%s\"} %f %d", query, cluster_group, cluster_name, namespace, deploymentName, podName, containerName, value, millis)
							} else {
								line += fmt.Sprintf("\n%s{cluster_group=\"%s\",cluster_name=\"%s\",namespace=\"%s\",deployment_name=\"%s\",pod_name=\"%s\",container_name=\"%s\"} %d %d", query, cluster_group, cluster_name, namespace, deploymentName, podName, containerName, uint_val, millis)
							}
							records_counter += 1
							_, err := fmt.Fprintln(file, line)
							if err != nil {
								fmt.Printf("Error occured")
								return records_counter, err
							}
						}
					}
				}
			}
		}
	}
	_, write_err := fmt.Fprintln(file, "# EOF")
	if write_err != nil {
		fmt.Printf("Error occured")
		return records_counter, err
	}
	wg.Done()
	return records_counter, nil
}

func deleteTxtFiles(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			err := os.Remove(path) // Delete .txt files
			if err != nil {
				return err
			}
			fmt.Printf("Deleted: %s\n", path)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Function to generate a random combination
func generateRandomCombination() string {

	// Generate random indices for each array
	nameIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(Names))))
	placeIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(Places))))
	animalIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(Animals))))
	thingIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(Things))))

	// Create a combination using the randomly generated indices
	combination := fmt.Sprintf("%s_%s_%s_%s", Names[nameIndex.Int64()], Places[placeIndex.Int64()], Animals[animalIndex.Int64()], Things[thingIndex.Int64()])

	return combination
}

func generateCPUUsage(usage_mode uint8) (float64, float64) {
	var cpu_usage float64 = LowModeMinCPU
	var cpu_throttle float64 = 0.0
	var cpu_limit float64 = LowModeCPULimit

	if usage_mode == 0 {
		cpu_limit = LowModeCPULimit

		cpu_usage_val, cpu_usage_err := getRandomInRange_FLOAT64(LowModeMinCPU, LowModeMaxCPU)
		if cpu_usage_err != nil {
			cpu_usage = LowModeMinCPU
		} else {
			cpu_usage = cpu_usage_val
		}

		if cpu_usage > cpu_limit {
			cpu_throttle_val, cpu_throttle_err := getRandomInRange_FLOAT64(LowModeMinCPU, LowModeMaxCPU)
			if cpu_throttle_err != nil {
				cpu_throttle = LowModeMinCPU
			} else {
				cpu_throttle = cpu_throttle_val
			}
		}
	} else if usage_mode == 1 {
		cpu_limit = AvgModeCPULimit

		cpu_usage_val, cpu_usage_err := getRandomInRange_FLOAT64(AvgModeMinCPU, AvgModeMaxCPU)
		if cpu_usage_err != nil {
			cpu_usage = AvgModeMinCPU
		} else {
			cpu_usage = cpu_usage_val
		}

		if cpu_usage > cpu_limit {
			cpu_throttle_val, cpu_throttle_err := getRandomInRange_FLOAT64(AvgModeMinCPU, AvgModeMaxCPU)
			if cpu_throttle_err != nil {
				cpu_throttle = AvgModeMinCPU
			} else {
				cpu_throttle = cpu_throttle_val
			}
		}
	} else if usage_mode == 2 {
		cpu_limit = HighModeCPULimit

		cpu_usage_val, cpu_usage_err := getRandomInRange_FLOAT64(HighModeMinCPU, HighModeMaxCPU)
		if cpu_usage_err != nil {
			cpu_usage = HighModeMinCPU
		} else {
			cpu_usage = cpu_usage_val
		}

		if cpu_usage > cpu_limit {
			cpu_throttle_val, cpu_throttle_err := getRandomInRange_FLOAT64(HighModeMinCPU, HighModeMaxCPU)
			if cpu_throttle_err != nil {
				cpu_throttle = HighModeMinCPU
			} else {
				cpu_throttle = cpu_throttle_val
			}
		}
	} else {
		cpu_limit = VeryHighModeCPULimit

		cpu_usage_val, cpu_usage_err := getRandomInRange_FLOAT64(VeryHighModeMinCPU, VeryHighModeMaxCPU)
		if cpu_usage_err != nil {
			cpu_usage = VeryHighModeMinCPU
		} else {
			cpu_usage = cpu_usage_val
		}

		if cpu_usage > cpu_limit {
			cpu_throttle_val, cpu_throttle_err := getRandomInRange_FLOAT64(VeryHighModeMinCPU, VeryHighModeMaxCPU)
			if cpu_throttle_err != nil {
				cpu_throttle = VeryHighModeMinCPU
			} else {
				cpu_throttle = cpu_throttle_val
			}
		}
	}
	return cpu_usage, cpu_throttle
}

func generateMemoryUsage(usage_mode uint8) uint64 {
	var mem_usage uint64 = LowModeMinMemory

	switch usage_mode {
	case 0:
		mem_usage_val, err := getRandomInRange_UINT64(LowModeMinMemory, LowModemaxMemory)
		if err != nil {
			mem_usage = LowModeMinMemory
		} else {
			mem_usage = mem_usage_val
		}
	case 1:
		mem_usage_val, err := getRandomInRange_UINT64(AvgModeMinMemory, AvgModemaxMemory)
		if err != nil {
			mem_usage = AvgModeMinMemory
		} else {
			mem_usage = mem_usage_val
		}
	case 2:
		mem_usage_val, err := getRandomInRange_UINT64(HighModeMinMemory, HighModemaxMemory)
		if err != nil {
			mem_usage = HighModeMinMemory
		} else {
			mem_usage = mem_usage_val
		}
	default:
		mem_usage_val, err := getRandomInRange_UINT64(VeryHighModeMinMemory, VeryHighModemaxMemory)
		if err != nil {
			mem_usage = VeryHighModeMinMemory
		} else {
			mem_usage = mem_usage_val
		}
	}
	return mem_usage
}

func getCPURequest(usage_mode uint8) float64 {
	switch usage_mode {
	case 0:
		return LowModeCPURequest
	case 1:
		return AvgModeCPURequest
	case 2:
		return HighModeCPURequest
	default:
		return VeryHighModeCPURequest
	}
}

func getCPULimit(usage_mode uint8) float64 {
	switch usage_mode {
	case 0:
		return LowModeCPULimit
	case 1:
		return AvgModeCPULimit
	case 2:
		return HighModeCPULimit
	default:
		return VeryHighModeCPULimit
	}
}

func getMemoryRequest(usage_mode uint8) float64 {
	switch usage_mode {
	case 0:
		return float64(LowModeMemoryRequest)
	case 1:
		return float64(AvgModeMemoryRequest)
	case 2:
		return float64(HighModeMemoryRequest)
	default:
		return float64(VeryHighModeMemoryRequest)
	}
}

func getMemoryLimit(usage_mode uint8) float64 {
	switch usage_mode {
	case 0:
		return float64(LowModeMemoryLimit)
	case 1:
		return float64(AvgModeMemoryLimit)
	case 2:
		return float64(HighModeMemoryLimit)
	default:
		return float64(VeryHighModeMemoryLimit)
	}
}
