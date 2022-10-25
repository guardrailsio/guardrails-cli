package guardrailsclient

import (
	"io"
	"time"
)

// ErrorResp is general error body returned from guardrails client.
type ErrorResp struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

// CreateUploadURLReq is CreateUploadURL http request body.
type CreateUploadURLReq struct {
	CLIToken string `json:"clitoken"`
	File     string `json:"file"`
}

// CreateUploadURLResp is CreateUploadURL http response body.
type CreateUploadURLResp struct {
	SignedURL string `json:"signedUrl"`
}

// UploadProjectReq is UploadProject http request body.
type UploadProjectReq struct {
	UploadURL string
	File      io.Reader
}

// TriggerScanReq is TriggerScan http request body.
type TriggerScanReq struct {
	CLIToken   string `json:"clitoken"`
	Repository string `json:"repository"`
	SHA        string `json:"sha"`
	Branch     string `json:"branch"`
	FileName   string `json:"fileName"`
}

// TriggerScanResp is TriggerScan http response body.
type TriggerScanResp struct {
	ScanID       string `json:"idScan"`
	DashboardURL string `json:"dashboardUrl"`
}

// GetScanDataReq contains GetScanData parameters required to call GetScanData API.
type GetScanDataReq struct {
	ScanID string
}

// GetScanDataResp is GetScanData http response body.
type GetScanDataResp struct {
	ScanID  string `json:"idScan"`
	Type    string `json:"type"`
	Branch  string `json:"branch"`
	SHA     string `json:"sha"`
	OK      bool   `json:"ok"`
	Results struct {
		Count *GetScanDataCountResp `json:"count"`
		Rules []GetScanDataRuleResp `json:"rules"`
	} `json:"results"`
	Repository struct {
		RepositoryID  int64     `json:"idRepository"`
		Name          string    `json:"name"`
		DefaultBranch string    `json:"defaultBranch"`
		Provider      string    `json:"provider"`
		FullName      string    `json:"fullName"`
		Description   string    `json:"description"`
		Language      string    `json:"language"`
		IsPrivate     bool      `json:"isPrivate"`
		IsEnabled     bool      `json:"isEnabled"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	} `json:"repository"`
	Report     string    `json:"report"`
	QueuedAt   time.Time `json:"queuedAt"`
	ScanningAt time.Time `json:"scanningAt"`
	FinishedAt time.Time `json:"finishedAt"`
}

type GetScanDataCountResp struct {
	Total    int `json:"total"`
	New      int `json:"new"`
	Open     int `json:"open"`
	Resolved int `json:"resolved"`
	Fixed    int `json:"fixed"`
	Findings int `json:"findings"`
}

type GetScanDataRuleResp struct {
	Rule struct {
		RuleID int64  `json:"idRule"`
		Title  string `json:"title"`
		Name   string `json:"name"`
		Docs   string `json:"docs"`
	} `json:"rule"`
	Languages       []string                         `json:"language"`
	Count           *GetScanDataCountResp            `json:"count"`
	Vulnerabilities []GetScanDataVulnerabilitiesResp `json:"vulnerabilities"`
}

type GetScanDataVulnerabilitiesResp struct {
	FindingID               string `json:"idFinding"`
	Status                  string `json:"status"`
	Language                string `json:"language"`
	Branch                  string `json:"branch"`
	Path                    string `json:"path"`
	PrimaryLocationLineHash string `json:"primaryLocationLineHash"`
	LineNumber              int64  `json:"lineNumber"`
	IntroducedBy            string `json:"introducedBy"`
	Type                    string `json:"type"`
	Metadata                struct {
		DependencyName  string   `json:"dependencyName"`
		CurrentVersion  string   `json:"currentVersion"`
		PatchedVersions string   `json:"patchedVersions"`
		References      []string `json:"references"`
		CvssSeverity    string   `json:"cvssSeverity"`
		CvssScore       string   `json:"cvssScore"`
		CvssVector      string   `json:"cvssVector"`
	} `json:"metadata"`
	Severity struct {
		SeverityID int64  `json:"idSeverity"`
		Name       string `json:"name"`
	} `json:"severity"`
	EngineRule struct {
		EngineRuleID int64  `json:"idEngineRule"`
		Title        string `json:"title"`
		Name         string `json:"name"`
		Docs         string `json:"docs"`
		EngineName   string `json:"engineName"`
		CvssSeverity string `json:"cvssSeverity"`
		CvssScore    float64`json:"cvssScore"`
		CvssVector   string `json:"cvssVector"`
	} `json:"engineRule"`
}
