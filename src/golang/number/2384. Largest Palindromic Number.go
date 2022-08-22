package number

import (
	"strconv"
	"strings"
)

func LargestPalindromic(num string) string {
	var l int = len(num)
	var cnt [10]int
	for i := 0;i < l;i++{
		cnt[num[i] - '0']++
	}
	var max_odd int = -1
	for i := 9;i >= 0;i--{
		if (cnt[i] | 1) == cnt[i]{
			max_odd = i
			break
		}
	}
	var ss strings.Builder
	for i := 9;i >= 0;i--{
		if ss.Len() == 0 && i == 0{
			break
		}
		for j := 0;j < cnt[i]/2;j++{
			ss.WriteString(strconv.Itoa(i))
		}
	}
	var left string = ss.String()
	if max_odd != -1{
		ss.WriteString(strconv.Itoa(max_odd))
	}else{
		if len(left) == 0 && cnt[0] > 0{
			return "0"
		}
	}
	for i := len(left) - 1;i >= 0;i--{
		ss.WriteString(string(left[i]))
	}
	return ss.String()
}