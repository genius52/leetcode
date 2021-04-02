package string_issue

import "strconv"

func compress(chars []byte) int{
	var l int = len(chars)
	var res int = 0
	var left int = 0
	var index int = 0
	for left < l{
		var right int = left
		var c byte = chars[left]
		for right < l && chars[left] == chars[right]{
			right++
		}

		chars[index] = c
		index++
		res++
		cur_len := right - left
		if cur_len > 1{
			slen := strconv.Itoa(cur_len)
			for _,c := range slen{
				chars[index] = byte(c)
				index++
			}
			res += len(slen)
		}
		left = right
	}
	return res
}