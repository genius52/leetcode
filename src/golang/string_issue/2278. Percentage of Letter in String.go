package string_issue

import "math"

func percentageLetter(s string, letter byte) int {
	var l int = len(s)
	var cnt int = 0
	for _,c := range s{
		if c == int32(letter){
			cnt++
		}
	}
	var res int = int(math.Floor(float64(cnt)/float64(l) * 100))
	return res
}