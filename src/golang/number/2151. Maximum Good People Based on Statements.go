package number

func check_maximumGood(statements [][]int,l int,state int)bool{
	var cur_state []int = make([]int,l)
	for i := 0;i < l;i++{
		cur_state[i] = -1
	}
	for i := 0;i < l;i++{
		if (state | (1 << i)) == state{
			for j := 0;j < l;j++{
				if statements[i][j] == 0{ //bad
					if (state | (1 << j)) == state{
						return false
					}
					if cur_state[j] == 1{
						return false
					}
					cur_state[j] = 0
				}else if statements[i][j] == 1{ //good
					if (state | (1 << j)) != state{
						return false
					}
					if cur_state[j] == 0{
						return false
					}
					cur_state[j] = 1
				}
			}
		}
	}
	return true
}

func calc_cnt(n int)int{
	var cnt int = 0
	for n > 0{
		if n | 1 == n{
			cnt++
		}
		n = n >> 1
	}
	return cnt
}

func MaximumGood(statements [][]int) int {
	var l int = len(statements)
	var mask int = 1 << l
	var max_cnt int = 0
	for i := 1;i < mask;i++{
		ret := check_maximumGood(statements,l,i)
		if ret{
			var cnt int = 0
			n := i
			for n > 0{
				if n | 1 == n{
					cnt++
				}
				n = n >> 1
			}
			if cnt > max_cnt{
				max_cnt = cnt
			}
		}
	}
	return max_cnt
}