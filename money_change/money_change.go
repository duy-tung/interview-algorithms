package money_change

func Change(amount int) int {
	sum := 0
	coins := []int{10, 5, 1}

	for _, coin := range coins {
		sum += amount / coin
		amount %= coin
	}

	return sum
}
