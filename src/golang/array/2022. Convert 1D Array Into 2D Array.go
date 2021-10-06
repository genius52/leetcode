package array

func construct2DArray(original []int, m int, n int) [][]int {
	var l int = len(original)
	if l != (m * n){
		return [][]int{}
	}
	var res [][]int = make([][]int,m)
	for i := 0;i < m;i++{
		res[i] = make([]int,n)
	}
	for i := 0;i < l;i++{
		r := i / n
		c := i % n
		res[r][c] = original[i]
	}
	return res
}