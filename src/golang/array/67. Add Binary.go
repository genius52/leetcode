package array

func addBinary(a string, b string) string {
	var len_a int = len(a)
	var len_b int = len(b)
	var pos_a int = len_a - 1
	var pos_b int = len_b - 1
	var res string
	var upgrade bool = false
	for pos_a >= 0 && pos_b >= 0{
		if a[pos_a] == '1' && b[pos_b] == '1'{
			if upgrade{
				res = "1" + res
			}else{
				res = "0" + res
				upgrade = true
			}
		}else if a[pos_a] == '1' || b[pos_b] == '1'{
			if upgrade{
				res = "0" + res
			}else{
				res = "1" + res
			}
		}else{
			if upgrade{
				res = "1" + res
				upgrade = false
			}else{
				res = "0" + res
			}
		}
		pos_a--
		pos_b--
	}
	for pos_a >= 0{
		if a[pos_a] == '0'{
			if upgrade{
				res = "1" + res
				upgrade = false
			}else{
				res = "0" + res
			}
		}else{
			if upgrade{
				res = "0" + res
			}else{
				res = "1" + res
			}
		}
		pos_a--
	}
	for pos_b >= 0{
		if b[pos_b] == '0'{
			if upgrade{
				res = "1" + res
				upgrade = false
			}else{
				res = "0" + res
			}
		}else{
			if upgrade{
				res = "0" + res
			}else{
				res = "1" + res
			}
		}
		pos_b--
	}
	if upgrade {
		res = "1" + res
	}
	return res
}