package diagram

func check_conflict(seats *[][]byte,rows int,columns int,r int,c int)bool{
	var dirs [][]int = [][]int{{-1,-1},{-1,1},{0,-1},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns{
			continue
		}
		if (*seats)[next_r][next_c] == '1'{
			return false
		}
	}
	return true
}

func dfs_maxStudents(seats *[][]byte,rows int,columns int,idx int,students int,space int,status int64,visited map[int64]int)int{
	if idx >= rows * columns{
		return 0
	}
	//if *max_val > (students + space){
	//	return
	//}
	r := idx / columns
	c := idx % columns
	if _,ok := visited[status];ok{
		return visited[status]
	}
	var res int = 0
	if (*seats)[r][c] == '#'{
		res = dfs_maxStudents(seats,rows,columns,idx + 1,students,space,status,visited)
	}
	res = max_int(res,dfs_maxStudents(seats,rows,columns,idx + 1,students,space - 1,status,visited))
	if check_conflict(seats,rows,columns,r,c){
		//add a student here
		(*seats)[r][c] = '1'
		next_status := status | (1 << idx)

		res = max_int(res,1 + dfs_maxStudents(seats,rows,columns,idx + 1,students + 1,space - 1,next_status,visited))
		//remove this student
		(*seats)[r][c] = '.'
	}
	//visited[r][status] = res
	return res
}

func MaxStudents(seats [][]byte) int {
	var rows int = len(seats)
	var columns int = len(seats[0])
	var space int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if seats[i][j] == '.'{
				space++
			}
		}
	}
	var visited map[int64]int = make(map[int64]int)
	return dfs_maxStudents(&seats,rows,columns,0,0,space,0,visited)
}