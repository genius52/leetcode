package string_issue

import "strings"

//func jieCheng(n int) int {
//	result := 1
//	for i := 2; i <= n; i++ {
//		result *= i
//		result %= 1000000007
//	}
//	return result
//}

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
			//if strings.HasSuffix(target,sub){
			//	if cnt2,ok2 := record[target[:target_len - sub_len]];ok2{
			//		res += cnt * cnt2
			//	}
			//}
		}
	}
	return res
}