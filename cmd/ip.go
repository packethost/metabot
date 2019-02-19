package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "github.com/packethost/metabot/metadata"
)

func init() {
  rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
  Use:   "facility",
  Short: "Print the packet ip information, subject to qualifiers",
  Long:  `Print the packet ip information, subject to qualifiers`,
  Run: func(cmd *cobra.Command, args []string) {
     // ip has lots of qualifiers
  },
}

