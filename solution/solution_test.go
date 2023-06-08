package solution

import "testing"

func TestSolution(t *testing.T) {
	result := Solution("10:00", "01:21")
	if result != 9 {
		t.Errorf("Expected 9, got %d", result)
	}
}
