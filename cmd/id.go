package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(idCmd)
}

var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Print the packet host id",
	Long:  `Print the packet host id`,
	Run: func(cmd *cobra.Command, args []string) {
		// id has no qualifiers
		reportAndExit(data.ID)
	},
}
