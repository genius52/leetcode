package diagram

//Input:
//
//[['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'M', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E']]
//
//Click : [3,0]
//
//Output:
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
func check_round(board [][]byte,rows int,columns int,row int,column int)bool{
	if row < 0 || column < 0 || row >= rows || column >= columns {
		return false
	}
	if board[row][column] == 'M'{
		return true
	}
	return false
}

func dfs_UpdateBoard(board [][]byte,rows int,columns int,row int,column int){
	if row < 0 || column < 0 || row >= rows || column >= columns{
		return
	}
	if board[row][column] == 'M'{
		board[row][column] = 'X'
	}else if board[row][column] == 'E'{
		board[row][column] = 'B'
		var mines int = 0
		if check_round(board,rows,columns,row - 1,column - 1){
			mines++
		}
		if check_round(board,rows,columns,row - 1,column){
			mines++
		}
		if check_round(board,rows,columns,row - 1,column + 1){
			mines++
		}
		if check_round(board,rows,columns,row,column - 1){
			mines++
		}
		if check_round(board,rows,columns,row,column + 1){
			mines++
		}
		if check_round(board,rows,columns,row + 1,column - 1){
			mines++
		}
		if check_round(board,rows,columns,row + 1,column){
			mines++
		}
		if check_round(board,rows,columns,row + 1,column + 1){
			mines++
		}
		if mines == 0{
			board[row][column] = 'B'
			dfs_UpdateBoard(board,rows,columns,row - 1,column - 1)
			dfs_UpdateBoard(board,rows,columns,row - 1,column)
			dfs_UpdateBoard(board,rows,columns,row - 1,column + 1)
			dfs_UpdateBoard(board,rows,columns,row,column - 1)
			dfs_UpdateBoard(board,rows,columns,row,column + 1)
			dfs_UpdateBoard(board,rows,columns,row + 1,column - 1)
			dfs_UpdateBoard(board,rows,columns,row + 1,column)
			dfs_UpdateBoard(board,rows,columns,row + 1,column + 1)
		}else{
			board[row][column] = byte(mines) + '0'
		}
	}
}

func UpdateBoard(board [][]byte, click []int) [][]byte {
	var rows int = len(board)
	var columns int = len(board[0])

	row := click[0]
	column := click[1]
	dfs_UpdateBoard(board,rows,columns,row,column)
	return board
}