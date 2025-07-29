package internal

import (
	"fmt"
	"os"

	"github.com/headboot/feel-expense/pkg/model"
	"github.com/jedib0t/go-pretty/v6/table"
)

type ConsolePrinter struct{}

func getHeaderPreference() table.Row {
	return table.Row{
		"ID", "Description", "Amount", "Date",
	}
}

func (printer ConsolePrinter) PrintTable(data []model.Expense) {
	tab := table.NewWriter()
	tab.SetOutputMirror(os.Stdout)

	header := getHeaderPreference()

	tab.AppendHeader(header)

	for _, expense := range data {
		tab.AppendRow(table.Row{
			expense.Id, expense.Description, expense.Amount, expense.Date,
		})
	}

	tab.Render()

}

func (printer ConsolePrinter) PrintSummary(sum int) {
	fmt.Printf("summary amount: %d$\n", sum)
}

func (printer ConsolePrinter) PrintHelpText(text string) {
	fmt.Println(text)
}