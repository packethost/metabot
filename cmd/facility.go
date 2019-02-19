package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(facilityCmd)
}

var facilityCmd = &cobra.Command{
  Use:   "facility",
  Short: "Print the packet facility id",
  Long:  `Print the packet facility id`,
  Run: func(cmd *cobra.Command, args []string) {
     // facility has no qualifiers
    fmt.Println(data.Facility)
    os.Exit(0)
  },
}

