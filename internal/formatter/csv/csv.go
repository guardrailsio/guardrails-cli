package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

func ScanResult(result *guardrailsclient.GetScanDataResp) error {
	var b []byte
	buf := bytes.NewBuffer(b)
	w := csv.NewWriter(buf)

	header := []string{"rule_id", "rule_title", "total", "finding_id", "path", "line_number", "docs"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, r := range result.Results.Rules {
		for _, v := range r.Vulnerabilities {
			ruleID := strconv.FormatInt(r.Rule.RuleID, 10)
			total := strconv.Itoa(r.Count.Total)
			lineNumber := strconv.FormatInt(v.LineNumber, 10)
			docs := fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s", v.Language, r.Rule.Docs)

			record := []string{ruleID, r.Rule.Title, total, v.FindingID, v.Path, lineNumber, docs}
			if err := w.Write(record); err != nil {
				return err
			}
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	fmt.Printf("\n%s", string(buf.Bytes()))

	return nil
}
