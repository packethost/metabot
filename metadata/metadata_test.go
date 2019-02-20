package metadata

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)
const (
	testFilePath = "../cmd/testdata/metadata.json"
)

var validMeta = &Metadata{
	ID:"b642678f-1d6e-45a2-aed1-bd0a63135fe5",
	Hostname:"kubetest",
	Iqn:"iqn.2019-02.net.packet:device.b642678f",
	Facility:"ewr1",
	Network: Network{
		Bonding: Bonding{Mode:0x4},
		Interfaces:[]Interface{
			Interface{
				Name:"eth0",
				Mac:"ec:0d:9a:c0:02:4c",
				Bond:"bond0",
			},
			Interface{
				Name:"eth1",
				Mac:"ec:0d:9a:c0:02:4d",
				Bond:"bond0",
			},
		},
		Addresses:[]Address{
			Address{
				ID:"362c8eb4-c156-499c-addc-6c1167253be8",
				AddressFamily:4,
				Netmask:"255.255.255.254",
				Public:true,
				Cidr:31,
				Management:true,
				Enabled:true,
				Network:"147.75.67.72",
				Address:"147.75.67.73",
				Gateway:"147.75.67.72",
				Parent: &Address{
					Network:"147.75.67.72",
					Netmask:"255.255.255.254",
					Cidr: 31,
				},
			},
			Address{
				ID:"2280c6bd-38b8-4f9f-bf7a-de01ad4e010c",
				AddressFamily:6,
				Netmask:"ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffe",
				Public:true,
				Cidr:127,
				Management:true,
				Enabled:true,
				Network:"2604:1380:1:9200::",
				Address:"2604:1380:1:9200::1",
				Gateway:"2604:1380:1:9200::",
				Parent: &Address{
					Network:"2604:1380:0001:9200:0000:0000:0000:0000",
					Netmask:"ffff:ffff:ffff:ff00:0000:0000:0000:0000",
					Cidr: 56,
				},
			},
			Address{
				ID:"83a52dd8-86bc-4405-97c0-f2561644ada6",
				AddressFamily:4,
				Netmask:"255.255.255.254",
				Public:false,
				Cidr:31,
				Management:true,
				Enabled:true,
				Network:"10.99.185.0",
				Address:"10.99.185.1",
				Gateway:"10.99.185.0",
				Parent: &Address{
					Network:"10.99.185.0",
					Netmask:"255.255.255.128",
					Cidr: 25,
				},
			},
		},
	},
}

func getTestData() (string, []byte, error) {
	jsonFile, err := filepath.Abs(testFilePath)
	if err != nil {
		return "", nil, err
	}
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return "", nil, err
	}
	return jsonFile, jsonData, nil
}

func TestGetMetadata(t *testing.T) {
	jsonFile, jsonData, err := getTestData()
	if err != nil {
		t.Fatalf("Unable to read test data file: %v", err)
	}
	tests := []struct{
		u string
		b []byte
		err error
	}{
		{"/foo", nil, fmt.Errorf("Get /foo: unsupported protocol scheme")},
		{"file:///no_such_file_foo", nil, fmt.Errorf("open /no_such_file_foo: no such file")},
		{"https://metadata.packet.net/metadata", nil, fmt.Errorf("non-200 return code: 422")},
		{fmt.Sprintf("file://%s", jsonFile), jsonData, nil},
	}
	for i, tt := range tests {
		b, err := getMetadata(tt.u)
		switch {
		case (err == nil && tt.err != nil) || (err != nil && tt.err == nil) || (err != nil && tt.err != nil && !strings.HasPrefix(err.Error(), tt.err.Error())):
			t.Errorf("%d: mismatched error, actual then expected", i)
			t.Logf("%v", err)
			t.Logf("%v", tt.err)
		case bytes.Compare(b, tt.b) != 0:
			t.Errorf("%d: mismatched bytes returned, actual then expected", i)
			t.Logf("%s", string(b))
			t.Logf("%s", string(tt.b))
		}

	}
}

func TestParseMetadata(t *testing.T) {
	_, jsonData, err := getTestData()
	if err != nil {
		t.Fatalf("Unable to read test data file: %v", err)
	}

	tests := []struct{
		b []byte
		m *Metadata
		err error
	}{
		{[]byte(`"bad json"`), nil, fmt.Errorf("json: cannot unmarshal string into Go value of type metadata.Metadata")},
		{[]byte(`{"a":1}`), &Metadata{}, nil},
		{jsonData, validMeta, nil},
	}
	for i, tt := range tests {
		m, err := parseMetadata(tt.b)
		switch {
		case (err == nil && tt.err != nil) || (err != nil && tt.err == nil) || (err != nil && tt.err != nil && !strings.HasPrefix(err.Error(), tt.err.Error())):
			t.Errorf("%d: mismatched error, actual then expected", i)
			t.Logf("%v", err)
			t.Logf("%v", tt.err)
		case (m == nil && tt.m != nil) || (m != nil && tt.m == nil) || (m != nil && tt.m != nil && !reflect.DeepEqual(*m, *tt.m)):
			t.Errorf("%d: mismatched structs returned, actual then expected", i)
			t.Logf("%#v", m)
			t.Logf("%#v", tt.m)
		}
	}
}

func TestGetAndParseMetadata(t *testing.T) {
	jsonFile, _, err := getTestData()
	if err != nil {
		t.Fatalf("Unable to read test data file: %v", err)
	}

	u := fmt.Sprintf("file://%s", jsonFile)
	m, err := GetAndParseMetadata(u)
	if err != nil {
		t.Fatalf("Unexpected error getting and parsing metadata: %v", err)
	}
	if m == nil {
		t.Fatalf("Unexpected non-nil metadata struct  getting and parsing metadata")
	}
	if !reflect.DeepEqual(*m, *validMeta) {
			t.Errorf("mismatched structs returned, actual then expected")
			t.Logf("%#v", *m)
			t.Logf("%#v", *validMeta)
	}
}

