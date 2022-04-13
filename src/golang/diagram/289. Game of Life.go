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
//func get_status(board [][]int,i,j int)int{
//	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
//		return 0
//	}
//	return board[i][j]
//}
//
//func gameOfLife(board [][]int)  {
//	var neighbour_cnt [][]int = make([][]int,len(board))
//	for i := 0;i < len(board);i++{
//		neighbour_cnt[i] = make([]int,len(board[0]))
//		for j := 0;j < len(board[0]);j++{
//			neighbour_cnt[i][j] = get_status(board,i - 1,j - 1) + get_status(board,i - 1,j) + get_status(board,i - 1,j + 1) + get_status(board,i,j - 1) + get_status(board,i,j + 1)+ get_status(board,i + 1,j - 1) + get_status(board,i + 1,j) + get_status(board,i + 1,j + 1)
//		}
//	}
//	for i := 0;i < len(board);i++{
//		for j := 0;j < len(board[0]);j++{
//			if board[i][j] == 1{
//				if neighbour_cnt[i][j] < 2 || neighbour_cnt[i][j] > 3{
//					board[i][j] = 0
//				}
//			}else{
//				if neighbour_cnt[i][j] == 3{
//					board[i][j] = 1
//				}
//			}
//		}
//	}
//}

func gameOfLife(board [][]int){
	var rows int = len(board)
	var columns int = len(board[0])
	var record [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		record[i] = make([]int,columns)
	}
	var dirs [][]int = [][]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if board[i][j] == 0{
				continue
			}
			for _,dir := range dirs{
				next_r := i + dir[0]
				next_c := j + dir[1]
				if next_r >= 0 && next_r < rows && next_c >= 0 && next_c < columns{
					record[next_r][next_c]++
				}
			}
		}
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if board[i][j] == 1{
				if record[i][j] < 2 || record[i][j] > 3{
					board[i][j] = 0
				}
			}else{
				if record[i][j] == 3{
					board[i][j] = 1
				}
			}
		}
	}
}