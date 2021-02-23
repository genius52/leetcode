package array

import "strconv"

//Example:
//X X X X
//X O O X
//X X O X
//X O X X
//After running your function, the board should be:
//
//X X X X
//X X X X
//X X X X
//X O X X
func fill_x(board [][]byte,r int,c int,rows int,columns int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if board[r][c] == 'X'{
		return
	}
	board[r][c] = 'X'
	fill_x(board,r - 1,c,rows,columns)
	fill_x(board,r + 1,c,rows,columns)
	fill_x(board,r,c - 1,rows,columns)
	fill_x(board,r,c + 1,rows,columns)
}

func can_be_replace(board [][]byte,r int,c int,rows int,columns int,visited map[string]bool)bool {
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return false
	}
	if board[r][c] == 'X'{
		return true
	}
	if r == 0 || r == (rows - 1) || c == 0 || c == (columns - 1){
		return false
	}
	k := strconv.Itoa(r) + "," + strconv.Itoa(c)
	if _,ok := visited[k];ok{
		return true
	}
	visited[k] = true
	return can_be_replace(board,r - 1,c,rows,columns,visited) && can_be_replace(board,r + 1,c,rows,columns,visited) &&
		can_be_replace(board,r,c - 1,rows,columns,visited) && can_be_replace(board,r,c + 1,rows,columns,visited)
}

func solve(board [][]byte)  {
	rows := len(board)
	if rows == 0{
		return
	}
	columns := len(board[0])
	var record map[string]bool = make(map[string]bool)
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if board[i][j] != 'X'{
				k := strconv.Itoa(i) + "," + strconv.Itoa(j)
				if _,ok := record[k];ok{
					continue
				}
				var visited map[string]bool = make(map[string]bool)
				ret := can_be_replace(board,i,j,rows,columns,visited)
				if ret{
					for k,v := range visited{
						record[k] = v
					}
					fill_x(board,i,j,rows,columns)
				}
			}
		}
	}
}

func dfs_solve(board [][]byte,rows int,columns int,r int,c int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if board[r][c] == 'X' || board[r][c] == 'Y'{
		return
	}
	board[r][c] = 'Y'
	dfs_solve(board,rows,columns,r - 1,c)
	dfs_solve(board,rows,columns,r + 1,c)
	dfs_solve(board,rows,columns,r,c - 1)
	dfs_solve(board,rows,columns,r,c + 1)
}

func solve2(board [][]byte){
	var rows int = len(board)
	var columns int = len(board[0])
	if rows <= 2 || columns <= 2{
		return
	}
	for i := 0;i < rows;i++{
		dfs_solve(board,rows,columns,i,0)
		dfs_solve(board,rows,columns,i,columns - 1)
	}
	for j := 0;j < columns;j++{
		dfs_solve(board,rows,columns,0,j)
		dfs_solve(board,rows,columns,rows - 1,j)
	}
	for i := 1;i < rows - 1;i++{
		for j := 1;j < columns - 1;j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			}else if board[i][j] == 'Y'{
				board[i][j] = 'O'
			}
		}
	}
	for i := 0;i < rows;i++{
		if board[i][0] == 'Y'{
			board[i][0] = 'O'
		}
		if board[i][columns - 1] == 'Y'{
			board[i][columns - 1] = 'O'
		}
	}
	for j := 0;j < columns;j++{
		if board[0][j] == 'Y'{
			board[0][j] = 'O'
		}
		if board[rows - 1][j] == 'Y'{
			board[rows - 1][j] = 'O'
		}
	}
}