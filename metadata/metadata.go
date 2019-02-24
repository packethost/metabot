package metadata

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
