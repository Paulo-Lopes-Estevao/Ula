package file

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type (
	IFilePath interface {
		OpenFile() (*os.File, error)
		Dir() string
		FileName() string
	}
	FilePath struct {
		Name string
		Path string
	}
)

var (
	ErrEmptyFilePath   = errors.New("empty file path")
	ErrInvalidFilePath = errors.New("invalid file path")
	ErrFileNotCsv      = errors.New("file is not a csv")
)

func NewFilePath(filename string) (*FilePath, error) {
	file := &FilePath{
		Name: filename,
		Path: PathRoot(filename),
	}

	if err := file.validate(); err != nil {
		return nil, err
	}

	if !file.validFileIsCsv() {
		return nil, ErrFileNotCsv
	}

	return file, nil
}

func PathRoot(filename string) string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "/../../example/file", filename)
}

func (f *FilePath) OpenFile() (*os.File, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}

	if file == nil {
		return nil, errors.New("file is nil")
	}

	return file, nil
}

func (f *FilePath) Dir() string {
	return filepath.Dir(f.Path)
}

func (f *FilePath) FileName() string {
	return f.Name
}

func (f *FilePath) validate() error {
	if len(f.Path) == 0 {
		return ErrEmptyFilePath
	}

	if filepath.Clean(f.Path) != f.Path {
		return ErrInvalidFilePath
	}
	return nil
}

func (f *FilePath) validFileIsCsv() bool {
	return filepath.Ext(f.Name) == ".csv"
}
