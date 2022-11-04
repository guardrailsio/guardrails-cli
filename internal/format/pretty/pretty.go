package pretty

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Error returns pretty formatted error message.
func Error(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(text.FgYellow.Sprintf("%s  Error: %s", emoji.Warning, err.Error()))
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

// Hyperlinks accepts links and texts and construct clickable hyperlinks in terminal.
// See: https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
func Hyperlinks(href, text string) string {
	return fmt.Sprintf(`\e]8;;%s\e\\%s\e]8;;\e\\\n`, href, text)
}
