package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("v%s\n", version)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
