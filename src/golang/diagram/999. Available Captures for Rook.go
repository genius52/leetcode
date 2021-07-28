package diagram

func numRookCaptures(board [][]byte) int {
	var rows int = len(board)
	var columns int = len(board[0])
	var start_row int = 0
	var start_col int = 0
	for i := 0;i < rows;i++{
loop:
		for j := 0;j < columns;j++{
			if board[i][j] == 'R'{
				start_row = i
				start_col = j
				break loop
			}
		}
	}
	var res int = 0
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		i := start_row
		j := start_col
		for i >= 0 && i < rows && j >= 0 && j < columns{
			if board[i][j] == 'B'{
				break
			}
			if board[i][j] == 'p'{
				res++
				break
			}
			i += dir[0]
			j += dir[1]
		}
	}
	return res
}

//func NumRookCaptures(board [][]byte) int {
//	var rows int = len(board)
//	var columns int = len(board[0])
//	var rook_row int = 0
//	var rook_column int = 0
//	for i := 0;i < rows;i++{
//		for j := 0;j < columns;j++{
//			if(board[i][j] == 'R'){
//				rook_row = i
//				rook_column = j
//				break
//			}
//		}
//	}
//	var res int = 0
//	var r int = rook_row - 1
//	for(r >= 0){
//		if(board[r][rook_column] == 'B'){
//			break
//		}
//		if(board[r][rook_column] == 'p'){
//			res++
//			break
//		}
//		r--
//	}
//	r = rook_row + 1
//	for(r < rows){
//		if(board[r][rook_column] == 'B'){
//			break
//		}
//		if(board[r][rook_column] == 'p'){
//			res++
//			break
//		}
//		r++
//	}
//	var c int = rook_column - 1
//	for(c >= 0){
//		if(board[rook_row][c] == 'B'){
//			break
//		}
//		if(board[rook_row][c] == 'p'){
//			res++
//			break
//		}
//		c--
//	}
//	c = rook_column + 1
//	for(c < columns){
//		if(board[rook_row][c] == 'B'){
//			break
//		}
//		if(board[rook_row][c] == 'p'){
//			res++
//			break
//		}
//		c++
//	}
//	return res
//}