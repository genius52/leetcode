package string_issue

func try_palindromic(s string,start int,end int) int{
	var cnt int = 0
	for start >= 0 && end < len(s){
		if s[start] == s[end]{
			cnt++
		}else{
			break
		}
		start--
		end++
	}
	return cnt
}

func countSubstrings(s string) int {
	var res int = 0
	for start := 0;start < len(s);start++{
		res += try_palindromic(s,start,start)
		res += try_palindromic(s,start,start+1)
	}
	return res
}