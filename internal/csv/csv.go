package csv

import (
	"encoding/json"
	"errors"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type DataCsv struct {
	Email string `json:"Email"`
}

type Csv struct {
	FileData [][]string
}

const emailHeader = "email"

var (
	titleUpper  = strings.ToUpper(emailHeader)
	targetTitle = cases.Title(language.English, cases.Compact).String(emailHeader)
	targetLower = strings.ToLower(emailHeader)
)

func NewCsv(fileData [][]string) *Csv {
	return &Csv{
		FileData: fileData,
	}
}

func (c *Csv) TotalRowCountInCsvFile() int {
	return len(c.FileData)
}

func (c *Csv) PositionOfTheEmailInTheCsvFile() int {
	var positionEmail int
	for _, row := range c.FileData {
		positionEmail = searchThePositionOfTheEmailHeader(row)
		if positionEmail != -1 {
			return positionEmail
		}
	}
	return positionEmail
}

func searchThePositionOfTheEmailHeader(row []string) int {
	for c, col := range row {
		if checkIfTextEmailEquals(col) {
			return c
		}
	}
	return -1
}

func checkIfTextEmailEquals(email string) bool {
	return email == titleUpper || email == targetTitle || email == targetLower
}

func searchForEmailsInTheposition(row []string, data *DataCsv, targetPosition int) {
	for c, field := range row {
		if c == targetPosition {
			if removeHeaderEmail(field) {
				data.Email = field
			}
		}
	}
}

func (c *Csv) checkIfTheEmailHeaderExists() error {
	positionEmail := c.PositionOfTheEmailInTheCsvFile()
	if positionEmail == -1 {
		return errors.New("email header not found")
	}
	for _, row := range c.FileData {
		if checkIfTextEmailEquals(row[positionEmail]) {
			return nil
		} else {
			return errors.New("email header not found")
		}
	}
	return nil
}

func (c *Csv) AddCsvDataInStructJson() ([]DataCsv, error) {

	if err := c.checkIfTheEmailHeaderExists(); err != nil {
		return nil, err
	}

	positionEmail := c.PositionOfTheEmailInTheCsvFile()

	dataList := []DataCsv{}
	for _, row := range c.FileData {
		data := &DataCsv{}
		searchForEmailsInTheposition(row, data, positionEmail)
		dataList = append(dataList, *data)
	}
	return dataList, nil
}

func removeHeaderEmail(email string) bool {
	return email != titleUpper && email != targetTitle && email != targetLower
}

func CsvToJson(dataCsv []DataCsv) ([]DataCsv, error) {
	jsonData, err := json.MarshalIndent(dataCsv, "", "  ")
	if err != nil {
		return nil, err
	}

	data := []DataCsv{}
	json.Unmarshal(jsonData, &data)

	return data, nil

}
