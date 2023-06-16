package max_saraly

import (
	"sort"
	"strings"
)

func largestNumber(nums []string) string {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]+nums[j] > nums[j]+nums[i]
	})

	result := strings.Join(nums, "")

	return result
}
