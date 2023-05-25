package number

import "strconv"

func dfs_punishmentNumber(s string, l int, idx int, target int, sum int) bool {
	if idx == l {
		return target == sum
	}
	for i := idx + 1; i <= l; i++ {
		n, _ := strconv.Atoi(s[idx:i])
		ret := dfs_punishmentNumber(s, l, i, target, sum+n)
		if ret {
			return true
		}
	}
	return false
}

func check_punishmentNumber(n int) bool {
	var s string = strconv.Itoa(n * n)
	var l int = len(s)
	return dfs_punishmentNumber(s, l, 0, n, 0)
}

func PunishmentNumber(n int) int {
	var res int = 0
	for i := 1; i <= n; i++ {
		if check_punishmentNumber(i) {
			res += i * i
		}
	}
	return res
}
