package last_digit_fibonacci

import "testing"

func TestLastDigitFibonacci(t *testing.T) {
	result := last_digit_fibonacci(91239)

	if result != 6 {
		t.Errorf("LastDigitFibonacci was incorrect, got: %d, want: %d.", result, 6)
	}
}
