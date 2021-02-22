package array

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0;i < numRows;i++{
		var tmp []int = make([]int,i+1)
		tmp[0] = 1
		res = append(res, tmp)
		for j := 1;j <= i;j++{
			if i == j{
				res[i][j] = 1
			} else{
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}