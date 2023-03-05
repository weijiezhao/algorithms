package mapp

// BM50 两数之和
// 知识点：map  时间复杂度：O(n)   空间复杂度：O(n)

func twoSum(numers []int, target int) []int {

	numIndexMap := make(map[int]int)
	for i, num := range numers {
		if k, ok := numIndexMap[target-num]; ok {
			return []int{k + 1, i + 1}
		}
		numIndexMap[num] = i
	}
	return nil
}
