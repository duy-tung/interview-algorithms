package fibonacci

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		prev, current = current, prev+current
	}
	return current
}
