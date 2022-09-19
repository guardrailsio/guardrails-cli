package json

import (
	"encoding/json"
	"fmt"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

func ScanResult(result *guardrailsclient.GetScanDataResp) error {
	manifestJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("\n%s", string(manifestJson))

	return nil
}
