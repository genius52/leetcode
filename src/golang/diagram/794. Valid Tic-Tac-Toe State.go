package diagram

func ValidTicTacToe(board []string) bool {
	var x_cnt int = 0
	var o_cnt int = 0
	for i := 0;i < 3;i++{
		for j := 0;j < 3;j++{
			if board[i][j] == 'X'{
				x_cnt++
			}else if board[i][j] == 'O'{
				o_cnt++
			}
		}
	}
	if o_cnt > x_cnt{
		return false
	}
	if x_cnt > (o_cnt + 1){
		return false
	}
	var x_lines int = 0
	var o_lines int = 0
	for i := 0;i < 3;i++{
		if board[i][0] == 'X' && board[i][1] == 'X' && board[i][2] == 'X'{
			x_lines++
		}
		if board[i][0] == 'O' && board[i][1] == 'O' && board[i][2] == 'O'{
			o_lines++
		}
		if board[0][i] == 'X' && board[1][i] == 'X' && board[2][i] == 'X'{
			x_lines++
		}
		if board[0][i] == 'O' && board[1][i] == 'O' && board[2][i] == 'O'{
			o_lines++
		}
	}
	if x_lines > 0 && o_lines > 0{
		return false
	}
	if x_lines > 2 || o_lines > 2{
		return false
	}
	if x_lines > o_lines{
		if x_cnt == o_cnt{
			return false
		}
	}
	if x_lines < o_lines{
		if x_cnt > o_cnt{
			return false
		}
	}
	var x_tilt_line int = 0
	var o_tilt_line int = 0
	if board[0][0] == 'X' && board[1][1] == 'X' && board[2][2] == 'X'{
		x_tilt_line++
	}
	if board[0][0] == 'O' && board[1][1] == 'O' && board[2][2] == 'O'{
		o_tilt_line++
	}
	if board[2][0] == 'X' && board[1][1] == 'X' && board[0][2] == 'X'{
		x_tilt_line++
	}
	if board[2][0] == 'O' && board[1][1] == 'O' && board[0][2] == 'O'{
		o_tilt_line++
	}

	if x_tilt_line > o_tilt_line {
		if x_cnt == o_cnt{
			return false
		}
	}
	if x_tilt_line < o_tilt_line{
		if x_cnt > o_cnt{
			return false
		}
	}
	return true
}