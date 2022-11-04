package guardrailsclient

import (
	"encoding/json"
	"errors"
	"net/http"

	httpClient "github.com/guardrailsio/guardrails-cli/internal/client"
)

var (
	ErrInvalidToken            = errors.New("invalid token, please provide a valid GuardRails CLI token, available from dashboard -> settings")
	ErrRepositoryNotFound      = errors.New("invalid repository, please provide an existing repository from the git provider account linked with GuardRails, available from dashboard -> repositories")
	ErrScanProcessNotCompleted = errors.New("scan process is not completed")
)

func parseHTTPRespStatusCode(funcName string, resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil

	case http.StatusNotFound:
		errRespBody := new(ErrorResp)
		if err := json.NewDecoder(resp.Body).Decode(errRespBody); err != nil {
			return err
		}
		if errRespBody.Message == "Account not found" {
			return ErrInvalidToken
		}
		if errRespBody.Message == "Repository not found" {
			return ErrRepositoryNotFound
		}

		return httpClient.UnexpectedHTTPResponseFormatter(funcName, resp.StatusCode, resp.Body)

	default:
		return httpClient.UnexpectedHTTPResponseFormatter(funcName, resp.StatusCode, resp.Body)
	}
}
