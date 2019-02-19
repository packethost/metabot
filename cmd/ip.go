package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
  Use:   "ip",
  Short: "Print the packet ip information, subject to qualifiers",
  Long:  `Print the packet ip information, subject to qualifiers`,
  Run: func(cmd *cobra.Command, args []string) {
     // ip has lots of qualifiers
    results := make([]string, 0)
    for _, addr := range data.Network.Addresses {
        results = append(results, fmt.Sprintf("%s/%d", addr.Address, addr.Cidr))
    }
    fmt.Println(results)
    os.Exit(0)
  },
}

