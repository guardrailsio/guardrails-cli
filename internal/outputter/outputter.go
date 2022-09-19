package outputter

import (
	"encoding/json"
	"io/ioutil"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

func SaveToFile(path string, result *guardrailsclient.GetScanDataResp) error {
	file, err := json.Marshal(result)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
