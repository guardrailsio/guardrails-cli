package sarif

import (
	"encoding/json"
	"fmt"
	"strconv"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/formatter/pretty"
)

const SCAVulnerabilityType = "sca"

type Schema struct {
	Version string `json:"version"`
	Schema  string `json:"$schema"`
	Runs    []Runs `json:"runs"`
}

type Runs struct {
	Tool struct {
		Driver ToolComponent `json:"driver"`
	} `json:"tool"`
	Results []Result `json:"results"`
}

type ToolComponent struct {
	Name  string                `json:"name"`
	Rules []ReportingDescriptor `json:"rules"`
}

type ReportingDescriptor struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ShortDescription struct {
		Text string `json:"text"`
	} `json:"shortDescription"`
	FullDescription struct {
		Text string `json:"text"`
	} `json:"fullDescription"`
	Help struct {
		Text     string `json:"text"`
		Markdown string `json:"markdown"`
	} `json:"help"`
	Properties struct {
		Tags    []string `json:"tags"`
		Problem struct {
			Severity string `json:"severity"`
		} `json:"problem"`
		SecuritySeverity string `json:"security-severity"`
	} `json:"properties"`
}

type Result struct {
	RuleID  string `json:"ruleId"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
	Locations           []Location `json:"locations"`
	PartialFingerprints struct {
		PrimaryLocationLineHash string `json:"primaryLocationLineHash"`
	} `json:"partialFingerprints"`
}

type Location struct {
	PhysicalLocation struct {
		ArtifactLocation struct {
			URI string `json:"uri"`
		} `json:"artifactLocation"`
	} `json:"physicalLocation"`
	Region struct {
		StartLine int `json:"startLine"`
	} `json:"region"`
}

func ScanResult(isQuiet bool, scanResult *guardrailsclient.GetScanDataResp) error {
	schema := &Schema{
		Schema:  "https://json.schemastore.org/sarif-2.1.0.json",
		Version: "2.1.0",
	}

	runs := make([]Runs, 0)
	rules := make([]ReportingDescriptor, 0)
	results := make([]Result, 0)
	for _, r := range scanResult.Results.Rules {
		for _, v := range r.Vulnerabilities {
			if v.Type == SCAVulnerabilityType {
				continue
			}

			docs := fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s", v.Language, r.Rule.Docs)
			helpText := "For more information, please see " + docs

			reportingDescriptor := ReportingDescriptor{
				ID:   strconv.FormatInt(v.EngineRule.EngineRuleID, 10),
				Name: v.EngineRule.EngineName,
				ShortDescription: struct {
					Text string `json:"text"`
				}{
					Text: v.EngineRule.Title,
				},
				FullDescription: struct {
					Text string `json:"text"`
				}{
					Text: v.EngineRule.Title,
				},
				Help: struct {
					Text     string `json:"text"`
					Markdown string `json:"markdown"`
				}{
					Text:     helpText,
					Markdown: helpText,
				},
				Properties: struct {
					Tags    []string `json:"tags"`
					Problem struct {
						Severity string `json:"severity"`
					} `json:"problem"`
					SecuritySeverity string `json:"security-severity"`
				}{
					Tags: []string{"security"},
					Problem: struct {
						Severity string `json:"severity"`
					}{
						Severity: "error",
					},
					SecuritySeverity: v.Metadata.CvssScore,
				},
			}
			rules = append(rules, reportingDescriptor)

			result := Result{
				RuleID: strconv.FormatInt(v.EngineRule.EngineRuleID, 10),
				Message: struct {
					Text string `json:"text"`
				}{
					Text: v.EngineRule.Title,
				},
				Locations: []Location{
					{
						PhysicalLocation: struct {
							ArtifactLocation struct {
								URI string `json:"uri"`
							} `json:"artifactLocation"`
						}{
							ArtifactLocation: struct {
								URI string `json:"uri"`
							}{
								URI: v.Path,
							},
						},
						Region: struct {
							StartLine int `json:"startLine"`
						}{
							StartLine: int(v.LineNumber),
						},
					},
				},
				PartialFingerprints: struct {
					PrimaryLocationLineHash string `json:"primaryLocationLineHash"`
				}{
					// TODO: provide the value when available
					PrimaryLocationLineHash: "",
				},
			}

			results = append(results, result)
		}
	}

	// TODO: assuming 1 driver (GuardRails) is single entry to runs
	if len(rules) > 0 && len(results) > 0 {
		runs = append(runs, Runs{
			Tool: struct {
				Driver ToolComponent `json:"driver"`
			}{
				Driver: ToolComponent{
					Name:  "GuardRails",
					Rules: rules,
				},
			},
			Results: results,
		})
	}

	schema.Runs = runs

	if !isQuiet && len(schema.Runs) == 0 && scanResult.Results.Count.Total > 0 {
		fmt.Printf("\n%s\n", prettyFmt.Success("No issues detected, well done!"))
		fmt.Printf("(%d vulnerabilities were detected, but SARIF format only reports on static analysis)\n", scanResult.Results.Count.Total)
	} else {
		manifestJson, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			return err
		}

		fmt.Printf("\n%s", string(manifestJson))
	}

	return nil
}
