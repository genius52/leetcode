package array

func RotateGrid(grid [][]int, k int) [][]int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
	}
	for m := 0;m < min_int(rows,columns)/2;m++{
		var cur_data []int
		for i := m;i <= rows - 1 - m;i++{
			cur_data = append(cur_data,grid[i][m])
		}
		for j := m + 1;j < columns - 1 - m;j++{
			cur_data = append(cur_data,grid[rows - 1 - m][j])
		}
		for i := rows - 1 - m;i >= m ;i--{
			cur_data = append(cur_data,grid[i][columns - 1 - m])
		}
		for j := columns - 1 - m - 1;j > m;j--{
			cur_data = append(cur_data,grid[m][j])
		}
		cur_len := len(cur_data)
		var move int = k % cur_len
		var move_data []int = cur_data[cur_len - move:]
		move_data = append(move_data,cur_data[:cur_len - move]...)

		var idx int = 0
		for i := m;i <= rows - 1 - m;i++{
			res[i][m] = move_data[idx]
			idx++
		}
		for j := m + 1;j < columns - 1 - m;j++{
			res[rows - 1 - m][j] = move_data[idx]
			idx++
		}
		for i := rows - 1 - m;i >= m ;i--{
			res[i][columns - 1 - m] = move_data[idx]
			idx++
		}
		for j := columns - 1 - m - 1;j > m;j--{
			res[m][j] = move_data[idx]
			idx++
		}
	}
	return res
}