package string_issue

import "strings"

func NumOfPairs(nums []string, target string) int {
	var l int = len(nums)
	var record map[string]int = make(map[string]int)
	for i := 0;i < l;i++{
		record[nums[i]]++
	}
	var res int = 0
	//var target_len int = len(target)
	for sub,cnt := range record{
		if (sub + sub) == target{
			res += cnt * (cnt - 1)
		}else{
			var sub_len int = len(sub)
			if strings.HasPrefix(target,sub){
				if cnt2,ok2 := record[target[sub_len:]];ok2{
					res += cnt * cnt2
				}
			}
		}
	}
	return res
}

func numOfPairs(nums []string, target string) int{
	var prefix map[string]int = make(map[string]int)
	var suffix map[string]int = make(map[string]int)
	for _,n := range nums{
		if strings.HasPrefix(target,n){
			prefix[n]++
		}
		if strings.HasSuffix(target,n){
			suffix[n]++
		}
	}
	var res int = 0
	for k1,v1 := range prefix{
		need := target[len(k1):]
		if k1 == need{
			res += v1 * (v1 - 1)
		}else{
			if v2,ok := suffix[need];ok{
				res += v1 * v2
			}
		}
	}
	return res
}