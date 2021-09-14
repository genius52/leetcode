package diagram

type Score_Path struct {
	score int
	path int64
}

func dp_pathsWithMaxScore(board *[]string,rows int,columns int,r int,c int,memo *[][]Score_Path)(int,int64){
	if r < 0 || c < 0{
		return 0,0
	}
	if r == 0 && c == 0{
		return 0,1
	}
	if (*board)[r][c] == 'X'{
		return 0,0
	}
	if (*memo)[r][c].path != -1{
		return (*memo)[r][c].score,(*memo)[r][c].path
	}

	var res Score_Path
	res.score,res.path = dp_pathsWithMaxScore(board,rows,columns,r - 1,c - 1,memo)
	s2,p2 := dp_pathsWithMaxScore(board,rows,columns,r - 1,c,memo)
	s3,p3 := dp_pathsWithMaxScore(board,rows,columns,r,c - 1,memo)
	if s2 > res.score{
		res.score = s2
		res.path = p2
		res.path %= (1e9 + 7)
	}else if s2 == res.score{
		res.path += p2
		res.path %= (1e9 + 7)
	}
	res.path %= (1e9 + 7)
	if s3 > res.score{
		res.score = s3
		res.path = p3
		res.path %= (1e9 + 7)
	}else if s3 == res.score{
		res.path += p3
		res.path %= (1e9 + 7)
	}
	val := int((*board)[r][c] - '0')
	(*memo)[r][c].score = res.score + val
	(*memo)[r][c].path = res.path
	return (*memo)[r][c].score,(*memo)[r][c].path
}

func PathsWithMaxScore(board []string) []int {
	var rows int = len(board)
	var columns int = len(board[0])
	var memo [][]Score_Path = make([][]Score_Path,rows)
	for i := 0;i < rows;i++{
		memo[i] = make([]Score_Path,columns)
		for j := 0;j < columns;j++{
			memo[i][j].path = -1
		}
	}
	var score int = 0
	var path int64 = 0
	s1,p1 := dp_pathsWithMaxScore(&board,rows,columns,rows - 2,columns - 1,&memo)
	s2,p2 := dp_pathsWithMaxScore(&board,rows,columns,rows - 1,columns - 2,&memo)
	s3,p3 := dp_pathsWithMaxScore(&board,rows,columns,rows - 2,columns - 2,&memo)
	if p1 != 0{
		if s1 > score{
			score = s1
			path = p1
		}
	}
	if p2 != 0{
		if s2 > score{
			score = s2
			path = p2
		}else if s2 == score{
			path += p2
			path %= (1e9 + 7)

		}
	}
	if p3 != 0 {
		if s3 > score{
			score = s3
			path = p3
		}else if s3 == score{
			path += p3
			path %= (1e9 + 7)
		}
	}
	return []int{score,int(path)}
}

func PathsWithMaxScore2(board []string) []int{
	var rows int = len(board)
	var columns int = len(board[0])
	var dp_score [][]int = make([][]int,rows)
	var dp_path [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp_score[i] = make([]int,columns)
		dp_path[i] = make([]int,columns)
	}
	dp_score[rows - 1][columns - 1] = 0
	dp_path[rows - 1][columns - 1] = 1
	var dirs [][]int = [][]int{{1,0},{0,1},{1,1}}
	for i := rows - 1;i >= 0;i--{
		for j := columns - 1;j >= 0;j--{
			if board[i][j] == 'S'{
				continue
			}else if board[i][j] == 'X'{
				dp_score[i][j] = 0
				dp_path[i][j] = 0
			}else{
				var max_score int = 0
				var max_path int = 0
				for _,dir := range dirs{
					pre_i := i + dir[0]
					pre_j := j + dir[1]
					if pre_i >= rows || pre_j >= columns{
						continue
					}
					if dp_path[pre_i][pre_j] == 0{
						continue
					}
					if dp_score[pre_i][pre_j] > max_score{
						max_score = dp_score[pre_i][pre_j]
						max_path = dp_path[pre_i][pre_j]
						max_path %= (1e9 + 7)
					}else if dp_score[pre_i][pre_j] == max_score{
						max_path += dp_path[pre_i][pre_j]
						max_path %= (1e9 + 7)
					}
				}
				if board[i][j] == 'E'{
					dp_score[i][j] = max_score
				}else{
					dp_score[i][j] = max_score + int(board[i][j] - '0')
				}
				dp_path[i][j] = max_path
			}
		}
	}
	return []int{dp_score[0][0],dp_path[0][0]}
}