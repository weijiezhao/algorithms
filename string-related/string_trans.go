package stringrelated

// BM83 字符串变形

import "strings"

func trans(s string, n int) string {
	newString := ""
	for i := 0; i < n; i++ {
		newString += changeCase(s[i])
	}
	ss := strings.Split(newString, " ")
	ss = reverse(ss)

	return strings.Join(ss, " ")
}

func reverse(ss []string) []string {
	left := 0
	right := len(ss) - 1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}
	return ss
}

func changeCase(c byte) string {
	if c >= 'A' && c <= 'Z' {
		return strings.ToLower(string(c))
	}

	if c >= 'a' && c <= 'z' {
		return strings.ToUpper(string(c))
	}

	return string(c)
}
