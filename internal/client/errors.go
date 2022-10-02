package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))
)

// UnexpectedHTTPResponseFormatter catches error response returned by HTTP client.
func UnexpectedHTTPResponseFormatter(funcName string, statusCode int, respBody io.Reader) error {
	body, err := io.ReadAll(respBody)
	if err != nil {
		return err
	}

	return fmt.Errorf("unexpected %s response with http status code %d: %s", funcName, statusCode, string(body))
}
