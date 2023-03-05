package simulation

func rotateMatrix(mat [][]int, n int) [][]int {
	// write code here
	mat = transpose(mat, n)

	for i := 0; i < n; i++ {
		left, right := 0, len(mat[i])-1
		for left < right {
			mat[i][left], mat[i][right] = mat[i][right], mat[i][left]
			left++
			right--
		}
	}
	return mat
}

func transpose(mat [][]int, n int) [][]int {
	m, n := len(mat)-1, len(mat[0])-1
	for i := 0; i <= m; i++ {
		for j := i; j <= n; j++ {
			if i == j {
				continue
			}
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	return mat
}
