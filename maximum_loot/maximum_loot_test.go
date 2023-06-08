package maximum_loot

import (
	"fmt"
	"testing"
)

func TestMaximumLoot(t *testing.T) {
	fmt.Println(maximum_loot(10, []int{30}, []int{500}))
	if maximum_loot(10, []int{30}, []int{500}) != 166.667 {
		t.Errorf("Expected 166.667, got %f", maximum_loot(10, []int{30}, []int{500}))
	}
}
