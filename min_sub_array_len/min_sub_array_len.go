package min_sub_array_len

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return nums[0]
		} else {
			return 0
		}
	}
	min := 0
	l, r := 0, 1
	sum := nums[l] + nums[r]

	for r < len(nums) {
		if sum != target {
			if nums[l] >= target || nums[r] >= target {
				return 1
			} else if sum < target {
				r++
				if r < len(nums) {
					sum += nums[r]
				}
			} else if r-l == 1 {
				min = 2
				l++
				r++
				sum = calSum(l, r, nums, sum)
			} else {
				min = getMin(min, l, r)
				l++
				sum -= nums[l-1]
			}
		} else {
			min = getMin(min, l, r)
			l++
			r++
			sum = calSum(l, r, nums, sum)
		}
	}

	return min
}

func calSum(l int, r int, nums []int, sum int) int {
	if r < len(nums) {
		sum += nums[r]
		sum -= nums[l-1]
	}

	return sum
}

func getMin(min int, l int, r int) int {
	if min == 0 {
		min = r - l + 1
	} else if r-l+1 < min {
		min = r - l + 1
	}

	return min
}
