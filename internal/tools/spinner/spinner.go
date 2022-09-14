package spinner

import (
	"time"

	"github.com/briandowns/spinner"
)

func New() *spinner.Spinner {
	return spinner.New(spinner.CharSets[26], 300*time.Millisecond)
}
