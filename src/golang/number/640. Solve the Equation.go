package number

import (
	"strconv"
	"strings"
)

func get_equation(equation string)(int,int){
	var num int = 0
	var x_cnt int = 0
	var l int = len(equation)
	var left int = 0
	for left < l{
		var right int = left + 1
		for right < l && equation[right] != '+' && equation[right] != '-'{
			right++
		}
		n,err := strconv.Atoi(equation[left:right])
		if err != nil{// 4X,x,-x
			if (right - left) == 1{
				x_cnt++
			}else{
				if equation[left] == '-' && (right - left) == 2{
					x_cnt--
				}else{
					n2,_ := strconv.Atoi(equation[left:right - 1])
					x_cnt += n2
				}
			}
		}else{
			num += n
		}
		if right >= l{
			break
		}
		if equation[right] == '+'{
			left = right + 1
		}else{
			left = right
		}
	}
	return num,x_cnt
}

func SolveEquation(equation string) string {
	var express []string = strings.Split(equation,"=")
	left_num,left_x := get_equation(express[0])
	right_num,right_x := get_equation(express[1])
	right_num -= left_num
	left_x -= right_x
	if left_x == 0 && right_num == 0{
		return "Infinite solutions"
	}
	if left_x == 0 && right_num != 0{
		return "No solution"
	}
	//s1 := strconv.FormatFloat(float64(right_num)/float64(left_x), 'E', -1, 32)
	return "x=" + strconv.Itoa(right_num/left_x)
}