package segments_cover

import "testing"

func TestSegmentsCover(t *testing.T) {
	tests := []struct {
		segments [][]int
		expected []int
	}{
		{[][]int{{1, 3}, {2, 5}, {3, 6}}, []int{3}},
		{[][]int{{4, 7}, {1, 3}, {2, 5}, {5, 6}}, []int{3, 6}},
		{[][]int{{1, 3}, {2, 5}, {3, 6}, {7, 9}}, []int{3, 9}},
		{[][]int{{1, 3}, {2, 5}, {3, 6}, {7, 9}, {8, 10}}, []int{3, 9}},
	}

	for _, test := range tests {
		actual := SegmentsCover(test.segments)
		if len(actual) != len(test.expected) {
			t.Errorf("SegmentsCover(%v) = %v; expected %v", test.segments, actual, test.expected)
		}
	}
}
