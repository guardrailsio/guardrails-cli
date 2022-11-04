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

		for _, r := range resp.Results.Rules {
			l.SetStyle(list.StyleBulletCircle)

			fmt.Fprintln(w, text.Bold.Sprintf("%s (%d)", r.Rule.Title, r.Count.Total))

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

				docs := r.Rule.Docs
				if v.EngineRule.Docs != "" {
					docs = v.EngineRule.Docs
				}

				vulnrDocs := fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s", v.Language, docs)
				vulnrLink := prettyFmt.Hyperlinks(vulnrDocs, vulnrTitle)

				blobCaller := fmt.Sprintf("%s:%d", v.Path, v.LineNumber)
				// example permalink: https://github.com/guardrailsio/guardrails-cli/blob/67dbb9394658f163491989d931f196442747c295/README.md?plain=1#L3
				// TODO: construct the permalink using provided API response when the API is ready
				// blobPermalink := ""
				// blobLink := prettyFmt.Hyperlinks(blobPermalink, blobCaller)
				blobLink := blobCaller

				l.AppendItem(fmt.Sprintf("(%s) %s - %s", cvssSecurityAbbr, vulnrLink, blobLink))
			}

			fmt.Fprintln(w, l.Render())
			fmt.Fprintln(w)

			l.Reset()
		}
	}

	return nil
}
