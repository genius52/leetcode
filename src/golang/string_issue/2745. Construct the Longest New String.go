package string_issue

import "strconv"

// x: aa
// y: bb
// z: ab
// aa bb ab aa
// bb ab aa bb
// ab aa bb ab

// bb ab aa bb ab aa
const (
	Undefined int = -1
	AA        int = 1
	BB        int = 2
	AB        int = 3
)

func dfs_longestString(x int, y int, z int, pre int, memo map[string]int) int {
	key := strconv.Itoa(pre) + "," + strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z)
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	var res int = 0
	if pre == Undefined {
		if x > 0 {
			res = max_int(res, 2+dfs_longestString(x-1, y, z, AA, memo))
		}
		if y > 0 {
			res = max_int(res, 2+dfs_longestString(x, y-1, z, BB, memo))
		}
		if z > 0 {
			res = max_int(res, 2+dfs_longestString(x, y, z-1, AB, memo))
		}
	} else if pre == AA {
		if y > 0 {
			res = max_int(res, 2+dfs_longestString(x, y-1, z, BB, memo))
		}
	} else if pre == BB || pre == AB {
		if x > 0 {
			res = max_int(res, 2+dfs_longestString(x-1, y, z, AA, memo))
		}
		if z > 0 {
			res = max_int(res, 2+dfs_longestString(x, y, z-1, AB, memo))
		}
	}
	memo[key] = res
	return res
}

func LongestString(x int, y int, z int) int {
	var memo map[string]int = make(map[string]int)
	return dfs_longestString(x, y, z, -1, memo)
}
