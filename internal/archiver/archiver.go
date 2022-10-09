package archiver

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
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

	return compress(projectPath, filepaths, outputFile)
}

// OutputZipToIOReader implements archiver.Archiver interface.
func (a *archiver) OutputZipToIOReader(projectPath string, filepaths []string) (io.Reader, error) {
	outputBuf := new(bytes.Buffer)
	if err := compress(projectPath, filepaths, outputBuf); err != nil {
		return nil, err
	}

	return outputBuf, nil
}

func compress(projectPath string, filepaths []string, output io.Writer) (err error) {
	gzipWriter, err := gzip.NewWriterLevel(output, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := gzipWriter.Close()
		if err == nil {
			err = closeErr
		}
	}()
	tarWriter := tar.NewWriter(gzipWriter)
	defer func() {
		closeErr := tarWriter.Close()
		if err == nil {
			err = closeErr
		}
	}()

	for _, path := range filepaths {
		// path variables only contains relative path of the project root, so we need to append project path to get absolute path.
		// path/filepath packages works cross platform which will use appropriate file separator based on OS.
		fullpath := filepath.Join(projectPath, path)

		file, err := os.Open(fullpath)
		if err != nil {
			return err
		}
		defer func() {
			closeErr := file.Close()
			if err == nil {
				err = closeErr
			}
		}()

		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			continue
		}

		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}
		header.Name = strings.Replace(strings.TrimPrefix(path, string(filepath.Separator)), "\\", "/", -1)
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(tarWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}
