package diagram

//'U' moves our position up one row, if the position exists on the board;
//'D' moves our position down one row, if the position exists on the board;
//'L' moves our position left one column, if the position exists on the board;
//'R' moves our position right one column, if the position exists on the board;
//'!' adds the character board[r][c] at our current position (r, c) to the answer.
func alphabetBoardPath(target string) string {
	//char <-> position
	// row = (char - 'a') / 5
	// column = (char - 'a') % 5
	//var dirs []string = []string{"U","D","L","R"}
	var res string
	var pre_row int = 0
	var pre_col int = 0
	for _,c := range target{
		var cur_row int = int((c - 'a')/5)
		var cur_col int = int((c - 'a')%5)
		if cur_row == 5{
			var step int = 0
			var dir string
			if cur_col > pre_col{
				step = 1
				dir = "R"
			}else if cur_col < pre_col{
				step = -1
				dir = "L"
			}
			for c := pre_col;c != cur_col;c += step{
				res += dir
			}
			if cur_row > pre_row{
				step = 1
				dir = "D"
			}else if cur_row < pre_row{
				step = -1
				dir = "U"
			}
			for r := pre_row;r != cur_row;r += step{
				res += dir
			}
		}else{
			var step int = 0
			var dir string
			if cur_row > pre_row{
				step = 1
				dir = "D"
			}else if cur_row < pre_row{
				step = -1
				dir = "U"
			}
			for r := pre_row;r != cur_row;r += step{
				res += dir
			}
			step = 0
			if cur_col > pre_col{
				step = 1
				dir = "R"
			}else if cur_col < pre_col{
				step = -1
				dir = "L"
			}
			for c := pre_col;c != cur_col;c += step{
				res += dir
			}
		}
		res += "!"
		pre_row = cur_row
		pre_col = cur_col
	}
	return res
}