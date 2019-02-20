package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/packethost/metabot/metadata"
)

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

func parseMetadata(b []byte) (*metadata.Metadata, error) {
	var m metadata.Metadata
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// GetAndParseMetadata retrieve metadata from a specific URL or Packet's standard
func GetAndParseMetadata(u string) (*metadata.Metadata, error) {
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
