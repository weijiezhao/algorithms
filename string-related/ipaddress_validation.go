package stringrelated

import (
	"strconv"
	"strings"
)

// BM85 验证IP地址

func solve(ip string) string {
	if strings.Contains(ip, ".") && isIpv4(ip) {
		return "IPv4"
	}

	if strings.Contains(ip, ":") && isIpv6(ip) {
		return "IPv6"
	}

	return "Neither"
}

func isIpv4(ip string) bool {
	ss := strings.Split(ip, ".")
	if len(ss) != 4 {
		return false
	}

	for _, v := range ss {
		if v != "0" && strings.HasPrefix(v, "0") {
			return false
		}
		i, err := strconv.ParseInt(v, 10, 32)

		if err != nil {
			return false
		}

		if i < 0 || i > 255 {
			return false
		}
	}
	return true
}

func isIpv6(ip string) bool {
	ss := strings.Split(ip, ":")
	if len(ss) != 8 {
		return false
	}

	for _, v := range ss {
		if v == "" || strings.HasPrefix(v, "00") {
			return false
		}

		_, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			return false
		}
	}

	return true
}
