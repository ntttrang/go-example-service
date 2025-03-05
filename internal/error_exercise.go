package internal

import (
	"fmt"
	"math"
)

// Define the custom error type
type ErrNegativeSqrt float64

// Implement the Error() method
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // Return an error for negative input
	}
	return math.Sqrt(x), nil // Use the built-in sqrt for positive numbers
}
