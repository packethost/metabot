package cmd

import (
        "fmt"
        "os"

        "github.com/spf13/cobra"
        "github.com/packethost/metabot/metadata"
)

const (
    metadataURL = "https://metadata.packet.net/metadata"
)

// to hold our parsed and ready metadata
var (
  data *metadata.Metadata
  u string
)

var rootCmd = &cobra.Command{
  Use:   "metabot",
  Short: "metabot is your Packet metadata robot",
  Long: `Metabot is a metadata retriever and parser for Packet instance metadata.
                Complete documentation and source are available at https://github.com/packethost/metabot`,
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
    var err error
    data, err = metadata.GetAndParseMetadata(u)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error retrieving and parsing metadata: %v\n", err);
        os.Exit(1)
    }
  },
}

func init() {
     u = metadataURL
     rootCmd.PersistentFlags().StringVar(&u, "url", metadataURL, fmt.Sprintf("metadata URL (default is %s)",metadataURL))
}

// Execute run the command processor
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintf(os.Stderr, "%v\n", err);
    os.Exit(1)
  }
}

