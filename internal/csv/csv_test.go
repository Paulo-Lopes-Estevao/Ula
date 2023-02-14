package csv_test

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/ebizno/Ula/internal/file"
)

func TestOpenFileCsv(t *testing.T) {
	file, err := file.NewFilePath("../../example/file/test.csv")
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

func TestTotalRowCountInCsvFile(t *testing.T) {
	file, err := file.NewFilePath("../../example/file/test.csv")
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
	//ReadCsvFile
	fileData, err := csv.NewReader(openFile).ReadAll()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	//TotalRowCountInCsvFile
	totalRowCount := len(fileData)
	if totalRowCount == 0 {
		t.Errorf("Expected totalRowCount to be 3 but got %d", totalRowCount)
	}

	t.Log("Total Row Count: ", totalRowCount)

}
