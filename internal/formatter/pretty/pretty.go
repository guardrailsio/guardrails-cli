package pretty

import (
	"fmt"

	"github.com/enescakir/emoji"
	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	"github.com/jedib0t/go-pretty/text"
)

// Error returns pretty formatted error message.
func Error(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s Error: %s", emoji.Warning, text.FgYellow.Sprint(err.Error()))
}

// Warning returns pretty formatted warning message.
func Warning(message string) string {
	if message == "" {
		return ""
	}

	return fmt.Sprintf("%s  %s", emoji.Warning, text.FgYellow.Sprint(message))
}

// Success returns pretty formatted success message.
func Success(message string) string {
	if message == "" {
		return ""
	}

	return fmt.Sprintf("%s  %s", emoji.CheckMark, text.FgGreen.Sprint(message))
}

func ScanResult(result *guardrailsclient.GetScanDataResp) {
	if result.OK {
		fmt.Printf("%s\n", Success("No issues detected, well done!"))
	} else {
		fmt.Printf("%s\n", Warning(fmt.Sprintf("We detected %d security issue", result.Results.Count.Total)))

		for _, r := range result.Results.Rules {
			fmt.Printf("%s (%d)\n", r.Rule.Title, r.Count.Total)

			for _, v := range r.Vulnerabilities {
				fmt.Println(text.FgCyan.Sprintf("%s (line %d)", v.Path, v.LineNumber))
			}

			fmt.Println("Not sure how to fix this ?")
			for _, l := range r.Languages {
				fmt.Println(text.FgBlue.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s\n", l, r.Rule.Docs))
			}
		}
	}
}
