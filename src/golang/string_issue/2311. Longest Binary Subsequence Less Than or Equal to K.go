package string_issue

func LongestSubsequence(s string, k int) int {
	var l int = len(s)
	var res int = 0
	var val int64 = 0
	var add int64 = 1
	for i := l - 1;i >= 0;i--{
		if s[i] == '1'{
			if add > 0 && add <= 2147483647 && (val + add) <= int64(k){
				res++
				val += add
			}
		}else{
			res++
		}
		add *= 2
	}
	return res
}