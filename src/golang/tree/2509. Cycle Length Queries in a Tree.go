package tree

import "math"

func cycleLengthQueries(n int, queries [][]int) []int {
	var l int = len(queries)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		n1 := queries[i][0]
		n2 := queries[i][1]
		var small int = min_int(n1,n2)
		var big int = max_int(n1,n2)
		var level1 int = int(math.Log2(float64(small)))
		var level2 int = int(math.Log2(float64(big)))
		var diff int = level2 - level1
		var cnt int = diff + 1
		for diff > 0{
			big >>= 1
			diff--
		}
		for small != big{
			small >>= 1
			big >>= 1
			cnt += 2
		}
		res[i] = cnt
	}
	return res
}