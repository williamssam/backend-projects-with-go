package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Amount      int        `json:"amount"`
	Category    string     `json:"category"`
	Month       time.Month `json:"month"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Expenses []Expense

func AddExpense(desc string, amount int) {
	fmt.Printf("Add expense with desc %v of amount %v", desc, amount)
}
