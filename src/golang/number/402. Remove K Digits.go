package number

import (
	"strconv"
	"strings"
)

//402
func dp_removeKdigits(num string,cur_num string,target_len int,cur_pos int)int{
	if len(cur_num) == target_len{
		val,_ := strconv.Atoi(cur_num)
		return val
	}
	if cur_pos >= len(num){
		val,_ := strconv.Atoi(num)
		return val
	}
	n1 := dp_removeKdigits(num,cur_num,target_len,cur_pos + 1)
	n2 := dp_removeKdigits(num,cur_num + string(num[cur_pos]),target_len,cur_pos + 1)
	return min_int(n1,n2)
}

func removeKdigits2(num string, k int) string {
	if len(num) <= k{
		return "0"
	}
	res := strconv.Itoa(dp_removeKdigits(num,"",len(num) - k,0))
	return res
}

func removeKdigits(num string, k int) string {
	if len(num) <= k{
		return "0"
	}
	var remove_cnt int = 0
	for remove_cnt < k && len(num) > 0{
		var delete bool = false
		for i := 0;i < len(num) - 1;i++{
			if num[i] > num[i+1]{
				num = num[:i] + num[i+1:]
				delete = true
				break;
			}
		}
		if !delete{
			num = num[:len(num)-1]
		}
		remove_cnt++
	}
	res := strings.TrimLeft(num,"0")
	if len(res) == 0{
		return "0"
	}
	return res
}