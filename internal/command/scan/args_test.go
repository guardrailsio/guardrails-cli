package scan

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestScanCommandArgsValidationFailed(t *testing.T) {
	type testCase struct {
		name  string
		input Args
	}

	testCases := []testCase{
		{
			name:  "missing_token",
			input: Args{},
		},
		{
			name: "invalid_format",
			input: Args{
				Token:  gofakeit.HexUint32(),
				Format: "invalid-format",
			},
		},
	}

	for _, tc := range testCases {
		err := tc.input.Validate()
		assert.NotNil(t, err)
	}
}
