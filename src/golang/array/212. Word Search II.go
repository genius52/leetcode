package array


//212
//Input:
//board = [
//  ['o','a','a','n'],
//  ['e','t','a','e'],
//  ['i','h','k','r'],
//  ['i','f','l','v']
//]
//words = ["oath","pea","eat","rain"]
//
//Output: ["eat","oath"]
func dfs_findWords(board [][]byte,i int,j int,s string,cur_pos int,visited [][]bool)bool{
	rows := len(board)
	columns := len(board[0])
	if i < 0 || i >= rows || j < 0 || j >= columns || visited[i][j]{
		return false
	}
	if cur_pos >= len(s){
		return true
	}
	visited[i][j] = true
	var res bool = false
	if board[i][j] == s[cur_pos]{
		if cur_pos == len(s) - 1{
			res = true
		}else{
			res = dfs_findWords(board,i - 1,j,s,cur_pos+1,visited) || dfs_findWords(board,i + 1,j,s,cur_pos+1,visited)||
				dfs_findWords(board,i,j - 1,s,cur_pos+1,visited) || dfs_findWords(board,i,j + 1,s,cur_pos+1,visited)
		}
	}
	visited[i][j] = false
	return res
}

func findWords2(board [][]byte, words []string) []string {
	rows := len(board)
	columns := len(board[0])
	var res []string
	var record map[string]bool = make(map[string]bool)
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			for k := 0;k < len(words);k++{
				if len(words[k]) == 0{
					continue
				}
				if _,ok := record[words[k]];ok{
					continue
				}
				var visited [][]bool = make([][]bool,rows)
				for c := 0;c < rows;c++{
					visited[c] = make([]bool,columns)
				}

				if dfs_findWords(board,i,j,words[k],0,visited){
					record[words[k]] = true
					res = append(res,words[k])
				}
			}
		}
	}
	return res
}