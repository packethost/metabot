package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
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
    fmt.Println(data.Hostname)
    os.Exit(0)
  },
}

