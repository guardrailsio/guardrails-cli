package scanformat

import (
	"fmt"
	"io"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/format/pretty"
	"github.com/jedib0t/go-pretty/text"
)

// GetScanDataJSONFormat parses guardrailsclient.GetScanDataResp to pretty format.
func GetScanDataPrettyFormat(w io.Writer, resp *guardrailsclient.GetScanDataResp) error {
	if resp.OK {
		fmt.Fprintf(w, "%s\n", prettyFmt.Success("No issues detected, well done!"))
	} else {
		issueStr := "issue"
		if resp.Results.Count.Total > 1 {
			issueStr += "s"
		}

		fmt.Fprintf(w, "%s\n", prettyFmt.Warning(fmt.Sprintf("We detected %d security %s", resp.Results.Count.Total, issueStr)))

		for _, r := range resp.Results.Rules {
			fmt.Fprintf(w, "%s (%d)\n", r.Rule.Title, r.Count.Total)

			for _, v := range r.Vulnerabilities {
				fmt.Fprintln(w, text.FgCyan.Sprintf("%s (line %d)", v.Path, v.LineNumber))
			}

			fmt.Fprintln(w, "Not sure how to fix this?")
			for _, l := range r.Languages {
				fmt.Fprintln(w, text.FgBlue.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s\n", l, r.Rule.Docs))
			}
		}
	}

	return nil
}
