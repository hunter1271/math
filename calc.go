package math

import (
	"github.com/GDG-Cloud-Innopolis/Go-begginners/l3/test"
)

// Calc is a function that return an operation function
func Calc(operation test.CalcOperation) func(a, b float64) float64 {
	switch operation {
	case test.OperationAdd:
		return add
	case test.OperationDivide:
		return divide
	case test.OperationMultiply:
		return multiply
	default:
		return subtract
	}
}

// Add operation. Returns sum of two operands
func add(a, b float64) float64 {
	return a + b
}

// Divide operation. Returns division result of first operand to second
func divide(a, b float64) float64 {
	return a / b
}

// Multiply operation. Returns product of two operands
func multiply(a, b float64) float64 {
	return a * b
}

// Subtract operation. Returns difference of two operands
func subtract(a, b float64) float64 {
	return a - b
}
