package string_issue

//"t", evaluating to True;
//"f", evaluating to False;
//"!(expr)", evaluating to the logical NOT of the inner expression expr;
//"&(expr1,expr2,...)", evaluating to the logical AND of 2 or more inner expressions expr1, expr2, ...;
//"|(expr1,expr2,...)", evaluating to the logical OR of 2 or more inner expressions expr1, expr2, ...

//Input: expression = "|(&(t,f,t),!(t))"
//Output: false
func helper_parseBoolExpr(s string,start int,end int,is_or bool)bool{
	var res bool = true
	if is_or{
		res = false
	}
	var tags int = 0
	var left int = 0
	//var l int = len(s)
	for i := start;i < end;i++{
		if s[i] == '('{
			if tags == 0{
				left = i
			}
			tags++
		}else if s[i] == ')'{
			tags--
			if tags == 0{
				sub_is_or := false
				if left > start{
					if s[left - 1] == '|'{
						sub_is_or = true
					}
				}
				sub := helper_parseBoolExpr(s,left + 1,i,sub_is_or)
				if left > start{
					if s[left - 1] == '!'{
						sub = !sub
					}
				}
				if is_or{
					res = res || sub
				}else{
					res = res && sub
				}
			}
		}else if s[i] == 't'{
			if tags == 0{
				if is_or{
					res = res || true
				}else{
					res = res && true
				}
			}
		}else if s[i] == 'f'{
			if tags == 0{
				if is_or{
					res = res || false
				}else{
					res = res && false
				}
			}
		}
	}
	return res
}


func ParseBoolExpr(expression string) bool {
	var l int = len(expression)
	if expression[0] == '!'{
		return !helper_parseBoolExpr(expression,2,l - 1,false)
	}else if expression[0] == '|'{
		return helper_parseBoolExpr(expression,2,l - 1,true)
	}else if expression[0] == '&'{
		return helper_parseBoolExpr(expression,2,l - 1,false)
	}else{
		return helper_parseBoolExpr(expression,1,l - 1,false)
	}
}