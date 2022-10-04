package outputwriter

import (
	"bytes"
	"io"
	"os"

	"github.com/acarl005/stripansi"
)

// OutputWriter determines which output that the applications written to with the default output being stdout.
// When the output path is specified, it will store the data to buffer before being flushed out to other media such as file.
type OutputWriter struct {
	Path   string
	Buffer *bytes.Buffer
	Writer io.Writer
}

// New instantiates new *outputwriter.OutputWriter.
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

// SaveBufferToFile saves data written on the buffer to designated file path.
func (o *OutputWriter) SaveBufferToFile() error {
	// strip colorizer from text
	plainText := stripansi.Strip(o.Buffer.String())
	return os.WriteFile(o.Path, []byte(plainText), 0644)
}
