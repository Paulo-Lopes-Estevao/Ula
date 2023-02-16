package worker_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"testing"

	csvIn "github.com/ebizno/Ula/internal/csv"
	"github.com/ebizno/Ula/internal/file"
)

var waitgroup sync.WaitGroup

func lengthEmail() int {
	email := []string{
		"test01@gmail.com",
		"test2@gmail.com",
		"test3@gmail.com",
		"test4@gmail.com",
	}

	return len(email)
}

func TestCreateWorkerQuantityWithExistingEmailQuantity(t *testing.T) {
	amountOfEmail := lengthEmail()
	if amountOfEmail != 4 {
		t.Error("Expected 4, got ", amountOfEmail)
	}

	waitgroup.Add(amountOfEmail)

	sum := 0
	for i := 0; i < amountOfEmail; i++ {
		go func() {
			defer waitgroup.Done()
			// do something
			sum += 1
			fmt.Println("do something", sum)
		}()
	}

}

func TestCreateWorkerQtyWithTheExisteingEmailQtyInTheCsvFile(t *testing.T) {
	// open csv file
	// read csv file
	// get the length of email
	// create worker quantity with the existing email quantity in the csv file

	file, err := file.NewFilePath("../../example/file/test.csv")
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	openFile, err := os.Open(file.Path)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	defer openFile.Close()

	fileData, err := csv.NewReader(openFile).ReadAll()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	csvData := csvIn.NewCsv(fileData)

	amountOfEmail := csvData.TotalRowCountInCsvFile()

	waitgroup.Add(amountOfEmail)

	sum := 0
	for i := 0; i < amountOfEmail; i++ {
		go func() {
			defer waitgroup.Done()
			// do something
			sum += 1
			fmt.Println("do something", sum)
		}()
	}
}

func TestReadExisteingEmailInTheCsvFile(t *testing.T) {

	file, err := file.NewFilePath("../../example/file/test.csv")
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	openFile, err := os.Open(file.Path)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	defer openFile.Close()

	fileData, err := csv.NewReader(openFile).ReadAll()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	csvData := csvIn.NewCsv(fileData)

	amountOfEmail := csvData.TotalRowCountInCsvFile()

	data, err := csvData.AddCsvDataInStructJson()
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	dataJson, err := csvIn.CsvToJson(data)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	waitgroup.Add(amountOfEmail)

	for i, v := range dataJson {
		go emailsReadByTheWorkers(v, i)
	}
}

func emailsReadByTheWorkers(data csvIn.DataCsv, index int) {
	defer waitgroup.Done()
	fmt.Println(index, "- Emails: ", data.Email)
}
