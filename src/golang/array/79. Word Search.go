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
func dfs_exist(board [][]byte,word string,pos int,r int,c int,rows int,columns int,visited [][]bool)bool{
	if pos >= len(word){
		return true
	}
	if r < 0 || c < 0 || r >= rows || c >= columns || visited[r][c]{
		return false
	}
	if word[pos] != board[r][c]{
		return false
	}
	visited[r][c] = true
	var res bool = false
	res = dfs_exist(board,word,pos + 1,r - 1,c,rows,columns,visited)
	if !res{
		res = dfs_exist(board,word,pos + 1,r + 1,c,rows,columns,visited)
	}
	if !res {
		res = dfs_exist(board, word, pos+1, r, c-1, rows, columns, visited)
	}
	if !res {
		res = dfs_exist(board, word, pos+1, r, c+1, rows, columns, visited)
	}
	visited[r][c] = false
	return res
}

func Exist(board [][]byte, word string) bool {
	var rows int = len(board)
	var columns int = len(board[0])
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		visited[i] = make([]bool,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			ret := dfs_exist(board,word,0,i,j,rows,columns,visited)
			if ret{
				return true
			}
		}
	}
	return false
}
