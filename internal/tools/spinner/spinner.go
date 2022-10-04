package spinner

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
)

// New instantiates new spinner.Spinner.
func New() *spinner.Spinner {
	return spinner.New(spinner.CharSets[26], 300*time.Millisecond, spinner.WithWriter(os.Stderr))
}
