package cmd

import (
	"context"

	"github.com/guardrailsio/guardrails-cli/internal/archiver"
	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	scan "github.com/guardrailsio/guardrails-cli/internal/command/scan"
	"github.com/guardrailsio/guardrails-cli/internal/config"
	outputwriter "github.com/guardrailsio/guardrails-cli/internal/output"
	"github.com/guardrailsio/guardrails-cli/internal/repository"
	spinner "github.com/guardrailsio/guardrails-cli/internal/tools/spinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	token  string
	path   string
	format string
	output string
	quiet  bool
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use: "scan",
	Run: func(_ *cobra.Command, _ []string) {
		ctx := context.Background()

		args := &scan.Args{
			Token:  token,
			Path:   path,
			Format: format,
			Output: output,
			Quiet:  quiet,
		}

		// set default value to fill up optional args that has empty value
		if err := args.SetDefault(); err != nil {
			fail(err)
		}
		if err := args.Validate(); err != nil {
			fail(err)
		}

		// instantiate configuration
		cfg := config.New()

		// setup git repository client
		repo, err := repository.New(args.Path)
		if err != nil {
			fail(err)
		}

		// setup archiver
		arc := archiver.New()

		// setup guardrails api client
		grclient := guardrailsclient.New(cfg.HttpClient, args.Token)

		// setup output writer
		outputWriter := outputwriter.New(args.Output)
		outputWriter.SetWriter()

		// setup spinner animation
		spinner := spinner.New()

		// inject all scan command dependencies and execute the command
		cmd := scan.New(args, spinner, cfg, repo, arc, outputWriter, grclient)
		if err := cmd.Execute(ctx); err != nil {
			fail(err)
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&token, "token", "t", "", "a valid Guardrails CLI token you can obtain from dashboard > settings")
	scanCmd.Flags().StringVarP(&path, "path", "p", "", "the path to the repository to scan, defaults to $PWD")
	scanCmd.Flags().StringVarP(&format, "format", "f", "pretty", "the output format for scan results, defaults to pretty")
	scanCmd.Flags().StringVarP(&output, "output", "o", "", "if provided, will save the output to the specified file path")
	scanCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "if provided, will only output scan results in --format and nothing else")

	// We can set token either from --token or GUARDRAILS_CLI_TOKEN envvar, with the later is more suitable in CICD usage
	// where secrets are usually stored in CICD's secret vault so it won't displayed in CICD pipeline logs.
	// If both are exists at the same time, the one from CLI params (--token) will override the one set in env var.
	if err := viper.BindEnv("token", "GUARDRAILS_CLI_TOKEN"); err != nil {
		fail(err)
	}

	if tokenEnv := viper.GetString("token"); token == "" && tokenEnv != "" {
		token = tokenEnv
	}

	rootCmd.AddCommand(scanCmd)
}
