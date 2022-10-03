package scan

import (
	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	formatter "github.com/guardrailsio/guardrails-cli/internal/command/scan/format"
)

// GetScanDataFormatter parses guardrailsclient.GetScanDataResp to chosen format.
func (h *Handler) GetScanDataFormatter(resp *guardrailsclient.GetScanDataResp) error {
	w := h.OutputWriter.Writer

	switch h.Args.Format {
	case "json":
		return formatter.GetScanDataJSONFormat(w, resp)
	case "csv":
		return formatter.GetScanDataCSVFormat(w, resp)
	case "sarif":
		return formatter.GetScanDataSARIFFormat(w, resp, h.Args.Quiet)
	default:
		return formatter.GetScanDataPrettyFormat(w, resp)
	}
}
