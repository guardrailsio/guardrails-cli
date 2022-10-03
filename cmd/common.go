package cmd

import (
	"fmt"
	"os"

	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/format/pretty"
)

func fail(err error) {
	fmt.Println(prettyFmt.Error(err))
	os.Exit(1)
}
