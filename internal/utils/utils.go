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
	if clusters >= 2 && clusters < 199 {
		validator.Clusters = uint16(clusters)
	} else {
		fmt.Printf("%sInvalid clusters: %d | Valid values: Min Clusters - %d, Max Clusters - %d\n", LogErrorPrefix, clusters, MinClusters, MaxClusters)
		validator.Validated = false
	}

	// Validate runMimik flag
	validator.RunMimik = runMimik

	// Validate days flag
	if days > 14 && days < 366 {
		validator.Days = days
	} else {
		fmt.Printf("%sInvalid days: %d | Valid values: Min Days - %d, Max Days - %d\n", LogErrorPrefix, days, MinDays, MaxDays)
		validator.Validated = false
	}

	return validator
}
