package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			twoSumII(nums, i, &result)
		}
	}

	return result
}

func twoSumII(nums []int, i int, result *[][]int) {
	left, right := i+1, len(nums)-1

	for left < right {
		sum := nums[i] + nums[left] + nums[right]
		if sum < 0 {
			left++
		} else if sum > 0 {
			right--
		} else {
			*result = append(*result, []int{nums[i], nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		}
	}
}
