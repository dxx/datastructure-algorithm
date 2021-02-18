package stack

import "testing"

func TestCalculator(t *testing.T) {
    calculator := NewCalculator()

    calculator.Calculate("3+5*3-6")
    calculator.Calculate("30+5*3-6")
    calculator.Calculate("130+5*3-6")
}
