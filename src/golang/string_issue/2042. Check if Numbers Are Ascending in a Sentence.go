package string_issue

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool{
	var record []string = strings.Split(s," ")
	var l int = len(record)
	var pre_num int = -1
	for i := 0;i < l;i++{
		n,err := strconv.Atoi(record[i])
		if err != nil{
			continue
		}
		if n <= pre_num{
			return false
		}
		pre_num = n
	}
	return true
}

func AreNumbersAscending(s string) bool {
	var record []string = strings.Split(s," ")
	var l int = len(record)
	var left int = 0
	for left < l{
		var left_num int = -1
		for left < l{
			n,err := strconv.Atoi(record[left])
			if err == nil{
				left_num = n
				break
			}
			left++
		}
		var right_num int = -1
		var right int = left + 1
		for right < l{
			n,err := strconv.Atoi(record[right])
			if err == nil{
				right_num = n
				break
			}
			right++
		}
		if right >= l{
			break
		}
		if left_num >= right_num{
			return false
		}
		left = right
	}
	return true
}