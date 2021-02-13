package array

//Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//Output: [1,2,3,4,8,12,11,10,9,5,6,7]
func SpiralOrder(matrix [][]int) []int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var dir [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	var total int = rows * columns
	var pos int = 0
	var res []int
	var i int = 0
	var j int = 0
	var horizon bool = true
	var row_steps int = rows
	var col_steps int = columns
	for total > 0{
		var limit int = row_steps
		if horizon{
			limit = col_steps
		}
		for m := 0;m < limit;m++{
			res = append(res,matrix[i][j])
			i += dir[pos][0]
			j += dir[pos][1]
			total--
		}
		i -= dir[pos][0]
		j -= dir[pos][1]
		if horizon{
			row_steps--
		}else{
			col_steps--
		}
		horizon = !horizon
		pos = (pos + 1) % 4
		i += dir[pos][0]
		j += dir[pos][1]
	}
	return res
}