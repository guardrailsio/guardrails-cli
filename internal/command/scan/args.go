package scan

import (
	"errors"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
		args.Format = "pretty"
	}

	return nil
}

func isFormatAllowed(value interface{}) error {
	input, ok := value.(string)
	if !ok {
		return errors.New("failed to parse format value")
	}

	allowedFormat := []string{"json", "csv", "sarif", "pretty"}

	var isAllowed bool
	for _, f := range allowedFormat {
		if input == f {
			isAllowed = true
		}
	}

	if !isAllowed {
		return errors.New("format unknown. Allowed format are json, csv, sarif, pretty")
	}
	return nil
}

func (args Args) Validate() error {
	return validation.ValidateStruct(&args,
		validation.Field(&args.Token, validation.Required.Error("can't find token in --token parameter or GUARDRAILS_CLI_TOKEN environment variable")),
		validation.Field(&args.Format, validation.By(isFormatAllowed)),
	)
}
