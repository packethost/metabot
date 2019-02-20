package util

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"github.com/packethost/metabot/internal"
	"github.com/packethost/metabot/metadata"
)
const (
	testFilePath = "../cmd/testdata/metadata.json"
)

func TestGetMetadata(t *testing.T) {
	jsonFile, jsonData, err := internal.GetTestData()
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
	_, jsonData, err := internal.GetTestData()
	if err != nil {
		t.Fatalf("Unable to read test data file: %v", err)
	}

	tests := []struct{
		b []byte
		m *metadata.Metadata
		err error
	}{
		{[]byte(`"bad json"`), nil, fmt.Errorf("json: cannot unmarshal string into Go value of type metadata.Metadata")},
		{[]byte(`{"a":1}`), &metadata.Metadata{}, nil},
		{jsonData, internal.GetValidMeta(), nil},
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
	jsonFile, _, err := internal.GetTestData()
	validMeta := internal.GetValidMeta()
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

