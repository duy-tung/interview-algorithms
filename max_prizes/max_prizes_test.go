package max_prizes

import "testing"

func TestMaxPrizes(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{2, []int{2}},
		{6, []int{1, 2, 3}},
		{8, []int{1, 2, 5}},
		{10, []int{1, 2, 3, 4}},
		{20, []int{1, 2, 3, 4, 10}},
		{100, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 22}},
	}

	for _, test := range tests {
		actual := MaxPrizes(test.n)
		if len(actual) != len(test.expected) {
			t.Errorf("MaxPrizes(%d) = %v; expected %v", test.n, actual, test.expected)
		}
	}
}
