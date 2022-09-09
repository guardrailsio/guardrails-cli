package archiver

import (
	"archive/zip"
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Archive list of tasks
// 1. Prevent recursive archiving
// 2. Check the result in other OS

// Archiver variable that does static check to make sure that archiver struct implements Archiver interface.
var _ Archiver = (*archiver)(nil)

//go:generate mockgen -destination=mock/archiver.go -package=mockarchiver . Archiver

// Archiver defines methods for archiving file / directory.
type Archiver interface {
	OutputZipToFile(outputPath, projectPath string) error
	OutputZipToReader(projectPath string) (io.Reader, error)
}

type archiver struct{}

// New instantiates new archiver.
func New() Archiver {
	return &archiver{}
}

// OutputZipToFile zips files and directories in projectPath and output it to designated path destination.
func (a *archiver) OutputZipToFile(outputPath, projectPath string) error {
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

	return zipper(outputFile, projectPath)
}

// OutputZipToFile zips files and directories in projectPath and output it to io.Reader to be consumed later.
func (a *archiver) OutputZipToReader(projectPath string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if err := zipper(buf, projectPath); err != nil {
		return nil, err
	}

	return buf, nil
}

func zipper(output io.Writer, projectPath string) error {
	var err error

	w := zip.NewWriter(output)
	defer func() {
		closeErr := w.Close()
		if err == nil {
			err = closeErr
		}
	}()

	err = filepath.Walk(projectPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		// Strip the absolute path up to the current directory, then trim off a leading
		// path separator (for Windows) and replace all instances of Windows path separators
		// with forward slashes as required by the w.Create method.
		f, err := w.Create(strings.Replace(strings.TrimPrefix(strings.TrimPrefix(path, projectPath), string(filepath.Separator)), "\\", "/", -1))
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
