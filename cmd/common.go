package cmd

import (
	"fmt"
	"os"

	"github.com/guardrailsio/guardrails-cli/internal/constant"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/format/pretty"
)

func exit(exitCode constant.ExitCode, err error) {
	if err != nil {
		fmt.Println(prettyFmt.Error(err))
	}
	os.Exit(int(exitCode))
}
