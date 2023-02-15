package csv_test

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"testing"

	csvIn "github.com/ebizno/Ula/internal/csv"
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

func TestGetAllEmailCsvFile(t *testing.T) {
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
	fileData, err := csv.NewReader(openFile).Read()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	//TotalRowCountInCsvFile
	totalRowCount := len(fileData)
	if totalRowCount == 0 {
		t.Errorf("Expected totalRowCount to be 3 but got %d", totalRowCount)
	}

	t.Log("Total Row Count: ", totalRowCount)

	//GetAllEmailCsvFile
	for _, row := range fileData {
		t.Log(row)
	}
}

func TestGetAllEmailCsv(t *testing.T) {
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

	totalRowCount := len(fileData)
	if totalRowCount == 0 {
		t.Errorf("Expected totalRowCount to be 3 but got %d", totalRowCount)
	}

	positionEmail := PositionEmail(t, fileData)

	for _, row := range fileData {
		for c, col := range row {
			if c == positionEmail {
				t.Log(col)
			}
		}
	}

}

func PositionEmail(t *testing.T, fileData [][]string) int {
	positionEmail := 0
	for _, row := range fileData {
		for c, col := range row {
			if col == "Email" || col == "email" {
				positionEmail = c
			}
		}
	}
	return positionEmail
}

type Data struct {
	Email string `json:"Email"`
}

func TestConvertCsvToJson(t *testing.T) {
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

	fileData, err := csv.NewReader(openFile).ReadAll()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	totalRowCount := len(fileData)
	if totalRowCount == 0 {
		t.Errorf("Expected totalRowCount to be 3 but got %d", totalRowCount)
	}

	positionEmail := PositionEmail(t, fileData)

	dataList := ConvertCsvToJson(fileData, positionEmail)

	jsonData, err := json.MarshalIndent(dataList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	data := []Data{}
	json.Unmarshal(jsonData, &data)

	for _, row := range data {
		if removeHeaderEmail(row.Email) {
			t.Log(row.Email)
		}

	}

}

func ConvertCsvToJson(fileData [][]string, positionEmail int) []Data {
	var dataList []Data
	for _, row := range fileData {
		var data Data
		for c, column := range row {
			if c == positionEmail {
				data.Email = column
			}
		}
		dataList = append(dataList, data)
	}

	return dataList
}

func removeHeaderEmail(email string) bool {
	return email != "Email" && email != "email"
}

func TestCSV(t *testing.T) {
	file, err := file.NewFilePath("../../example/file/test.csv")
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	openFile, err := os.Open(file.Path)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}
	fileData, err := csv.NewReader(openFile).ReadAll()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	csvData := csvIn.NewCsv(fileData)

	if err := csvData.CheckIfTheEmailHeaderExists(); err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	data := csvData.AddCsvDataInStructJson()
	dataJson, err := csvIn.CsvToJson(data)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}
	for _, row := range dataJson {
		if row.Email != "" {
			t.Log(row.Email)
		}
	}
}