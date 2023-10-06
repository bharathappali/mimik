package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mimik/internal/mimik"
	"mimik/internal/utils"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch Mimik",
	Run: func(cmd *cobra.Command, args []string) {
		flags := utils.ValidateFlags(cmd)

		if !flags.Validated {
			fmt.Println("Flags validation failed.")
			return
		}

		mimik.CreateNamespaceFiles(flags)
	},
}

func init() {
	launchCmd.Flags().Int("clusters", utils.DefaultClusters, "Number of clusters")
	launchCmd.Flags().Bool("faux", utils.DefaultFaux, "Run background job for providing metrics")
	launchCmd.Flags().Int("days", utils.DefaultDays, "Number of days")

	rootCmd.AddCommand(launchCmd)
}
