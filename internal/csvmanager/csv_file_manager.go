package csvmanager

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/headboot/feel-expense/pkg"

	"github.com/headboot/feel-expense/pkg/model"
)

type CsvManager struct{}

const fileName string = "expense.csv"

func isFileExist() bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func (manager CsvManager) openFile() *os.File {
	if !isFileExist() {
		file, err := os.Create(fileName)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}


		return file
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND , os.ModeAppend)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}


	return file
}

func (manager CsvManager) ReadFile() []model.Expense {
	file := manager.openFile()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		pkg.ExitWithError(err)
	}

	result := []model.Expense{}

	for _, exp := range data {
		expModel := model.ExpenseFromCSV(exp)
		result = append(result, expModel)
	}

	defer file.Close()

	return result
}

func (manager CsvManager) WriteToFile(data model.Expense) {
	file := manager.openFile()

	writer := csv.NewWriter(file)

	csvData := data.ToCSV()

	err := writer.Write(csvData)


	defer file.Close()

	defer writer.Flush()
	

	if err != nil {
		pkg.ExitWithError(err)
	}
}

func (manager CsvManager) OverwriteFileData(data []model.Expense) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, os.ModeAppend)

	if err != nil {
		pkg.ExitWithError(err)
	}

	writer :=  csv.NewWriter(file)

	var csvData [][]string

	len := len(data)
	for i := 0; i < len; i++ {
		 csvData = append(csvData, data[i].ToCSV())
	}

	writer.WriteAll(csvData)
}