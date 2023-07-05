package max_area

import "testing"

func TestMaxArea(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Errorf("Max area was incorrect, got: %d, want: %d.", result, 49)
	}
}
