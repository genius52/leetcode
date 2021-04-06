package array

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var len_g int = len(g)
	var len_s int = len(s)
	i := 0
	j := 0
	for j < len_s && i < len_g{
		if s[j] < g[i]{
			j++
		}else{
			i++
			j++
		}
	}
	return i
}