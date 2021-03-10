package diagram

//Input:
//[
//  [0,1,0],
//  [0,0,1],
//  [1,1,1],
//  [0,0,0]
//]
//Output:
//[
//  [0,0,0],
//  [1,0,1],
//  [0,1,1],
//  [0,1,0]
//]
func get_status(board [][]int,i,j int)int{
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
		return 0
	}
	return board[i][j]
}

func gameOfLife(board [][]int)  {
	var neighbour_cnt [][]int = make([][]int,len(board))
	for i := 0;i < len(board);i++{
		neighbour_cnt[i] = make([]int,len(board[0]))
		for j := 0;j < len(board[0]);j++{
			neighbour_cnt[i][j] = get_status(board,i - 1,j - 1) + get_status(board,i - 1,j) + get_status(board,i - 1,j + 1) + get_status(board,i,j - 1) + get_status(board,i,j + 1)+ get_status(board,i + 1,j - 1) + get_status(board,i + 1,j) + get_status(board,i + 1,j + 1)
		}
	}
	for i := 0;i < len(board);i++{
		for j := 0;j < len(board[0]);j++{
			if board[i][j] == 1{
				if neighbour_cnt[i][j] < 2 || neighbour_cnt[i][j] > 3{
					board[i][j] = 0
				}
			}else{
				if neighbour_cnt[i][j] == 3{
					board[i][j] = 1
				}
			}
		}
	}
}