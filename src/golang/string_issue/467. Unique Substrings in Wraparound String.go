package string_issue

import "math"

func JC(num int) int{
	var count int = 1;
	if (num == 1 || num == 0) {
		return 1;
	}
	for i := 1; i <= num;i++ {
		count = count*i;
	}
	return count;
}

func FindSubstringInWraproundString(p string) int {
	var alpha []int = make([]int,26)
	var l int = len(p)
	if(l <= 1){
		return l
	}
	var begin int = 0
	var end int = begin + 1
	alpha[p[0] - 'a'] = 1
	for end < l{
		if p[end] - p[end - 1] == 1 || p[end - 1] - p[end] == 25{
			alpha[p[end] - 'a'] = int(math.Max(float64(alpha[p[end] - 'a']),float64(end - begin + 1)))
		}else{
			alpha[p[end] - 'a'] = int(math.Max(float64(alpha[p[end] - 'a']),1))
			begin = end
		}
		end++
	}
	var res int = 0
	for i := 0;i < 26;i++{
		res += alpha[i]
	}
	return res
}