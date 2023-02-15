package csv

import (
	"encoding/json"
	"errors"
)

type DataCsv struct {
	Email string `json:"Email"`
}

type Csv struct {
	FileData [][]string
}

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
		for c, col := range row {
			if col == "email" || col == "Email" {
				positionEmail = c
			}
		}
	}
	return positionEmail
}

func (c *Csv) CheckIfTheEmailHeaderExists() error {
	positionEmail := c.PositionOfTheEmailInTheCsvFile()
	for _, row := range c.FileData {
		if row[positionEmail] == "email" || row[positionEmail] == "Email" {
			return nil
		} else {
			return errors.New("email field not found")
		}
	}
	return nil
}

func (c *Csv) AddCsvDataInStructJson() []DataCsv {
	positionEmail := c.PositionOfTheEmailInTheCsvFile()
	var dataList []DataCsv
	for _, row := range c.FileData {
		var data DataCsv
		for index, field := range row {
			if index == positionEmail {
				if removeHeaderEmail(field) {
					data.Email = field
				}
			}
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func removeHeaderEmail(email string) bool {
	return email != "Email" && email != "email"
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
