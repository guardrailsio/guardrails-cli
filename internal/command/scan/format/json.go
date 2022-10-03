package scanformat

import (
	"encoding/json"
	"fmt"
	"io"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

// GetScanDataJSONFormat parses guardrailsclient.GetScanDataResp to json format.
func GetScanDataJSONFormat(w io.Writer, resp *guardrailsclient.GetScanDataResp) error {
	manifestJson, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\n", string(manifestJson))

	return nil
}
