package main

import "github.com/guardrailsio/guardrails-cli/cmd"

var (
	version = "latest"
)

func main() {
	cmd.Execute(version)
}
