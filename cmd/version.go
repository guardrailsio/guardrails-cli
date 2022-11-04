package cmd

import (
	"fmt"
	"os"

	"github.com/guardrailsio/guardrails-cli/internal/constant"
	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(_ *cobra.Command, _ []string) {
		if version == "unknown" {
			fmt.Println("version unknown")
		} else {
			fmt.Printf("v%s\n", version)
		}
		os.Exit(int(constant.SuccessExitCode))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
