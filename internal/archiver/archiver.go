package archiver

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Archiver variable that does static check to make sure that archiver struct implements Archiver interface.
var _ Archiver = (*archiver)(nil)

//go:generate mockgen -destination=mock/archiver.go -package=mockarchiver . Archiver

// Archiver defines methods for archiving file / directory.
type Archiver interface {
	// OutputZipToFile zips filepaths and directories in projectPath and output it to designated path destination.
	OutputZipToFile(projectPath, outputPath string, filepaths []string) error
	// OutputZipToIOReader zips filepaths and directories in projectPath and output it to io.Reader to be consumed later.
	OutputZipToIOReader(projectPath string, filepaths []string) (io.Reader, error)
}

type archiver struct{}

// New instantiates new archiver.
func New() Archiver {
	return &archiver{}
}

// OutputZipToFile implements archiver.Archiver interface.
func (a *archiver) OutputZipToFile(projectPath, outputPath string, filepaths []string) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := outputFile.Close()
		if err == nil {
			err = closeErr
		}
	}()

	return zipper(projectPath, filepaths, outputFile)
}

// OutputZipToIOReader implements archiver.Archiver interface.
func (a *archiver) OutputZipToIOReader(projectPath string, filepaths []string) (io.Reader, error) {
	outputBuf := new(bytes.Buffer)
	if err := zipper(projectPath, filepaths, outputBuf); err != nil {
		return nil, err
	}

	return outputBuf, nil
}

func zipper(projectPath string, filepaths []string, output io.Writer) (err error) {
	w := zip.NewWriter(output)
	defer func() {
		closeErr := w.Close()
		if err == nil {
			err = closeErr
		}
	}()

	for _, path := range filepaths {
		// path variables only contains relative path of the project root, so we need to append project path to get absolute path.
		// path/filepath packages works cross platform which will use appropriate file separator based on OS.
		fullpath := filepath.Join(projectPath, path)

		in, err := os.Open(fullpath)
		if err != nil {
			return err
		}
		defer func() {
			closeErr := in.Close()
			if err == nil {
				err = closeErr
			}
		}()

		fileInfo, err := in.Stat()
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			continue
		}

		// Strip the absolute path up to the current directory, then trim off a leading
		// path separator (for Windows) and replace all instances of Windows path separators
		// with forward slashes as required by the w.Create method.
		f, err := w.Create(strings.Replace(strings.TrimPrefix(path, string(filepath.Separator)), "\\", "/", -1))
		if err != nil {
			return err
		}

		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
	}

	return nil
}
