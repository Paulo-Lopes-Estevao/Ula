package worker

import (
	"log"
	"sync"

	"github.com/ebizno/Ula/internal/csv"
	"github.com/ebizno/Ula/internal/email"
	"github.com/ebizno/Ula/internal/file"
)

var waitgroup sync.WaitGroup

type Worker struct {
	Email     *email.Email
	FileName  string
	MaxWorker int
}

func NewWorker(email *email.Email, fileName string) Worker {
	return Worker{
		Email:    email,
		FileName: fileName,
	}
}

func (w Worker) WorkerMaxWorker(maxWorker int) Worker {
	return Worker{
		MaxWorker: maxWorker,
	}
}

func (w Worker) Start() {
	waitgroup.Add(w.MaxWorker)
}

func (w Worker) Stop() {
	defer waitgroup.Done()
}

func (w Worker) Wait() {
	waitgroup.Wait()
}

func (w Worker) Job() {
	w.Start()
}

func (w Worker) WorkerContentTypePlain(DataCsv csv.DataCsv) {
	defer w.Stop()
	if DataCsv.Email != "" {
		email, err := email.NewEmail(DataCsv.Email, w.Email.Subject, w.Email.Body, w.Email.IEmailCredential)
		if err != nil {
			log.Fatal(err)
		}
		email.SendEmailPlain()
	}
}

func (w Worker) WorkerSendEmail() []csv.DataCsv {
	filePath, err := file.NewFilePath(w.FileName)
	if err != nil {
		panic(err)
	}
	file, err := filePath.OpenFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	recordsData, err := csv.NewReaderFileCsv(file)
	if err != nil {
		panic(err)
	}
	csvData := csv.NewCsv(recordsData)

	amountOfEmail := csvData.TotalRowCountInCsvFile()

	data, err := csvData.AddCsvDataInStructJson()
	if err != nil {
		panic(err)
	}

	dataJson, err := csv.CsvToJson(data)
	if err != nil {
		panic(err)
	}

	worker := w.WorkerMaxWorker(amountOfEmail)
	worker.Start()

	return dataJson

}

func (w Worker) WorkerPlain() {
	dataJson := w.WorkerSendEmail()
	for _, v := range dataJson {
		go w.WorkerContentTypePlain(v)
	}
	w.Wait()
}
