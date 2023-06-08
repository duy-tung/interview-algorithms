package max_dot_product

import "sort"

func max_dot_product(clicks []int, prices []int) int {
	revenue := 0
	sort.Ints(clicks)
	sort.Ints(prices)

	for i := 0; i < len(clicks); i++ {
		revenue += clicks[i] * prices[i]
	}
	return revenue
}
