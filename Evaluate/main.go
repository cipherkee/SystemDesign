package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

type Rule struct {
	ID         string
	Name       string
	RuleString string
}

type RuleStorage struct {
	Rules map[string]Rule // Keyedby rule ID
}

type RuleEngine struct {
	Storage RuleStorage
}

func (re *RuleEngine) AddRule(rule Rule) {
	re.Storage.Rules[rule.ID] = rule
}

func (re *RuleEngine) RemoveRule(ruleID string) {
	delete(re.Storage.Rules, ruleID)
}

func (re *RuleEngine) GetRule(ruleID string) (Rule, bool) {
	rule, exists := re.Storage.Rules[ruleID]
	return rule, exists
}

func (re *RuleEngine) Evaluate(context map[string]interface{}) (bool, error) {
	for key, val := range re.Storage.Rules {
		result, err := re.EvaluateRule(val.RuleString, context)
		if err != nil {
			return false, fmt.Errorf("error evaluating rule %s: %w", key, err)
		}
		if !result {
			return false, nil // If any rule fails, return false
		}
	}
	return true, nil // All rules passed
}

func (re *RuleEngine) EvaluateRule(rule string, context map[string]interface{}) (bool, error) {
	// Parse and compile the rule expression
	compiledExpression, err := expr.Compile(rule, expr.Env(context))
	if err != nil {
		return false, fmt.Errorf("failed to compile rule: %w", err)
	}

	// Evaluate the compiled expression
	result, err := expr.Run(compiledExpression, context)
	if err != nil {
		return false, fmt.Errorf("failed to evaluate rule: %w", err)
	}

	// Ensure the result is a boolean
	booleanResult, ok := result.(bool)
	if !ok {
		return false, fmt.Errorf("rule evaluation did not return a boolean result")
	}

	return booleanResult, nil
}

func main() {
	ruleEngine := RuleEngine{
		Storage: RuleStorage{
			Rules: make(map[string]Rule),
		},
	}

	// Example rule
	restaurantRule := Rule{
		ID:         "1",
		Name:       "restaurant_expense <= 75",
		RuleString: "restaurant_expense <= 75",
	}

	ruleEngine.AddRule(restaurantRule)

	// Example rule for no airfare expense
	airfareRule := Rule{
		ID:         "2",
		Name:       "no_airfare_expense",
		RuleString: "airfare_expense == 0",
	}

	ruleEngine.AddRule(airfareRule)

	entertainmentRule := Rule{
		ID:         "3",
		Name:       "no_entertainment_expense",
		RuleString: "entertainment_expense == 0",
	}
	ruleEngine.AddRule(entertainmentRule)

	totalExpenseRule := Rule{
		ID:         "4",
		Name:       "total_expense <= 100",
		RuleString: "total_expense <= 100",
	}
	ruleEngine.AddRule(totalExpenseRule)

	context := map[string]interface{}{
		"restaurant_expense":    50,
		"airfare_expense":       0,
		"entertainment_expense": 0,
		"total_expense":         90,
	}

	expense := &Expense{context: context, re: &ruleEngine}
	err := expense.Evaluate()
	if err != nil {
		fmt.Println("Error evaluating expense:", err)
	} else {
		fmt.Println("Expense status:", expense.GetStatus())
	}
}
