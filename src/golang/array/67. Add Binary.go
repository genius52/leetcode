package array

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var len_a int = len(a)
	var len_b int = len(b)
	var pos_a int = len_a - 1
	var pos_b int = len_b - 1
	var res string
	var upgrade bool = false
	for pos_a >= 0 || pos_b >= 0 || upgrade{
		if pos_a >= 0 && pos_b >= 0{
			val := a[pos_a] - '0' + b[pos_b] - '0'
			if upgrade{
				val++
			}
			if val >= 2{
				upgrade = true
				val %= 2
			}else{
				upgrade = false
			}
			res = strconv.Itoa(int(val)) + res
		}else if pos_a >= 0{
			val := a[pos_a] - '0'
			if upgrade{
				val++
			}
			if val >= 2{
				val %= 2
				upgrade = true
			}else{
				upgrade = false
			}
			res = strconv.Itoa(int(val)) + res
		}else if pos_b >= 0{
			val := b[pos_b] - '0'
			if upgrade{
				val++
			}
			if val >= 2{
				val %= 2
				upgrade = true
			}else{
				upgrade = false
			}
			res = strconv.Itoa(int(val)) + res
		}else if upgrade{
			res = "1" + res
			upgrade = false
		}
		pos_a--
		pos_b--
	}
	return res
}