package string_issue

import (
	"math"
	"strconv"
)


func findDifferentBinaryString(nums []string) string {
	var l int = len(nums)
	var str_l int = len(nums[0])
	var exist map[string]bool = make(map[string]bool)
	for i := 0;i < l;i++{
		exist[nums[i]] = true
	}
	high := int(math.Pow(2,float64(str_l))) - 1
	for i := 0;i <= high;i++{
		s := convertToBin(i,str_l)
		if _,ok := exist[s];!ok{
			return s
		}
	}
	return ""
}

func convertToBin(num int,target_len int) string {
	s := ""
	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	var cur_len int = len(s)
	for i := 0;i < (target_len - cur_len);i++{
		s = "0" + s
	}
	return s
}
