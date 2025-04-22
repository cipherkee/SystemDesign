package main

type IOperator interface {
	Evaluate([]float64) float64
}

type AddOp struct {
}

func (a *AddOp) Evaluate(arr []float64) float64 {
	x := arr[0]
	y := arr[1]
	return x + y
}

type SubOp struct {
}

func (a *SubOp) Evaluate(arr []float64) float64 {
	x := arr[0]
	y := arr[1]
	return x - y
}

type AbsOp struct {
}

func (a *AbsOp) Evaluate(arr []float64) float64 {
	x := arr[0]
	if x < 0 {
		return -1 * x
	}
	return x
}
