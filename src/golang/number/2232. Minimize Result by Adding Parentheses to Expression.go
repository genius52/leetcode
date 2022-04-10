package number

import "strconv"

func MinimizeResult(expression string) string {
	var l int = len(expression)
	var plus_idx int = 0
	for i := 0;i < l;i++{
		if expression[i] == '+'{
			plus_idx = i
			break
		}
	}
	var left_exp []string
	for i := 0;i < plus_idx;i++{
		cur := expression[:i] + "(" + expression[i:plus_idx]
		left_exp = append(left_exp,cur)
	}
	var right_exp []string
	for i := plus_idx + 2;i <= l;i++{
		cur := expression[plus_idx:i] + ")" + expression[i:]
		right_exp = append(right_exp,cur)
	}
	var min_val int = 2147483647
	var min_exp string
	for i := 0;i < len(left_exp);i++{
		for j := 0;j < len(right_exp);j++{
			exp := left_exp[i] + right_exp[j]
			var k1 int = 0
			for k1 < len(exp){
				if exp[k1] == '('{
					break
				}
				k1++
			}
			var left_val int = 1
			var err error = nil
			left_val,err = strconv.Atoi(exp[:k1])
			if err != nil{
				left_val = 1
			}
			var k2 int = k1
			for k2 < len(exp){
				if exp[k2] == '+'{
					break
				}
				k2++
			}
			var mid_val1 int = 0
			//err = nil
			mid_val1,_ = strconv.Atoi(exp[k1 + 1:k2])
			var k3 int = k2
			for k3 < len(exp){
				if exp[k3] == ')'{
					break
				}
				k3++
			}
			var mid_val2 int = 0
			mid_val2,_ = strconv.Atoi(exp[k2 + 1:k3])
			var right_val int = 0
			right_val,err = strconv.Atoi(exp[k3 + 1:])
			if err != nil{
				right_val = 1
			}
			cur_val := left_val * (mid_val1 + mid_val2) * right_val
			if cur_val < min_val{
				min_val = cur_val
				min_exp = exp
			}
		}
	}
	return min_exp
}