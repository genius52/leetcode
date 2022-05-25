package string_issue

import "container/list"

//Input: "(()"
//Output: 2
//Explanation: The longest valid parentheses substring is "()"
//Input: ")()())"
//Output: 4
//Explanation: The longest valid parentheses substring is "()()"
func longestValidParentheses(s string) int{
	l := len(s)
	var stack list.List
	var max int = 0
	var start int = 0
	for i := 0;i < l;i++{
		if s[i] == ')' {
			if stack.Len() == 0{
				start = i + 1
			}else{
				top := stack.Back()
				stack.Remove(top)
				if stack.Len() == 0{//)()()()()
					max = max_int(i - start + 1,max)//start will not move,if string is valid
				}else{
					last_left := stack.Back().Value.(int)
					max = max_int(i - last_left,max)//(  () ()  length = i - last_open_ position
				}
			}
		}else{
			stack.PushBack(i)
		}
	}
	return max
}

func longestValidParentheses2(s string) int{
	l := len(s)
	if l <= 1{
		return 0
	}
	var dp []int = make([]int,l)
	dp[0] = 0
	var max int = 0
	for i := 1;i < l;i++{
		if s[i] == ')'{
			head := i - dp[i-1] - 1
			if head >= 0 && s[head] == '('{
				dp[i] = dp[i - 1] + 2
				if i - dp[i-1] - 2 > 0{
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > max{
				max = dp[i]
			}
		}
	}
	return max
}