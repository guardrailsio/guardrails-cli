package main

import "github.com/guardrailsio/guardrails-cli/cmd"

var (
	version = "unknown"
)

func main() {
	cmd.Execute(version)
}
