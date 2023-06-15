package segments_cover

import "testing"

func TestSegmentsCover(t *testing.T) {
	segments := [][]int{
		{4, 7},
		{1, 3},
		{2, 5},
		{5, 6},
	}

	points := SegmentsCover(segments)

	if len(points) != 2 {
		t.Error("Expected 2 points, got", len(points))
	}

	if points[0] != 3 {
		t.Error("Expected 3, got", points[0])
	}

	if points[1] != 6 {
		t.Error("Expected 6, got", points[1])
	}
}
