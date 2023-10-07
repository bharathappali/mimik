package mimik_core

import (
	"fmt"
	"mimik/internal/data"
)

func CreateNamespaceFiles(flags data.LaunchFlags) {
	fmt.Println("Creating Namespace files")

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
			break
		}
	}

	fmt.Println(selectedData)
}
