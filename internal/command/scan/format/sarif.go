package scanformat

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/samber/lo"

	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/format/pretty"
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

// GetScanDataJSONFormat parses guardrailsclient.GetScanDataResp to SARIF format.
func GetScanDataSARIFFormat(w io.Writer, resp *guardrailsclient.GetScanDataResp, isQuiet bool) error {
	schema := &Schema{
		Schema:  "https://json.schemastore.org/sarif-2.1.0.json",
		Version: "2.1.0",
	}

	runs := make([]Runs, 0)
	rules := make([]ReportingDescriptor, 0)
	results := make([]Result, 0)
	for _, r := range resp.Results.Rules {
		for _, v := range r.Vulnerabilities {
			if v.Type == SCAVulnerabilityType {
				continue
			}

			id := strconv.FormatInt(v.EngineRule.EngineRuleID, 10)
			docs := fmt.Sprintf("https://docs.guardrails.io/docs/vulnerabilities/%s/%s", v.Language, r.Rule.Docs)
			helpText := "For more information, please see " + docs

			reportingDescriptor := ReportingDescriptor{
				ID:   id,
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
				RuleID: id,
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
					PrimaryLocationLineHash: v.PrimaryLocationLineHash,
				},
			}

			results = append(results, result)
		}
	}

	// TODO: assuming 1 driver (GuardRails) is single entry to runs
	runs = append(runs, Runs{
		Tool: struct {
			Driver ToolComponent `json:"driver"`
		}{
			Driver: ToolComponent{
				Name: "GuardRails",
				// rules must be unique. Identified by EngineRule.EngineRuleID
				Rules: lo.UniqBy(rules, func(r ReportingDescriptor) string {
					return r.ID
				}),
			},
		},
		Results: results,
	})

	schema.Runs = runs

	if (len(results) == 0 && resp.Results.Count.Total > 0) || resp.Results.Count.Total == 0 {
		if !isQuiet {
			fmt.Fprintf(w, "%s\n", prettyFmt.Success("No issues detected, well done!"))
			fmt.Fprintf(w, "(%d vulnerabilities were detected, but SARIF format only reports on static analysis)\n", resp.Results.Count.Total)
		} else {
			manifestJson, err := json.MarshalIndent(schema, "", "  ")
			if err != nil {
				return err
			}

			fmt.Fprintf(w, "%s\n", string(manifestJson))
		}
	} else {
		manifestJson, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			return err
		}

		fmt.Fprintf(w, "%s\n", string(manifestJson))
	}

	return nil
}
