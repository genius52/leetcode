package number

import (
	"sort"
	"strconv"
)

func cmp_largestMultipleOfThree(s1,s2 string)bool{
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 > l2{
		return true
	}
	for i := 0;i < l1;i++{
		if s1[i] > s2[i]{
			return true
		}
	}
	return false
}
func LargestMultipleOfThree(digits []int) string {
	var l int = len(digits)
	sort.Ints(digits)
	if digits[l - 1] == 0{
		return "0"
	}
	var pre [3]string//record[i] 余数是i的最大字符串
	pre[digits[0] % 3] = strconv.Itoa(digits[0])
	for i := 1;i < l;i++{
		mod := digits[i] % 3
		var cur [3]string
		for j := 0;j <= 2;j++{//当前字符 + 之前余数的字符
			if len(pre[j]) == 0{ //如果之前余数字符为空
				continue
			}
			next := (j + mod) % 3
			cur[next] = strconv.Itoa(digits[i]) + pre[j]
		}
		if len(cur[mod]) == 0{
			cur[mod] = strconv.Itoa(digits[i])
		}
		for j := 0;j <= 2;j++{
			ret := cmp_largestMultipleOfThree(pre[j],cur[j])
			if !ret{
				pre[j] = cur[j]
			}
		}
	}
	return pre[0]
}