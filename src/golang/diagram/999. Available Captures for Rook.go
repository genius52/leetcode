package diagram

func NumRookCaptures(board [][]byte) int {
	var rows int = len(board)
	var columns int = len(board[0])
	var rook_row int = 0
	var rook_column int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if(board[i][j] == 'R'){
				rook_row = i
				rook_column = j
				break
			}
		}
	}
	var res int = 0
	var r int = rook_row - 1
	for(r >= 0){
		if(board[r][rook_column] == 'B'){
			break
		}
		if(board[r][rook_column] == 'p'){
			res++
			break
		}
		r--
	}
	r = rook_row + 1
	for(r < rows){
		if(board[r][rook_column] == 'B'){
			break
		}
		if(board[r][rook_column] == 'p'){
			res++
			break
		}
		r++
	}
	var c int = rook_column - 1
	for(c >= 0){
		if(board[rook_row][c] == 'B'){
			break
		}
		if(board[rook_row][c] == 'p'){
			res++
			break
		}
		c--
	}
	c = rook_column + 1
	for(c < columns){
		if(board[rook_row][c] == 'B'){
			break
		}
		if(board[rook_row][c] == 'p'){
			res++
			break
		}
		c++
	}
	return res
}