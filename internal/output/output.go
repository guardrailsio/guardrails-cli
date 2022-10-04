package outputwriter

import (
	"bytes"
	"io"
	"os"

	"github.com/acarl005/stripansi"
)

type OutputWriter struct {
	Path   string
	Buffer *bytes.Buffer
	Writer io.Writer
}

func New(outputPath string) *OutputWriter {
	var w io.Writer
	buf := new(bytes.Buffer)

	if outputPath == "" {
		w = os.Stdout
	} else {
		w = io.MultiWriter(os.Stdout, buf)
	}

	return &OutputWriter{
		Path:   outputPath,
		Buffer: buf,
		Writer: w,
	}
}

func (o *OutputWriter) SaveBufferToFile() error {
	// strip colorizer from text
	plainText := stripansi.Strip(o.Buffer.String())
	return os.WriteFile(o.Path, []byte(plainText), 0644)
}
