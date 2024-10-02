package expense

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewExpense(id int64, description string, amount float64) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func AddExpense(description string, amount float64) (int, error) {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return 0, err
	}

	var newExpenseId int64
	if len(expenses) == 0 {
		newExpenseId = 1
	} else {
		newExpenseId = expenses[len(expenses)-1].ID + 1
	}

	newExpense := NewExpense(newExpenseId, description, amount)
	expenses = append(expenses, *newExpense)

	if err := WriteExpensesToFile(expenses); err != nil {
		return 0, err
	}

	return int(newExpenseId), nil
}

func ListExpenses() ([]Expense, error) {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return nil, err
	}

	if len(expenses) == 0 {
		return nil, fmt.Errorf("no expenses found")
	}

	return expenses, nil
}

func DeleteExpense(id int64) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	var found bool
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	return WriteExpensesToFile(expenses)
}

func SummaryExpenses(month int) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		return fmt.Errorf("no expenses found")
	}

	var total float64
	var filteredExpenses []Expense

	if month > 0 {
		for _, expense := range expenses {
			if int(expense.CreatedAt.Month()) == month {
				total += expense.Amount
				filteredExpenses = append(filteredExpenses, expense)
			}
		}
		if len(filteredExpenses) == 0 {
			return fmt.Errorf("no expenses found for month %d", month)
		}
		monthName := time.Month(month).String()
		fmt.Printf("Total expenses for %s: $%.0f\n", monthName, total)
	} else {
		for _, expense := range expenses {
			total += expense.Amount
		}
		fmt.Printf("Total expenses: $%.0f\n", total)
	}

	return nil
}
