package sum_of_digits

import "testing"

func TestSum(t *testing.T) {
	result := sum_of_digits(3, 5)
	if result != 8 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 8)
	}
}
