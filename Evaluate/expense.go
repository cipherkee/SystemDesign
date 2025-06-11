package main

type ExpenseStatus string

const (
	pendining   ExpenseStatus = "pending"
	approved    ExpenseStatus = "approved"
	rejected    ExpenseStatus = "rejected"
	underReview ExpenseStatus = "under_review"
)

type Expense struct {
	ID        string // Unique identifier for the expense
	context   map[string]interface{}
	approvers []string      // List of approvers for the expense
	status    ExpenseStatus // Status of the expense (e.g., "pending", "approved", "rejected")
	re        *RuleEngine
}

func (ex *Expense) UpdateStatus(newStatus ExpenseStatus) {
	ex.status = newStatus
}

func (ex *Expense) GetStatus() ExpenseStatus {
	return ex.status
}

func (ex *Expense) Evaluate() error {

	if ex.re == nil || len(ex.re.Storage.Rules) == 0 {
		return nil // No rules to evaluate
	}

	booleanResult, err := ex.re.Evaluate(ex.context)
	if err != nil {
		return err // Error evaluating rules
	}

	if !booleanResult {
		ex.UpdateStatus(underReview)
	} else {
		ex.UpdateStatus(approved)
	}
	return nil
}
