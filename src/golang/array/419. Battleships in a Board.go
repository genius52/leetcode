package array


//419
//var total int = 0
//func dfs_419(board [][]byte,x int,y int,last_is_ship bool) bool{
//	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0]) - 1{
//		return false
//	}
//	if board[x][y] == 'X'{
//
//	}
//
//	dfs_419(board,x + 1,y,board[x][y] == 'X')
//	dfs_419(board,x ,y + 1,board[x][y] == 'X')
//}
func point_is_ship(board [][]byte,x int,y int) bool{
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0]) - 1{
		return false
	}
	return board[x][y] == 'X'
}

func countBattleships(board [][]byte) int {
	var res int
	for i := 0; i < len(board);i++{
		for j := 0; j < len(board[0]);j++{
			if board[i][j] == 'X' && !point_is_ship(board,i+1,j) && !point_is_ship(board,i,j+1){
				res++
			}
		}
	}
	return res
}