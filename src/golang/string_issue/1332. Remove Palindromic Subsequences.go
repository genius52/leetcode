package string_issue

func is_palindromic(s string)bool{
	l := len(s)
	begin := 0
	end := l - 1
	for begin < end{
		if s[begin] != s[end]{
			return false;
		}
	}
	return true
}
func removePalindromeSub(s string) int {
	if len(s) == 0{
		return 0
	}
	if is_palindromic(s){
		return 1
	}
	return 2//step1: remove all 'a',step2: remove all 'b'
}