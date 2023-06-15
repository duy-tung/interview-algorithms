package segments_cover

import (
	"sort"
)

func SegmentsCover(segments [][]int) []int {
	points := make([]int, 0)

	sort.Slice(segments, func(i, j int) bool {
		return segments[i][1] < segments[j][1]
	})

	minRight := -1

	for i := 0; i < len(segments); i++ {
		if minRight < segments[i][0] {
			minRight = segments[i][1]
			points = append(points, minRight)
		}
	}

	return points
}
