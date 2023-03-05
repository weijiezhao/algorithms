package stringrelated

import "strconv"

func solveBigNumberAdd(s string, t string) string {
	tmp, n, m := 0, len(s)-1, len(t)-1
	res := ""

	for n >= 0 || m >= 0 || tmp != 0 {
		i, j := 0, 0
		if n >= 0 {
			i = int(s[n] - '0')
			n--
		}
		if m >= 0 {
			j = int(t[m] - '0')
			m--
		}
		tmp += i + j
		res = strconv.Itoa(tmp%10) + res
		tmp = tmp / 10
	}

	return res
}
