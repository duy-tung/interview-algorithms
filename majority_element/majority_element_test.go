package majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := majorityElement([]int{2, 2, 3})
		if result != 1 {
			t.Errorf("Expected 1, got %d", result)
		}
	})
}
