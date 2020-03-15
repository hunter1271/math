package math

import (
	"fmt"
	"github.com/GDG-Cloud-Innopolis/Go-begginners/l3/test"
	"math"
	"testing"
)

const precision = 1e-10

func TestCalc(t *testing.T) {
	tests := []struct {
		operation test.CalcOperation
		first     float64
		second    float64
		want      float64
	}{
		{
			operation: test.OperationAdd,
			first:     1.001,
			second:    144,
			want:      145.001,
		},
		{
			operation: test.OperationAdd,
			first:     4,
			second:    13,
			want:      17,
		},
		{
			operation: test.OperationSubtract,
			first:     3,
			second:    4,
			want:      -1,
		},
		{
			operation: test.OperationSubtract,
			first:     199,
			second:    27,
			want:      172,
		},
		{
			operation: test.OperationMultiply,
			first:     6,
			second:    7,
			want:      42,
		},
		{
			operation: test.OperationMultiply,
			first:     16,
			second:    0.25,
			want:      4,
		},
		{
			operation: test.OperationDivide,
			first:     10,
			second:    3,
			want:      3.33333333333,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Checking %f %s %f...", tt.first, tt.operation, tt.second)
		t.Run(name, func(t *testing.T) {
			operation := Calc(tt.operation)
			got := operation(tt.first, tt.second)

			if math.Abs(tt.want-got) > precision {
				t.Errorf("Expected %f, got %f", tt.want, got)
			}
		})
	}
}
