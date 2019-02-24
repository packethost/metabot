package cmd

import (
	"github.com/packethost/metabot/internal"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestIpReport(t *testing.T) {
	tests := []struct {
		qualifiers string
		result     []string
	}{
		{"6 4", []string{}},
		{"6", []string{"2604:1380:1:9200::1/127"}},
		{"4", []string{"147.75.67.73/31", "10.99.185.1/31"}},
		{"", []string{"147.75.67.73/31", "10.99.185.1/31", "2604:1380:1:9200::1/127"}},
		{"address", []string{"147.75.67.73", "10.99.185.1", "2604:1380:1:9200::1"}},
		{"private", []string{"10.99.185.1/31"}},
		{"private 4 parent network", []string{"10.99.185.0/25"}},
	}
	for i, tt := range tests {
		result := ipReport(internal.GetValidMeta(), strings.Split(tt.qualifiers, " "))
		// sort so we can compare easily
		sort.Strings(result)
		sort.Strings(tt.result)
		if !reflect.DeepEqual(result, tt.result) {
			t.Errorf("%d %v: mismatched results, actual then expected", i, tt.qualifiers)
			t.Logf("%s\n", result)
			t.Logf("%s\n", tt.result)
		}
	}
}
