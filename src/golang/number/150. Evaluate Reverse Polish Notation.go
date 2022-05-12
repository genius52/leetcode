package number

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	var data []int
	var res int = 0
	for i := 0;i < len(tokens);i++{
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/"{
			n,err := strconv.Atoi(tokens[i])
			if err == nil{
				data = append(data,n)
			}
		}else{
			first := data[len(data) - 2]
			second := data[len(data) - 1]
			data = data[:len(data) - 2]
			var result int = 0
			switch(tokens[i]){
			case "+":
				result = first + second
			case "-":
				result = first - second
			case "*":
				result = first * second
			case "/":
				result = first / second
			}
			res += result
			data = append(data, res)
			res = 0
		}
	}
	return data[0]
}

func EvalRPN2(tokens []string) int{
	var l int = len(tokens)
	var q list.List
	for i := 0;i < l;i++{
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/"{
			n,err := strconv.Atoi(tokens[i])
			if err == nil{
				q.PushBack(n)
			}
		}else{
			var result int = 0
			num2 := q.Back().Value.(int)
			q.Remove(q.Back())
			num1 := q.Back().Value.(int)
			q.Remove(q.Back())
			switch(tokens[i]){
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2
			}
			q.PushBack(result)
		}
	}
	return q.Back().Value.(int)
}

func EvalRPN3(tokens []string) int{
	var l int = len(tokens)
	var char_opt map[string]func(i int,j int)int = make(map[string]func(i int,j int)int)
	char_opt["+"] = func(i int, j int) int {
		return i + j
	}
	char_opt["-"] = func(i int, j int) int {
		return i - j
	}
	char_opt["*"] = func(i int, j int) int {
		return i * j
	}
	char_opt["/"] = func(i int, j int) int {
		return i / j
	}
	var q list.List
	for i := 0;i < l;i++{
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/"{
			n,err := strconv.Atoi(tokens[i])
			if err == nil{
				q.PushBack(n)
			}
		}else{
			num2 := q.Back().Value.(int)
			q.Remove(q.Back())
			num1 := q.Back().Value.(int)
			q.Remove(q.Back())
			q.PushBack(char_opt[tokens[i]](num1,num2))
		}
	}
	return q.Back().Value.(int)
}