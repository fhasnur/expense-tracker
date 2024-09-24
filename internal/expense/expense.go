package expense

import "time"

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
