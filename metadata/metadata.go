package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Bonding mode
type Bonding struct {
	Mode uint32
}

// Interface network interface
type Interface struct {
	Name string
	Mac  string
	Bond string
}

// Address network address
type Address struct {
	ID            string
	AddressFamily uint32 `json:"address_family"`
	Netmask       string
	Public        bool
	Cidr          uint32
	Management    bool
	Enabled       bool
	Network       string
	Address       string
	Gateway       string
	Parent        *Address `json:"parent_block"`
}

// Network complete metadata network information
type Network struct {
	Bonding    Bonding
	Interfaces []Interface
	Addresses  []Address
}

// Metadata entire metadata
type Metadata struct {
	ID       string
	Hostname string
	Iqn      string
	Facility string
	Network  Network
}

func getMetadata(u string) ([]byte, error) {
	// get metadata
	var body []byte
	parsed, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("unable to parse url %s: %v", u, err)
	}
	if parsed.Scheme == "file" {
		body, err = ioutil.ReadFile(parsed.Path)
		if err != nil {
			return nil, err
		}
        } else {
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("non-200 return code: %d", resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}
	return body, nil
}

func parseMetadata(b []byte) (*Metadata, error) {
	var m Metadata
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// GetAndParseMetadata retrieve metadata from a specific URL or Packet's standard
func GetAndParseMetadata(u string) (*Metadata, error) {
	if u == "" {
		return nil, fmt.Errorf("cannot parse empty metadata URL")
	}
	b, err := getMetadata(u)
	if err != nil {
		return nil, err
	}
	parsed, err := parseMetadata(b)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
