package diagram

func check_placeWordInCrossword(board *[][]byte,rows int,columns,r int,c int,step []int,word string,l int)bool{
	var idx int = 0
	for idx < l && r >= 0 && r < rows && c >= 0 && c < columns{
		if (*board)[r][c] != ' ' && (*board)[r][c] != word[idx]{
			return false
		}
		idx++
		r += step[0]
		c += step[1]
	}
	return idx == l
}

func PlaceWordInCrossword(board [][]byte, word string) bool {
	var l int = len(word)
	var word2 string
	for i := 0;i < l;i++{
		word2 = string(word[i]) + word2
	}
	var rows int = len(board)
	var columns int = len(board[0])
	for i := 0;i < rows;i++{
		for j := 0;j <= (columns - l);{
			if board[i][j] == '#'{
				j++
				continue
			}
			var space_len int = 0
			var move int = j
			for move < columns && board[i][move] != '#' {
				move++
				space_len++
			}
			if space_len == l{
				res := check_placeWordInCrossword(&board,rows,columns,i,j,[]int{0,1},word,l)
				if res{
					return true
				}
				res = check_placeWordInCrossword(&board,rows,columns,i,j,[]int{0,1},word2,l)
				if res{
					return true
				}
			}
			j = move
		}
	}
	for j := 0;j < columns;j++{
		for i := 0;i <= (rows - l);{
			if board[i][j] == '#'{
				i++
				continue
			}
			var space_len int = 0
			var move int = i
			for move < rows && board[move][j] != '#' {
				move++
				space_len++
			}
			if space_len == l{
				res := check_placeWordInCrossword(&board,rows,columns,i,j,[]int{1,0},word,l)
				if res{
					return true
				}
				res = check_placeWordInCrossword(&board,rows,columns,i,j,[]int{1,0},word2,l)
				if res{
					return true
				}
			}
			i = move
		}
	}
	return false
}