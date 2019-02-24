package internal

import (
	"github.com/packethost/metabot/metadata"
	"io/ioutil"
	"path"
	"runtime"
)

const (
	testFilePath = "./testdata/metadata.json"
)

func GetValidMeta() *metadata.Metadata {
	return &metadata.Metadata{
		ID:       "b642678f-1d6e-45a2-aed1-bd0a63135fe5",
		Hostname: "kubetest",
		Iqn:      "iqn.2019-02.net.packet:device.b642678f",
		Facility: "ewr1",
		Network: metadata.Network{
			Bonding: metadata.Bonding{Mode: 0x4},
			Interfaces: []metadata.Interface{
				metadata.Interface{
					Name: "eth0",
					Mac:  "ec:0d:9a:c0:02:4c",
					Bond: "bond0",
				},
				metadata.Interface{
					Name: "eth1",
					Mac:  "ec:0d:9a:c0:02:4d",
					Bond: "bond0",
				},
			},
			Addresses: []metadata.Address{
				metadata.Address{
					ID:            "362c8eb4-c156-499c-addc-6c1167253be8",
					AddressFamily: 4,
					Netmask:       "255.255.255.254",
					Public:        true,
					Cidr:          31,
					Management:    true,
					Enabled:       true,
					Network:       "147.75.67.72",
					Address:       "147.75.67.73",
					Gateway:       "147.75.67.72",
					Parent: &metadata.Address{
						Network: "147.75.67.72",
						Netmask: "255.255.255.254",
						Cidr:    31,
					},
				},
				metadata.Address{
					ID:            "2280c6bd-38b8-4f9f-bf7a-de01ad4e010c",
					AddressFamily: 6,
					Netmask:       "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffe",
					Public:        true,
					Cidr:          127,
					Management:    true,
					Enabled:       true,
					Network:       "2604:1380:1:9200::",
					Address:       "2604:1380:1:9200::1",
					Gateway:       "2604:1380:1:9200::",
					Parent: &metadata.Address{
						Network: "2604:1380:0001:9200:0000:0000:0000:0000",
						Netmask: "ffff:ffff:ffff:ff00:0000:0000:0000:0000",
						Cidr:    56,
					},
				},
				metadata.Address{
					ID:            "83a52dd8-86bc-4405-97c0-f2561644ada6",
					AddressFamily: 4,
					Netmask:       "255.255.255.254",
					Public:        false,
					Cidr:          31,
					Management:    true,
					Enabled:       true,
					Network:       "10.99.185.0",
					Address:       "10.99.185.1",
					Gateway:       "10.99.185.0",
					Parent: &metadata.Address{
						Network: "10.99.185.0",
						Netmask: "255.255.255.128",
						Cidr:    25,
					},
				},
			},
		},
	}
}

func GetTestData() (string, []byte, error) {
	jsonFile := path.Join(path.Dir(GetMyCaller()), testFilePath)
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return "", nil, err
	}
	return jsonFile, jsonData, nil
}

func GetMyCaller() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}
