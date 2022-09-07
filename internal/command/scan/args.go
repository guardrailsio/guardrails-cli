package scan

import (
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

func (args Args) Validate() error {
	return validation.ValidateStruct(&args,
		validation.Field(&args.Token, validation.Required.Error("can't find token in --token parameter or GUARDRAILS_CLI_TOKEN environment variable")),
	)
}
