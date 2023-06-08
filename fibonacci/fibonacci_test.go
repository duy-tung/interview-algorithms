package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	result := fibonacci(45)

	if result != 1134903170 {
		t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", result, 1134903170)
	}
}
