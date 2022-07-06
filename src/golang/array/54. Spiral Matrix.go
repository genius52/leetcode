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

func spiralOrder(matrix [][]int) []int{
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var res []int = make([]int, rows * columns)
	var idx int = 0
	var r_start int = 0
	var c_start int = 0
	var r_end int = rows - 1
	var c_end int = columns - 1
	for idx < rows * columns{
		//To right
		for j := c_start;j <= c_end;j++{
			res[idx] = matrix[r_start][j]
			idx++
		}
		r_start++
		if r_start > r_end{
			break
		}
		//To bottom
		for i := r_start;i <= r_end;i++{
			res[idx] = matrix[i][c_end]
			idx++
		}
		c_end--
		if c_end < c_start{
			break
		}
		//To left
		for j := c_end;j >= c_start;j--{
			res[idx] = matrix[r_end][j]
			idx++
		}
		r_end--
		if r_end < r_start{
			break
		}
		//To top
		for i := r_end;i >= r_start;i--{
			res[idx] = matrix[i][c_start]
			idx++
		}
		c_start++
		if c_start > c_end{
			break
		}
	}
	return res
}


func spiralOrder2(matrix [][]int) []int{
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var res []int = make([]int, rows * columns)
	var idx int = 0
	var dirs [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	var dir_idx int = 0
	var r_start int = 0
	var c_start int = 0
	var r_end int = rows - 1
	var c_end int = columns - 1
	var r int = 0
	var c int = 0
	for idx < rows * columns{
		for r >= r_start && r <= r_end && c >= c_start && c <= c_end {
			res[idx] = matrix[r][c]
			r += dirs[dir_idx][0]
			c += dirs[dir_idx][1]
			idx++
		}
		r -= dirs[dir_idx][0]
		c -= dirs[dir_idx][1]
		if dir_idx == 0{
			r_start++
		}else if dir_idx == 1{
			c_end--
		}else if dir_idx == 2{
			r_end--
		}else if dir_idx == 3{
			c_start++
		}
		dir_idx = (dir_idx + 1) % 4
		r += dirs[dir_idx][0]
		c += dirs[dir_idx][1]
	}
	return res
}