package josephus_problem

func JosephusProblem(n int, k int) int {
	if n == 1 {
		return 0
	}
	return (JosephusProblem(n-1, k) + k) % n
}
