package main

import "fmt"

type Expense struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	MonthBudget int    `json:"month_budget"`
}

type Expenses []Expense

func AddExpense(desc string, amount int) {
	fmt.Printf("Add expense with desc %v of amount %v", desc, amount)
}
