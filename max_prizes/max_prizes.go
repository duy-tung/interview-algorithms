package max_prizes

func MaxPrizes(n int) []int {
	k := 1
	for k*(k+1)/2 <= n {
		k++
	}
	k--

	delta := n - k*(k+1)/2
	prizes := make([]int, 0)

	for i := 1; i < k; i++ {
		prizes = append(prizes, i)
	}

	prizes = append(prizes, k+delta)
	return prizes
}
