package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := threeSum([]int{-1, 0, 1, 2, -1, -4, 3, 1})
		if !reflect.DeepEqual(result, [][]int{{-4, 1, 3}, {-1, -1, 2}, {-1, 0, 1}}) {
			t.Errorf("Expected [[-4, 1, 3], [-1, -1, 2], [-1, 0, 1]], got %v", result)
		}
	})
}
