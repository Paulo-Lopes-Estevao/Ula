package file

import (
	"path/filepath"
	"runtime"
	"testing"
)

func configPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../../file", "test.csv")
}

func TestFilePath(t *testing.T) {

	basepath := configPath()

	base := filepath.Base(basepath)

	if base != "file" {
		t.Errorf("Expected base to be 'file' but got %s", base)
	}

	if filepath.Clean(basepath) != basepath {
		t.Errorf("Expected basepath to be clean but got %s", basepath)
	}

}

func TestFileNameCsv(t *testing.T) {

	basepath := configPath()

	base := filepath.Base(basepath)

	ext := filepath.Ext(base)

	if ext != ".csv" {
		t.Errorf("Expected base to be '.csv' but got %s", ext)
	}

	if filepath.Clean(basepath) != basepath {
		t.Errorf("Expected basepath to be clean but got %s", basepath)
	}

}

func TestNewPath(t *testing.T) {

	basepath := "../../file/test.csv"

	filePath, err := NewFilePath(basepath)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	if filePath.Path != basepath {
		t.Errorf("Expected filePath.Path to be %s but got %s", basepath, filePath.Path)
	}

}
