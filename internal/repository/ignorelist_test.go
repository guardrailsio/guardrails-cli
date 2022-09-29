package repository

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestIsFileTypeIgnored(t *testing.T) {
	for _, ft := range ignoredFileTypes {
		filename := fmt.Sprintf("%s%s", gofakeit.AppName(), ft)
		assert.True(t, isFileTypeIgnored(filename))
	}
}
