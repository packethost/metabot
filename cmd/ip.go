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
		// ip has lots of qualifiers
		qualifiers := map[string]bool{
			"4":       false,
			"6":       false,
			"private": false,
			"public":  false,
			"address": false,
			"netmask": false,
			"parent":  false,
		}
		for _, a := range args {
			qualifiers[a] = true
		}

		// first use the filter qualifiers to filter out which data set we want
		addresses := make([]*metadata.Address, 0)
		for _, addr := range data.Network.Addresses {
			// did we exclude it for a particular reason?
			if addr.Public && qualifiers["private"] {
				continue
			}
			if !addr.Public && qualifiers["public"] {
				continue
			}
			if addr.AddressFamily == 6 && qualifiers["4"] {
				continue
			}
			if addr.AddressFamily == 4 && qualifiers["6"] {
				continue
			}
			saver := &addr
			if qualifiers["parent"] {
				saver = addr.Parent
			}
			addresses = append(addresses, saver)
		}
		results := make([]string, 0)
		for _, addr := range addresses {
			addrValue := addr.Address
			if qualifiers["network"] {
				addrValue = fmt.Sprintf("%s/%d", addr.Network, addr.Cidr)
			}
			saver := fmt.Sprintf("%s/%d", addrValue, addr.Cidr)
			switch {
			case qualifiers["netmask"]:
				saver = fmt.Sprintf("%s/%s", addrValue, addr.Netmask)
			case qualifiers["address"]:
				saver = fmt.Sprintf("%s", addrValue)
			}
			results = append(results, saver)
		}
		reportAndExit(strings.Join(results, " "))
	},
}
