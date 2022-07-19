package array

func Generate(numRows int) [][]int {
	var res [][]int
	for i := 0;i < numRows;i++{
		var tmp []int = make([]int,i + 1)
		tmp[0] = 1
		tmp[i] = 1
		res = append(res, tmp)
		for j := 1;j < i;j++{
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}