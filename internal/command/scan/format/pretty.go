package scanformat

import (
	"fmt"
	"io"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/format/pretty"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/text"
)

// GetScanDataPrettyFormat parses guardrailsclient.GetScanDataResp to pretty format.
func GetScanDataPrettyFormat(w io.Writer, resp *guardrailsclient.GetScanDataResp) error {
	if resp.OK {
		fmt.Fprintf(w, "%s\n", prettyFmt.Success("No issues detected, well done!"))
	} else {
		issueStr := "issue"
		if resp.Results.Count.Total > 1 {
			issueStr += "s"
		}

		fmt.Fprintf(w, "%s\n", prettyFmt.Warning(fmt.Sprintf("We detected %d security %s", resp.Results.Count.Total, issueStr)))

		l := list.NewWriter()
		l.SetStyle(list.StyleBulletCircle)

		for _, r := range resp.Results.Rules {
			fmt.Fprintf(w, text.Bold.Sprintf("%s (%d)\n", r.Rule.Title, r.Count.Total))

			l.Indent()

			// get cvssSecurity abbreviation default value : N/A
			cvssSecurityAbbr := guardrailsclient.GetCVSSSeverityAbbreviation("")
			for _, v := range r.Vulnerabilities {
				vulnrTitle := v.EngineRule.Title

				if v.Metadata != nil {
					if v.Metadata.CvssSeverity != "" {
						cvssSecurityAbbr = guardrailsclient.GetCVSSSeverityAbbreviation(v.Metadata.CvssSeverity)
					} else if v.EngineRule.CvssSeverity != "" {
						cvssSecurityAbbr = guardrailsclient.GetCVSSSeverityAbbreviation(v.EngineRule.CvssSeverity)
					}

					if r.Rule.Title == "Vulnerable Libraries" {
						ok, err := v.Metadata.IsDependencyNameContainsVersion()
						if err != nil {
							return err
						}

						if ok {
							vulnrTitle = v.Metadata.DependencyName
						} else {
							vulnrTitle = fmt.Sprintf("%s@%s", v.Metadata.DependencyName, v.Metadata.CurrentVersion)
						}
					}
				}

				vulnrDocs := fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/general/%s\n", r.Rule.Docs)
				if v.EngineRule.Docs != "" {
					vulnrDocs = fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s\n", v.Language, v.EngineRule.Docs)
				}
				vulnrLink := prettyFmt.Hyperlinks(vulnrDocs, vulnrTitle)

				blobCaller := fmt.Sprintf("%s:%d", v.Path, v.LineNumber)
				// example permalink: https://github.com/guardrailsio/guardrails-cli/blob/67dbb9394658f163491989d931f196442747c295/README.md?plain=1#L3
				// TODO: how to construct the permalink using provided API response ?
				blobPermalink := "https://example.com"
				blobLink := prettyFmt.Hyperlinks(blobPermalink, blobCaller)

				l.AppendItem(text.FgCyan.Sprintf("(%s) %s - %s", cvssSecurityAbbr, vulnrLink, blobLink))
			}

			fmt.Fprintln(w, l.Render())
		}
	}

	return nil
}
