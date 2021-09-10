package diagram


func Tictactoe(moves [][]int) string{
	var a [8]int//[0-2]row [3-5]column [6-7]tilt
	var b [8]int
	var l int = len(moves)
	for i := 0;i < l;i++{
		r := moves[i][0]
		c := moves[i][1]
		if (i | 1) == i{
			b[r]++
			b[c + 3]++
			if r == 1 && c == 1{
				b[6]++
				b[7]++
			}else if r == c{
				b[6]++
			}else if (r + c) == 2{
				b[7]++
			}
		}else{
			a[r]++
			a[c + 3]++
			if r == 1 && c == 1{
				a[6]++
				a[7]++
			}else if r == c{
				a[6]++
			}else if (r + c) == 2{
				a[7]++
			}
		}
	}
	for i := 0;i < 8;i++{
		if a[i] == 3{
			return "A"
		}
		if b[i] == 3{
			return "B"
		}
	}
	if l == 9{
		return "Draw"
	}
	return "Pending"
}