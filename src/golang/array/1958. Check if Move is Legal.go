package array

func valid_checkMove(board [][]byte,rows int,columns int,start_x int,start_y int,dir []int,color byte)bool{
	next_r := start_x + dir[0]
	next_c := start_y + dir[1]
	var target_color byte = 'B'
	if color == 'B'{
		target_color = 'W'
	}
	var cnt int = 0
	for next_r >= 0 && next_r < rows && next_c >= 0 && next_c < columns &&
		board[next_r][next_c] == target_color{
		cnt++
		next_r += dir[0]
		next_c += dir[1]
	}
	if cnt == 0 || next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns || board[next_r][next_c] != color{
		return false
	}
	return true
}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	var rows int = len(board)
	var columns int = len(board[0])
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1},{-1,-1},{-1,1},{1,-1},{1,1}}
	for _,dir := range dirs{
		ret := valid_checkMove(board,rows,columns,rMove,cMove,dir,color)
		if ret{
			return true
		}
	}
	return false
}