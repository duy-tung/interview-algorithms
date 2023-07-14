package binary_search_with_duplicates

import "testing"

func TestBinarySearchWithDuplicates(t *testing.T) {
	arr := []int{2, 4, 4, 4, 7, 7, 9}
	target := 4
	result := binary_search_with_duplicates(arr, target)
	if result != 1 {
		t.Errorf("Expected 0, got %d", result)
	}
}
