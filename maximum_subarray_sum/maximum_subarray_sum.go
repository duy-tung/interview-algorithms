package maximum_subarray_sum

func maximumSubarraySum(nums []int, k int) int64 {
	max := 0
	temp := 0
	i := 0
	m := make(map[int]int)

	for ; i-k+1 < 0 && i < len(nums); i++ {
		add(m, nums, i)
		if max < max+nums[i] {
			max = max + nums[i]
		}
	}

	max = max + nums[i]
	add(m, nums, i)

	if i == len(nums)-1 && isDuplicate(m, k) {
		return 0
	}

	if (i == len(nums)-1) && !isDuplicate(m, k) {
		return int64(max)
	}

	if isDuplicate(m, k) {
		temp = max
		max = 0
	} else {
		temp = max
	}

	for i = i + 1; i < len(nums); i++ {
		add(m, nums, i)
		remove(m, nums, i-k)
		temp = temp + nums[i] - nums[i-k]
		if max < temp && !isDuplicate(m, k) {
			max = temp
		}
	}

	return int64(max)
}

func isDuplicate(m map[int]int, k int) bool {
	return len(m) < k
}

func add(m map[int]int, nums []int, i int) map[int]int {
	m[nums[i]]++
	return m
}

func remove(m map[int]int, nums []int, i int) map[int]int {
	if m[nums[i]] == 1 {
		delete(m, nums[i])
	}

	if m[nums[i]] > 1 {
		m[nums[i]]--
	}

	return m
}
