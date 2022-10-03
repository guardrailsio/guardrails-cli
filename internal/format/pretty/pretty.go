package pretty

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/jedib0t/go-pretty/text"
)

// Error returns pretty formatted error message.
func Error(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s Error: %s", emoji.Warning, text.FgYellow.Sprint(err.Error()))
}

// Warning returns pretty formatted warning message.
func Warning(message string) string {
	if message == "" {
		return ""
	}

	return fmt.Sprintf("%s  %s", emoji.Warning, text.FgYellow.Sprint(message))
}

// Success returns pretty formatted success message.
func Success(message string) string {
	if message == "" {
		return ""
	}

	return fmt.Sprintf("%s  %s", emoji.CheckMark, text.FgGreen.Sprint(message))
}