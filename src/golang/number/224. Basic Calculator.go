package number

import (
	"strconv"
	"strings"
)

//s consists of digits, '+', '-', '(', ')', and ' '.
//Input: s = "(1+(4+5+2)-3)+(6+8)"
//Output: 23
func recursive_calculate224(s string,l int,start int)(int,int){
	if start >= l{
		return 0,l - 1
	}
	var total int = 0
	var positive bool = true
	var i int = start
	var num_begin int = start
	for i < l{
		if s[i] == '('{
			val,pos := recursive_calculate224(s,l,i + 1)
			if positive{
				total += val
			}else{
				total -= val
			}
			i = pos
		}else if s[i] == '+'|| s[i] == '-'|| s[i] == ')'{
			n,err := strconv.Atoi(s[num_begin:i])
			if err == nil{
				if positive{
					total += n
				}else{
					total -= n
				}
			}
			if s[i] == '+'{
				positive = true
			}else if s[i] == '-'{
				positive = false
			}else if s[i] == ')'{
				break
			}
			i++
			num_begin = i
		}else{
			if i == l - 1{
				n,err := strconv.Atoi(s[num_begin:i + 1])
				if err == nil{
					if positive{
						total += n
					}else{
						total -= n
					}
				}
				break
			}
			i++
		}
	}
	return total,i + 1
}

func Calculate224(s string) int {
	s = strings.ReplaceAll(s," ","")
	var l int = len(s)
	val,_ := recursive_calculate224(s,l,0)
	return val
}