package array

//79
//board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//Given word = "ABCCED", return true.
//Given word = "SEE", return true.
//Given word = "ABCB", return false.
// 1: up 2: right 3:down 4:left
func dfs_exist(board [][]byte,word string,pos int,r int,c int,rows int,columns int)bool{
	if pos >= len(word){
		return true
	}
	if r < 0 || c < 0 || r >= rows || c >= columns || board[r][c] == 0{
		return false
	}
	if word[pos] != board[r][c]{
		return false
	}
	tmp := board[r][c]
	board[r][c] = 0
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		res := dfs_exist(board,word,pos + 1,next_r,next_c,rows,columns)
		if res{
			return true
		}
	}
	board[r][c] = tmp
	return false
}

func Exist(board [][]byte, word string) bool {
	var rows int = len(board)
	var columns int = len(board[0])
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			ret := dfs_exist(board,word,0,i,j,rows,columns)
			if ret{
				return true
			}
		}
	}
	return false
}
