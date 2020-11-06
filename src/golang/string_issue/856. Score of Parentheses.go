package string_issue

//import "container/list"

func rec_scoreOfParentheses(S string,l int,pos int)(int,int){
	if pos >= l{
		return 0,pos
	}
	if S[pos] == ')'{
		return 1,pos
	}
	var res int = 0
	for pos < l{
		if S[pos] == '('{
			if S[pos + 1] == ')'{
				res++
				pos = pos + 2
			}else{
				sub_res,sub_pos := rec_scoreOfParentheses(S,l,pos + 1)
				res += 2 * sub_res
				pos = sub_pos + 1
			}
		}else{
			return res,pos
		}
	}
	return res,pos
}

func ScoreOfParentheses(S string) int {
	var l int = len(S)
	res,_ := rec_scoreOfParentheses(S,l,0)
	return res
}