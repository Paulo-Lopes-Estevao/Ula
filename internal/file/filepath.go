package file

import (
	"errors"
	"path/filepath"
)

type FilePath struct {
	Path string
}

var (
	ErrEmptyFilePath   = errors.New("empty file path")
	ErrInvalidFilePath = errors.New("invalid file path")
	ErrFileNotCsv      = errors.New("file is not a csv")
)

func NewFilePath(path string) (*FilePath, error) {
	p := &FilePath{
		Path: path,
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	if !p.validFileIsCsv() {
		return nil, ErrFileNotCsv
	}

	return p, nil

}

func (f *FilePath) FileName() string {
	return filepath.Base(f.Path)
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
	return filepath.Ext(f.Path) == ".csv"
}
