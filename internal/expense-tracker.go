package internal

import (
	"fmt"

	"github.com/headboot/feel-expense/internal/csvmanager"
	"github.com/headboot/feel-expense/pkg"
	"github.com/headboot/feel-expense/pkg/model"
)

type ExpenseTracker struct{}

type TrackerOptions struct {
	Description string
	Amount      int
	Id          int
	Month       uint8
}

func (tracker ExpenseTracker) Execute(command model.Command, options TrackerOptions) {
	switch command {
	case model.Add:
		{
			add(options)
		}
	case model.List:
		{
			list()
		}
	case model.Summary:
		{
			summary()
		}
	case model.Delete:
		{
			delete(options)
		}

	case model.Help:
		{
			help()
		}

	default:
		{
			pkg.ExitWithError(fmt.Errorf("unknown command %d", command))
		}
	}
}

func help() {
	printer := ConsolePrinter{}

	const helpText string = "Welcome to Feel-Expense \n———————————\nHere u can track your expences by using: \n\n1. --description <text> --amount <sum> add — to add a new expence \n2. list — to list all your expences\n3. summary — to get summary of your expenses \n4. --id <id> delete — to delete expense from list"

	printer.PrintHelpText(helpText)
}

func list() {
	fileManager := csvmanager.CsvManager{}

	expenses := fileManager.ReadFile()

	printer := ConsolePrinter{}

	printer.PrintTable(expenses)

}

func summary() {
	fileManager := csvmanager.CsvManager{}

	expenses := fileManager.ReadFile()

	var summary int

	expensLen := len(expenses)

	for i := 0; i < expensLen; i++ {
		summary += expenses[i].Amount
	}
	printer := ConsolePrinter{}

	printer.PrintSummary(summary)
}

func add(options TrackerOptions) {
	fileManager := csvmanager.CsvManager{}

	if options.Description == "" {
		pkg.ExitWithError(fmt.Errorf("nothing to save, at leats you  need to add --description"))
	}

	expense := model.New(options.Description, options.Amount)

		fileManager.WriteToFile(expense)
	
}

func delete(options TrackerOptions) {

		if options.Id < 0 {
			pkg.ExitWithError(fmt.Errorf("id cannot be less than 0"))
		}

	fileManager := csvmanager.CsvManager{}

	expenses := fileManager.ReadFile()

	for index, exp := range expenses {
		if exp.Id != uint(options.Id) {
			if len(expenses)-1 == index  {
				pkg.ExitWithError(fmt.Errorf("not found by id %d", options.Id))
			}
			continue
		}
		expenses = append(expenses[:index], expenses[index+1:]...)
		fmt.Println(exp.Id)
		fmt.Println(expenses)
		break
	}
	
	fileManager.OverwriteFileData(expenses)
}