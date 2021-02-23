package string_issue

import "strings"

func isPalindrome(s string) bool {
	ls := strings.ToLower(s)
	begin,end := 0,len(ls) - 1
	for begin < end{
		for begin < end{
			if (ls[begin] >= 'a' && ls[begin] <= 'z') || (ls[begin] >= '0' && ls[begin] <= '9'){
				break
			}else{
				begin++
			}
		}
		for end > begin{
			if (ls[end] >= 'a' && ls[end] <= 'z') || (ls[end] >= '0' && ls[end] <= '9'){
				break
			}else{
				end--
			}
		}
		if begin < end{
			if ls[begin] != ls[end]{
				return false
			}
			begin++
			end--
		}
	}
	return true
}