package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "github.com/packethost/metabot/metadata"
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
  },
}

