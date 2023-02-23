package file

import (
	"path/filepath"
	"runtime"
	"testing"
)

func configPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../../example/file", "ula.csv")
}

func TestFilePath(t *testing.T) {

	file, err := NewFilePath("ula.csv")
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	if file.Path != configPath() {
		t.Errorf("Expected file.Path to be %s but got %s", configPath(), file.Path)
	}

	fileOpen, err := file.OpenFile()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	if fileOpen.Name() != configPath() {
		t.Errorf("Expected fileOpen.Name() to be %s but got %s", configPath(), fileOpen.Name())
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

func TestDirPath(t *testing.T) {

	file, err := NewFilePath("ula.csv")
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	if file.Path != configPath() {
		t.Errorf("Expected file.Path to be %s but got %s", configPath(), file.Path)
	}

	dir := file.Dir()

	if dir != "../../example/file" {
		t.Errorf("Expected dir to be %s but got %s", "../../example/file", dir)
	}

}

func TestNewPath(t *testing.T) {

	basepath := "../../file/ula.csv"

	filePath, err := NewFilePath(basepath)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	if filePath.Path != basepath {
		t.Errorf("Expected filePath.Path to be %s but got %s", basepath, filePath.Path)
	}

}
