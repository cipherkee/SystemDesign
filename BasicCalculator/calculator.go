package main

import (
	"errors"
)

type Calculator struct {
	operatorMap map[Operation]IOperator
}

func NewCalculator(ops []Operation) *Calculator {
	operatorMap := map[Operation]IOperator{}
	for _, o := range ops {
		switch o {
		case Add:
			operatorMap[Add] = &AddOp{}
		case Sub:
			operatorMap[Sub] = &SubOp{}
		case Abs:
			operatorMap[Abs] = &AbsOp{}
		}
	}
	return &Calculator{
		operatorMap: operatorMap,
	}
}

func (c *Calculator) Evaluate(o Operation, nums []float64) (float64, error) {
	operator, ok := c.operatorMap[o]
	if ok {
		return operator.Evaluate(nums), nil
	}
	return 0, errors.New("Operator not supported")
}
