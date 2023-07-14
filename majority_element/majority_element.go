package majority_element

func majorityElement(nums []int) int {
	hash := make(map[int]float64)

	for _, v := range nums {
		hash[v]++

		if hash[v] > float64(len(nums)/2) {
			return 1
		}
	}

	return 0
}
