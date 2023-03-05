package simulation

func spiralOrder(matrix [][]int) []int {
	// write code here
	up, down := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	res := make([]int, 0)
	for left <= right && up <= down {
		for i := left; i < right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i < down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
