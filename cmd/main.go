package main

import (
	"flag"
	"os"

	"github.com/headboot/feel-expense/internal"
	"github.com/headboot/feel-expense/pkg/model"
)

func main() {
	args := os.Args

	flagDescription := flag.String("description", "", "type your expense description")
	flagAmount := flag.Int("amount", 0, "type amount of an expense")
	flagId := flag.Int("id", -1, "type id of the task")
	flagMonth := flag.Int("month", -1, "number of month u want to get summary")

	flag.Parse()
	
	options := internal.TrackerOptions{
		Description: *flagDescription,
		Amount:      *flagAmount,
		Id:          *flagId,
		Month:       uint8(*flagMonth),
	}
	
	tracker := internal.ExpenseTracker{}

	last := len(args)-1
	command := model.FromStringToCommand(args[last])
	tracker.Execute(command, options)
}
