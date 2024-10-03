# Expense Tracker CLI

A Command-Line Interface (CLI) tool for tracking and managing personal expenses. It allows users to add expenses, delete records, view expense history, and generate summaries. The sample solution for the [expense-tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Installation

To use the Expense Tracker CLI, you need to have Go installed on your machine.

**Clone the Repository:**
```bash
git clone https://github.com/fhasnur/expense-tracker.git
```

**Navigate to the project directory:**
```bash
cd expense-tracker
```

**Install the required Go modules::**
```bash
go mod tidy
```

**Build the CLI:**
```bash
go build -o task-cli
```

## Usage

Once built, you can run the CLI tool from your terminal. Below are the example commands:

### Add an Expense

```bash
./expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
```

### List All Expenses

```bash
./expense-tracker list
# ID  | Date/Time            | Description     | Amount
# -------------------------------------------------------
# 1   | 2024-09-24 23:06     | Lunch           | $20
# 2   | 2024-09-30 18:31     | Dinner          | $10
# 3   | 2024-10-02 18:58     | Shopping        | $30
```

### Delete an Expense

```bash
./expense-tracker delete --id 1
# Expense deleted successfully
```

### View Expenses Summary

```bash
./expense-tracker summary
# Total expenses: $60
```

### Summary for a Specific Month

```bash
./expense-tracker summary --month 10
# Total expenses for October: $30
```

## File Structure

Expenses are stored in a JSON file located at `expenses.json` in the root folder. Each expense entry is structured as follows:

```json
{
  "id": 1,
  "description": "Lunch",
  "amount": 20,
  "created_at": "2024-09-24T23:06:02.7613895+07:00",
  "updated_at": "2024-09-24T23:06:02.7613895+07:00"
}
```

## Contributing

Feel free to submit pull requests or open issues for new features, improvements, or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.