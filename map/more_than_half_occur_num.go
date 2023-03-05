package mapp

// BM51 数组中出现次数超过一半的数字
// 知识点：map  时间复杂度：O(n)   空间复杂度：O(1)
func MoreThanHalfNum_Solution(numbers []int) int {
	res := 0
	numOccurTimesMap := make(map[int]int)
	for _, val := range numbers {
		numOccurTimesMap[val]++
		if numOccurTimesMap[val] > len(numbers)/2 {
			res = val
			break
		}
	}
	return res
}
