package outputter

import (
	"encoding/json"
	"os"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

func SaveScanDataToFile(path string, result *guardrailsclient.GetScanDataResp) error {
	file, err := json.Marshal(result)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
