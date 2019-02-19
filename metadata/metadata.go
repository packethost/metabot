package main

import (
        "http"
        "fmt"
        "io/ioutil"
        "os"

        "github.com/packethost/metabot/cmd"
)

const (
    metadataUrl = "http://packet.metadata.net/metadata"
)

func getMetadata() ([]byte, error) {
    // get metadata
    resp, err := http.Get(metadataUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

func parseMetadata(b []byte) (error) {

}

func GetAndParseMetadata() (error) {
    b, err := getMetadata()
    if err != nil {
        return err
    }
    parsed, err := parseMetdata(b)
    if err != nil {
        return err
    }
    return parsed, nil
}


//        fmt.Fprintf(os.Stderr, "error retrieving packet metadata: %v\n", err)
//        os.Exit(1)
