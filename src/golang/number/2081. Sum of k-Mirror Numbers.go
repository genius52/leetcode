package number

import (
	"strconv"
)

func match_k(k int,n int)bool{
	var s string
	for n > 0{
		mod := n % k
		s = s + strconv.Itoa(mod)
		n = n / k
	}
	var left int = 0
	var right int = len(s) - 1
	for left < right{
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

func get_next_num(n int)int{
	s := strconv.Itoa(n)
	var l int = len(s)
	var data []uint8 = make([]uint8,l)
	for i := 0;i < l;i++{
		data[i] = s[i]
	}
	mid := (l - 1) / 2
	var find bool = false
	for i := mid;i >= 0;i--{
		if data[i] != '9'{
			add_one := data[i] + 1
			data[i] = add_one
			data[l - 1 - i] = add_one
			find = true
			break
		}else{
			data[i] = '0'
			data[l - 1 - i] = '0'
		}
	}
	if find{
		var next string
		for i := 0;i < l;i++{
			next += string(data[i])
		}
		ret,_ :=  strconv.Atoi(next)
		return ret
	}else{
		var next string
		next += "1"
		for i := 1;i < l;i++{
			next += "0"
		}
		next += "1"
		ret,_ :=  strconv.Atoi(next)
		return ret
	}
}

func KMirror(k int, n int) int64 {
	var res int64 = 0
	cur_num := 0
	for n > 0{
		next := get_next_num(cur_num)
		if match_k(k,next){
			res += int64(next)
			n--
		}
		cur_num = next
	}
	return res
}