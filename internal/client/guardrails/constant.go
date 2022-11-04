package guardrailsclient

import "strings"

type CVSSSeverityAbbr int

const (
	Informational CVSSSeverityAbbr = iota
	Low
	Medium
	High
	Critical
	NotAvailable
)

func (c CVSSSeverityAbbr) String() string {
	return [...]string{"I", "L", "M", "H", "C", "N/A"}[c]
}

var cvssSeverityAbbrMap = map[string]CVSSSeverityAbbr{
	"informational": Informational,
	"low":           Low,
	"medium":        Medium,
	"high":          High,
	"critical":      Critical,
}

// GetCVSSSeverityAbbreviation returns abbreviations of CVSSSeverity value.
func GetCVSSSeverityAbbreviation(value string) string {
	value = strings.ToLower(value)

	abbr, ok := cvssSeverityAbbrMap[value]
	if !ok {
		return NotAvailable.String()
	}

	return abbr.String()
}
