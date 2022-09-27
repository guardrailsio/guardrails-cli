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

	fmt.Printf("%s\n", string(manifestJson))

	return nil
}
