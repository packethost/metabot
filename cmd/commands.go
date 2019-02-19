package cmd

// to hold our parsed and ready metadata
var metadata
var rootCmd = &cobra.Command{
  Use:   "metabot",
  Short: "metabot is your Packet metadata robot",
  Long: `Metabot is a metadata retriever and parser for Packet instance metadata.
                Complete documentation and source are available at https://github.com/packethost/metabot`,
  Run: func(cmd *cobra.Command, args []string) {
     var err error
     metadata, err = metadata.GetAndParseMetadata()
     if err != nil {
         fmt.Errorf("Error retrieving and parsing metadata: %v\n", err);
         os.Exit(1)
     }
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

