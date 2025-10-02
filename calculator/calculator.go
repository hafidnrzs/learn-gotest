package calculator

import (
	"fmt"
	"math"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

// Complex Types
type Calculator struct {
	precision int
}

func (c *Calculator) DivideWithPrecision(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	return math.Round(result*math.Pow10(c.precision)) / math.Pow10(c.precision), nil
}
