package main

import (
        "http"
        "fmt"
        "io/ioutil"
        "os"

        "github.com/spf13/cobra"
        "github.com/packethost/metabot/cmd"
)

func main() {
    cmd.Execute()
}

