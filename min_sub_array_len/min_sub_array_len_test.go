package min_sub_array_len

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})
}
