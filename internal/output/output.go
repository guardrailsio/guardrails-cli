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
	return &OutputWriter{
		Path:   outputPath,
		Buffer: new(bytes.Buffer),
	}
}

func (o *OutputWriter) SetWriter() {
	if o.Path == "" {
		o.Writer = os.Stdout
	} else {
		o.Writer = io.MultiWriter(os.Stdout, o.Buffer)
	}
}

func (o *OutputWriter) SaveBufferToFile() error {
	// strip colorizer from text
	plainText := stripansi.Strip(o.Buffer.String())
	return os.WriteFile(o.Path, []byte(plainText), 0644)
}
