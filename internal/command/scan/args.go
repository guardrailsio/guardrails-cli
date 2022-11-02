package scan

import (
	"errors"
	"fmt"
	"os"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	FormatPretty = "pretty"
	FormatJSON   = "json"
	FormatCSV    = "csv"
	FormatSARIF  = "sarif"
)

var (
	ErrMissingToken       = errors.New("missing token, please provide your GuardRails CLI token via -—token option or GUARDRAILS_CLI_TOKEN environment variable")
	ErrInvalidFormatParam = errors.New("failed to parse format value")
)

// Args provides arguments for scan command handler.
type Args struct {
	Token  string
	Path   string
	Format string
	Output string
	Quiet  bool
}

// SetDefault sets default value for some of args variables.
func (args *Args) SetDefault() error {
	if args.Path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		args.Path = cwd
	}
	if args.Format == "" {
		args.Format = FormatPretty
	}

	return nil
}

func isFormatAllowed(value interface{}) error {
	input, ok := value.(string)
	if !ok {
		return ErrInvalidFormatParam
	}

	allowedFormat := []string{FormatJSON, FormatCSV, FormatSARIF, FormatPretty}

	var isAllowed bool
	for _, f := range allowedFormat {
		if input == f {
			isAllowed = true
		}
	}

	if !isAllowed {
		return fmt.Errorf("unknown format. Allowed formats are %s", strings.Join(allowedFormat, ", "))
	}
	return nil
}

func (args Args) Validate() error {
	return validation.ValidateStruct(&args,
		validation.Field(&args.Token, validation.
			Required.Error(ErrMissingToken.Error())),
		validation.Field(&args.Format, validation.By(isFormatAllowed)),
	)
}
