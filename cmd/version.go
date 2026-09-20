package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// version is the current build version. Set manually until releases are automated.
const version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("spotify-cli %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
