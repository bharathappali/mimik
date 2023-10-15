package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mimik/internal/mimik_core"
	"mimik/internal/utils"
	"os"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch Mimik",
	Run: func(cmd *cobra.Command, args []string) {
		flags := utils.ValidateFlags(cmd)

		if !flags.Validated {
			fmt.Println("[ERR] Flags validation failed.")
			return
		}

		mimik_core.CreateData(flags)
	},
}

func init() {
	// Get the current working directory, to set the data path
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Log the current working directory
	fmt.Println("Current Working Directory:", cwd)

	// Load flags
	launchCmd.Flags().Int("clusters", utils.DefaultClusters, "Number of clusters")
	launchCmd.Flags().Bool("faux", utils.DefaultFaux, "Run background job for providing metrics")
	launchCmd.Flags().Int("days", utils.DefaultDays, "Number of days")
	launchCmd.Flags().String("data-path", cwd, "Data path to create the metric files")

	rootCmd.AddCommand(launchCmd)
}
