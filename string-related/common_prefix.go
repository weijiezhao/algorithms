package stringrelated

import "strings"

// BM84 最长公共前缀

func longestCommonPrefix(ss []string) string {
	if len(ss) == 0 {
		return ""
	}

	commonPrefix := ""
	firstString := ss[0]
	for i := 0; i < len(firstString); i++ {
		commonPrefix += string(firstString[i])
		flag := 0
		for i := 1; i < len(ss); i++ {
			if !strings.HasPrefix(ss[i], commonPrefix) {
				commonPrefix = commonPrefix[0 : len(commonPrefix)-1]
				flag++
				break
			}
		}
		if flag != 0 {
			break
		}
	}
	return commonPrefix
}
