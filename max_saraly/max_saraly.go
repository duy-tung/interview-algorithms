package max_saraly

import "strings"

func largestNumber(nums []string) string {

	for j := 0; j < len(nums)-1; j++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i]+nums[i+1] < nums[i+1]+nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	result := strings.Join(nums, "")

	return result
}
