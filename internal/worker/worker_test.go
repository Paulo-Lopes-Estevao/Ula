package worker_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"testing"

	csvIn "github.com/ebizno/Ula/internal/csv"
	"github.com/ebizno/Ula/internal/email"
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

func TestWorkersSendingEmailsPlain(t *testing.T) {
	var waitgroup sync.WaitGroup
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

	for index, v := range dataJson {
		go sendMail(v, t, index, &waitgroup)
	}

	waitgroup.Wait()

	fmt.Println("Done")

}

func sendMail(v csvIn.DataCsv, t *testing.T, index int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	if v.Email != "" {
		fmt.Println("Email: ", v)
		emailCredential, err := email.NewEmailCredential("ula@gmail.com", "xxxx", 587, "smtp.gmail.com")
		if err != nil {
			t.Errorf("Expected no error but got %s", err)
		}

		body := fmt.Sprintf("Sou Paulo estou a fazer teste de envio de email WORKER %d", index)
		subject := fmt.Sprintln("Teste Ula")

		email, err := email.NewEmail(v.Email, subject, body, emailCredential)
		if err != nil {
			t.Errorf("Expected no error but got %s", err)
		}
		email.SendEmailPlain()
	}
}
