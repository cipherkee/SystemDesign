package main

import (
	"fmt"
)

func main() {
	cal := NewCalculator([]Operation{Add, Sub, Abs})
	fmt.Println(cal.Evaluate(Add, []float64{2, 3}))
	fmt.Println(cal.Evaluate(Sub, []float64{9, 3}))
	fmt.Println(cal.Evaluate(Abs, []float64{-10}))
}
