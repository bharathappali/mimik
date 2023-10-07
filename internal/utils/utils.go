package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"mimik/internal/data"
)

// ValidateFlags validates the clusters, runMimik, and days flags.
func ValidateFlags(cmd *cobra.Command) data.LaunchFlags {
	validator := data.LaunchFlags{}

	clusters, _ := cmd.Flags().GetInt("clusters")
	runMimik, _ := cmd.Flags().GetBool("run-mimik")
	days, _ := cmd.Flags().GetInt("days")

	// Setting it to true by default
	validator.Validated = true

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
		ContinentCountryLocationMap: make(map[string]map[string][]string),
		Validated:                   false, // Initially not validated
	}

	for continent, countries_map := range data.ContinentCountryLocationMap {
		for country, locations_list := range countries_map {
			locationsCovered = locationsCovered + 1

			if _, ok := selectedData.ContinentCountryLocationMap[continent]; !ok {
				selectedData.ContinentCountryLocationMap[continent] = make(map[string][]string)
			}

			var continent_map = selectedData.ContinentCountryLocationMap[continent]
			if _, ok := continent_map[country]; !ok {
				continent_map[country] = locations_list
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
