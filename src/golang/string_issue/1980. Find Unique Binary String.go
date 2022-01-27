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

type Tie_binary struct{
	next [2]*Tie_binary
	//val int
}

func dfs_findDifferentBinaryString(tie *Tie_binary,l int)string{
	if tie == nil{

	}
	if tie.next[0] == nil{
		var s string
		for i := 0;i < l;i++{
			s += "0"
		}
		return s
	}else if tie.next[1] == nil{
		var s string
		for i := 0;i < l;i++{
			s += "1"
		}
		return s
	}
	zero_s := dfs_findDifferentBinaryString(tie.next[0],l - 1)
	one_s := dfs_findDifferentBinaryString(tie.next[1],l - 1)
	if len(zero_s) != 0{
		return "0" + zero_s
	}else{
		return "1" + one_s
	}
}

func findDifferentBinaryString2(nums []string) string{
	var l int = len(nums[0])
	var root Tie_binary
	for _,num := range nums{
		var visit *Tie_binary = &root
		for i := 0;i < l;i++{
			if num[i] == '0'{
				if visit.next[0] == nil{
					visit.next[0] = new(Tie_binary)
				}
				visit = visit.next[0]
			}else{
				if visit.next[1] == nil{
					visit.next[1] = new(Tie_binary)
				}
				visit = visit.next[1]
			}
		}
	}
	res := dfs_findDifferentBinaryString(&root,l)
	return res
}