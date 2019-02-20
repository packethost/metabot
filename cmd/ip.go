package cmd

import (
	"fmt"
	"github.com/packethost/metabot/metadata"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Print the packet ip information, subject to qualifiers",
	Long:  `Print the packet ip information, subject to qualifiers`,
	Run: func(cmd *cobra.Command, args []string) {
		reportAndExit(ipReport(args))
	},
}

func ipReport(qualifiers []string) string {
	// ip has lots of qualifiers
	qualifierActive := map[string]bool{
		"4":       false,
		"6":       false,
		"private": false,
		"public":  false,
		"address": false,
		"netmask": false,
		"parent":  false,
	}
	for _, a := range qualifiers {
		qualifierActive[a] = true
	}
	// first use the filter qualifiers to filter out which data set we want
	addresses := make([]metadata.Address, 0)
	for _, addr := range data.Network.Addresses {
		// did we exclude it for a particular reason?
		if addr.Public && qualifierActive["private"] {
			continue
		}
		if !addr.Public && qualifierActive["public"] {
			continue
		}
		if addr.AddressFamily == 6 && qualifierActive["4"] {
			continue
		}
		if addr.AddressFamily == 4 && qualifierActive["6"] {
			continue
		}
		saver := addr
		if qualifierActive["parent"] {
			saver = *addr.Parent
		}
		addresses = append(addresses, saver)
	}
	results := make([]string, 0)
	for _, addr := range addresses {
		addrValue := addr.Address
		if qualifierActive["network"] {
			addrValue = addr.Network
		}
		saver := fmt.Sprintf("%s/%d", addrValue, addr.Cidr)
		switch {
		case qualifierActive["netmask"]:
			saver = fmt.Sprintf("%s/%s", addrValue, addr.Netmask)
		case qualifierActive["address"]:
			saver = fmt.Sprintf("%s", addrValue)
		}
		results = append(results, saver)
	}
	return strings.Join(results, " ")
}
