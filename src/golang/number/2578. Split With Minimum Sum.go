package number

import (
	"sort"
	"strconv"
	"strings"
)

func SplitNum(num int) int {
	data := strconv.Itoa(num)
	var s []string = strings.Split(data, "")
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	var l int = len(s)
	var n1 int = 0
	for i := 0; i < l; i += 2 {
		n1 *= 10
		n, _ := strconv.Atoi(s[i])
		n1 += n
	}
	var n2 int = 0
	for i := 1; i < l; i += 2 {
		n2 *= 10
		n, _ := strconv.Atoi(s[i])
		n2 += n
	}
	return n1 + n2
}
