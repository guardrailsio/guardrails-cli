package pretty

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/jedib0t/go-pretty/text"
)

// Error returns pretty formatted error.
func Error(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s  %s", emoji.Warning, text.FgYellow.Sprint(err.Error()))
}
