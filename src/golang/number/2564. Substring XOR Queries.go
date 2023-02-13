package number

import (
	"strconv"
	"strings"
)

func converToBianry(n int) string {
	if n == 0{
		return "0"
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func SubstringXorQueries(s string, queries [][]int) [][]int {
	var l2 int = len(queries)
	var res [][]int = make([][]int,l2)
	for i := 0;i < l2;i++{
		sub := converToBianry(queries[i][0] ^ queries[i][1])
		idx := strings.Index(s,sub)
		if idx == -1{
			res[i] = []int{-1,-1}
		}else{
			res[i] = []int{idx,idx + len(sub) - 1}
		}
	}
	return res
}