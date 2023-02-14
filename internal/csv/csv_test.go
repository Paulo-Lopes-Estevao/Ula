package csv_test

import (
	"os"
	"testing"

	"github.com/ebizno/Ula/internal/file"
)

func TestOpenFileCsv(t *testing.T) {
	file, err := file.NewFilePath("../../file/test.csv")
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	openFile, err := os.Open(file.Path)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	defer openFile.Close()

	if openFile == nil {
		t.Errorf("Expected openFile to be not nil")
	}

}
