package array

//Input: R = 1, C = 4, r0 = 0, c0 = 0
//Output: [[0,0],[0,1],[0,2],[0,3]]
func SpiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	var dirs [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	var res [][]int
	res = append(res,[]int{r0,c0})
	var total int = R * C - 1
	var row int = r0
	var col int = c0
	var steps int = 0
	var i int = 0
	for total > 0{
		if i | 1 != i{
			steps++
		}
		row_end := row + dirs[i][0] * steps
		col_end := col + dirs[i][1] * steps
		for !(row == row_end && col == col_end){
			row += dirs[i][0]
			col += dirs[i][1]
			if row >= 0 && row < R && col >= 0 && col < C{
				res = append(res,[]int{row,col})
				total--
				if total == 0{
					break
				}
			}
		}
		i++
		i = i % 4
	}
	return res
}