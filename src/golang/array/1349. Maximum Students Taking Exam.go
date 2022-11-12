package array

func check_conflict(seats [][]byte,rows int,columns int,r int,pre_status int,cur_status int)bool{
	var dirs [][]int = [][]int{{-1,-1},{-1,1},{0,-1},{0,1}}
	for j := 0;j < columns;j++{
		if (cur_status & (1 << j)) == 0{//当前位置是0，表示没有人
			continue
		}
		if seats[r][j] == '#'{ //当前位置有人，但是没有座位
			return true
		}
		for _,dir := range dirs{
			next_r := r + dir[0]
			next_c := j + dir[1]
			if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns{
				continue
			}
			if seats[next_r][next_c] == '.'{
				if next_r == r{
					if (cur_status & (1 << next_c)) != 0{
						return true
					}
				}else{
					if (pre_status & (1 << next_c)) != 0{
						return true
					}
				}
			}
		}
	}
	return false
}

func dp_maxStudents(seats [][]byte,rows int,columns int,r int,pre_status int,memo [][]int)int{
	if r == rows{
		return 0
	}
	if memo[r][pre_status] != -1{
		return memo[r][pre_status]
	}
	var res int = 0
	for i := 0;i < 1 << columns;i++{
		if !check_conflict(seats,rows,columns,r,pre_status,i){
			var n int = i
			var one_cnt int = 0
			for n > 0{
				if (n | 1) == n{
					one_cnt++
				}
				n >>= 1
			}
			res = max_int(res,one_cnt + dp_maxStudents(seats,rows,columns,r + 1,i,memo))
		}
	}
	memo[r][pre_status] = res
	return res
}

func MaxStudents(seats [][]byte) int {
	var rows int = len(seats)
	var columns int = len(seats[0])
	var memo [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		memo[i] = make([]int,1 << columns)
		for j := 0;j < 1 << columns;j++{
			memo[i][j] = -1
		}
	}
	return dp_maxStudents(seats,rows,columns,0,0,memo)
}

//func check_conflict(seats [][]byte,rows int,columns int,r int,c int)bool{
//	var dirs [][]int = [][]int{{-1,-1},{-1,1},{0,-1},{0,1}}
//	for _,dir := range dirs{
//		next_r := r + dir[0]
//		next_c := c + dir[1]
//		if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns{
//			continue
//		}
//		if seats[next_r][next_c] == '1'{
//			return true
//		}
//	}
//	return false
//}
//
//func dp_maxStudents(seats [][]byte,rows int,columns int,r int,c int,pre int,cur int,visited map[string]int)int{
//	if r == rows{
//		return 0
//	}
//	next_r := r + (c + 1)/columns
//	next_c := (c + 1) % columns
//	if seats[r][c] == '#'{
//		return dp_maxStudents(seats,rows,columns,next_r,next_c,pre,cur,visited)
//	}
//	key := strconv.Itoa(pre) + "," + strconv.Itoa(cur) + "," + strconv.Itoa(r)  + "," + strconv.Itoa(c)
//	if r > 0{
//		if _,ok := visited[key];ok{
//			return visited[key]
//		}
//	}
//	var res int = 0
//	ret := check_conflict(seats,rows,columns,r,c)
//	if ret{
//		res = dp_maxStudents(seats,rows,columns,next_r,next_c,pre,cur,visited)
//	}else{
//		//skip current seat
//		res = dp_maxStudents(seats,rows,columns,next_r,next_c,pre,cur,visited)
//		seats[r][c] = '1'
//		if next_r == r{
//			res = max_int(res,1 + dp_maxStudents(seats,rows,columns,next_r,next_c,pre,cur | (1 << c),visited))
//		}else{
//			res = max_int(res,1 + dp_maxStudents(seats,rows,columns,next_r,next_c,cur | (1 << c),0,visited))
//		}
//		seats[r][c] = '.'
//	}
//	visited[key] = res
//	return res
//}
//
//func MaxStudents(seats [][]byte) int {
//	var rows int = len(seats)
//	var columns int = len(seats[0])
//	var visited map[string]int = make(map[string]int)
//	return dp_maxStudents(seats,rows,columns,0,0,0,0,visited)
//}