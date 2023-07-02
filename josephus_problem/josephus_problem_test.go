package josephus_problem

import "testing"

func TestJosephusProblem(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := JosephusProblem(1, 1)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})
}
