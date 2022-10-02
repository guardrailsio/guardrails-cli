package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	tmpl := `
  ____                     _ ____       _ _        ____ _     ___
 / ___|_   _  __ _ _ __ __| |  _ \ __ _(_) |___   / ___| |   |_ _|
| |  _| | | |/ _' | '__/ _' | |_) / _' | | / __| | |   | |    | |
| |_| | |_| | (_| | | | (_| |  _ < (_| | | \__ \ | |___| |___ | |
 \____|\__,_|\__,_|_|  \__,_|_| \_\__,_|_|_|___/  \____|_____|___|

Usage: guardrails <command> [<args>]

Commands:
  scan [-t,--token=<token>][-p,--path=<path>][-f,--format=json,csv,sarif,pretty]
       [-o,--output=<path>][-q,--quiet]
  help

scan scans a repository for vulnerabilities and output results
  -t, --token  a valid Guardrails CLI token you can obtain from dashboard > settings
  -p, --path   the path to the repository to scan, defaults to $PWD
  -f, --format the output format for scan results, defaults to pretty
  -o, --output if provided, will save the output to the specified file path
  -q, --quiet  if provided, will only output scan results in --format and nothing else

help displays this help menu

Environment variables:
GUARDRAILS_CLI_TOKEN  if set, will be used as token when --token is not provided
`

	rootCmd.SetHelpTemplate(tmpl)
}
