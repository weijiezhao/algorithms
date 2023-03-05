package mapp

// BM52 数组中只出现一次的两个数字
// 知识点：map  时间复杂度：O(1)   空间复杂度：O(n)

func FindNumsAppearOnce(array []int) []int {
	// 1、统计出现次数
	numAppearTimesMap := make(map[int]int)
	for _, v := range array {
		numAppearTimesMap[v]++
	}

	// 2、取出仅出现一次的两个数值
	res := make([]int, 0, 2)
	for k, v := range numAppearTimesMap {
		if v == 1 {
			res = append(res, k)
		}
	}

	// 3、按非降序排列
	if res[0] < res[1] {
		return res
	}

	return []int{res[1], res[0]}
}
