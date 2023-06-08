package max_pairwise_product

func max_pairwise_product(arr []int) int {
	max1, max2 := -1, -1
	for _, num := range arr {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}
	return max1 * max2
}
