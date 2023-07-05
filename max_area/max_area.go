package max_area

import "math"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0.0

	for left < right {
		result = math.Max(float64(result), cal(left, right, height))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return int(result)
}

func cal(a int, b int, height []int) float64 {
	return math.Min(float64(height[a]), float64(height[b])) * float64(b-a)
}
