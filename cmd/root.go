package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command: the program itself, run without a subcommand.
var rootCmd = &cobra.Command{
	Use:   "spotify",
	Short: "Control Spotify from the terminal",
	Long: `spotify is a command line client for Spotify.

It lets you play, pause, skip, search and queue tracks
without leaving the terminal.`,
}

// Execute runs the root command. Cobra prints the error itself, so we only
// need to set a non-zero exit status.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
