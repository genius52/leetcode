package number

import "strconv"

func MinMaxDifference(num int) int {
	s := strconv.Itoa(num)
	var l int = len(s)
	var max_num int = 0
	i := 0
	for ; i < l; i++ {
		max_num *= 10
		c := int(s[i] - '0')
		if c != 9 {
			break
		}
		max_num += c
	}
	var max_digit int = -1
	for ; i < l; i++ {
		max_num *= 10
		c := int(s[i] - '0')
		if max_digit != -1 {
			if c == max_digit {
				max_num += 9
			} else {
				max_num += c
			}
		} else {
			max_digit = c
			max_num += 9
		}
	}
	var min_num int = 0
	var min_digit int = -1
	for j := 0; j < l; j++ {
		min_num *= 10
		c := int(s[j] - '0')
		if min_digit != -1 {
			if c == min_digit {
				min_num += 0
			} else {
				min_num += c
			}
		} else {
			min_digit = c
			min_num += 0
		}
	}
	return max_num - min_num
}
