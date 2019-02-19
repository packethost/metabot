package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "github.com/packethost/metabot/metadata"
)

func init() {
  rootCmd.AddCommand(hostnameCmd)
}

var hostnameCmd = &cobra.Command{
  Use:   "hostname",
  Short: "Print the packet hostname",
  Long:  `Print the packet hostname`,
  Run: func(cmd *cobra.Command, args []string) {
     // hostname has no qualifiers
  },
}

