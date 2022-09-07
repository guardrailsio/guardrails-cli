package cmd

import (
	"fmt"
	"os"

	prettyOut "github.com/guardrailsio/guardrails-cli/internal/output/pretty"
)

func fail(err error) {
	fmt.Println(prettyOut.Error(err))
	os.Exit(1)
}
