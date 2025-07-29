package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/headboot/feel-expense/pkg"
)

const (
	YYYYMMDD = "2006-01-02"
)

type Expense struct {
	Id          uint
	Description string
	Amount      int
	Date        string
}

func New(description string, amount int) Expense {
	return Expense{
		Id:          uint(rand.Uint32()),
		Description: description,
		Amount:      amount,
		Date:        time.Now().Format(YYYYMMDD),
	}
}

type UpdateOptions struct {
	Description string
	Amount      int
}

func (e *Expense) Update(updated UpdateOptions) {
	if updated.Description != "" {
		e.Description = updated.Description
	}

	if updated.Amount != 0 && updated.Amount > 0 {
		e.Amount = updated.Amount
	}
}

func (e *Expense) ToCSV() []string {
	id := strconv.FormatUint(uint64(e.Id), 10)
	return []string{
		id, e.Description, fmt.Sprint(e.Amount), e.Date,
	}
}

func ExpenseFromCSV(s []string) Expense {
	id, err := strconv.ParseUint(s[0], 10, 64)
	amount, errAmount := strconv.ParseInt(s[2],10, 32)
	if err != nil {
		pkg.ExitWithError(err)
	}

	if errAmount != nil {
		pkg.ExitWithError(errAmount)
	}
	return Expense{
		Id: uint(id),
		Description: s[1],
		Amount: int(amount),
		Date: s[3],
	}
}