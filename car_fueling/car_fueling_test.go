package car_fueling

import "testing"

func TestCarFueling(t *testing.T) {
	if car_fueling(10, 3, []int{1, 2, 5, 9}) != -1 {
		t.Errorf("Expected -1, got %d", car_fueling(10, 3, []int{1, 2, 5, 9}))
	}
}
