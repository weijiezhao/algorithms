package mapp

func minNumberDisappeared(nums []int) int {

	numBoolMap := make(map[int]bool)
	max := 0
	for _, v := range nums {
		numBoolMap[v] = true
		if v > max {
			max = v
		}
	}

	res := 0
	for i := 1; i <= max+1; i++ {
		if !numBoolMap[i] {
			res = i
			break
		}
	}
	return res
}
