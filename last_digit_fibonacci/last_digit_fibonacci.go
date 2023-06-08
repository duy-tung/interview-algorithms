package last_digit_fibonacci

func last_digit_fibonacci(n int) int {
	// 0 1 1 2 3 5 8
	if n < 2 {
		return n
	}

	prev, current := 0, 1

	for i := 2; i <= n; i++ {
		prev, current = current, (prev+current)%10
	}

	return current
}
