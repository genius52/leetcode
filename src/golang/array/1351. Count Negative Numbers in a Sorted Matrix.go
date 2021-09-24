package array

//[4,3,2,-1]
//[3,2,1,-1]
//[1,1,-1,-2]
//[-1,-1,-2,-3]
func CountNegatives(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = 0
	var last_column int = columns - 1
	for i := 0;i < rows;i++{
		j := min_int(last_column,columns - 1)
		for ;j >= 0;j--{
			if grid[i][j] >= 0{
				break
			}
		}
		if j == -1{
			res += ((rows - i) * columns)
			break
		}
		last_column = j + 1;
		res += (columns - 1 -  j)
	}
	return res
}