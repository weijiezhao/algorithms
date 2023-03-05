package simulation

func solve(n int, m int, a []int) []int {
	q := m % n
	reverse(a, 0, n-1)
	reverse(a, 0, n-q-1)
	reverse(a, n-q, n-1)
	return a
}

func reverse(a []int, start, end int) {
	for start < end {
		swap(a, start, end)
		start++
		end--
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
