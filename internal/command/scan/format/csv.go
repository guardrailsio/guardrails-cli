package scanformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

// GetScanDataCSVFormat parses guardrailsclient.GetScanDataResp to csv format.
func GetScanDataCSVFormat(w io.Writer, resp *guardrailsclient.GetScanDataResp) error {
	var b []byte
	buf := bytes.NewBuffer(b)
	csvWriter := csv.NewWriter(buf)

	header := []string{"rule_id", "rule_title", "total", "finding_id", "path", "line_number", "docs"}
	if err := csvWriter.Write(header); err != nil {
		return err
	}

	for _, r := range resp.Results.Rules {
		for _, v := range r.Vulnerabilities {
			ruleID := strconv.FormatInt(r.Rule.RuleID, 10)
			total := strconv.Itoa(r.Count.Total)
			lineNumber := strconv.FormatInt(v.LineNumber, 10)
			docs := fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s", v.Language, r.Rule.Docs)

			record := []string{ruleID, r.Rule.Title, total, v.FindingID, v.Path, lineNumber, docs}
			if err := csvWriter.Write(record); err != nil {
				return err
			}
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	fmt.Fprintf(w, "%s", buf.String())

	return nil
}
