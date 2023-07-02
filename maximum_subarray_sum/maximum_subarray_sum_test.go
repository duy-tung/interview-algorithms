package maximum_subarray_sum

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := maximumSubarraySum([]int{1, 2, 3, 4, 5}, 2)
		if result != 9 {
			t.Errorf("Expected 6, got %d", result)
		}
	})
}
