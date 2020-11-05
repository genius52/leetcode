package number

import "strconv"

func dp_splitIntoFibonacci(S string,pos int,num1 int,num2 int,nums *[]int)bool{
	var l int = len(S)
	if pos >= l{
		return true
	}
	var target int = num1 + num2
	var target_len int = 0
	for target > 0{
		target_len++
		target = target / 10
	}
	if target_len == 0{
		target_len++
	}
	for i := pos + target_len;i <= l;i++{
		if S[pos] == '0' && i > pos + target_len{
			break
		}
		cur_num,_ := strconv.Atoi(S[pos:i])
		if cur_num > 2147483647{
			break
		}
		if cur_num > num1 + num2{
			break
		}
		if cur_num == num1 + num2{
			*nums = append(*nums,cur_num)
			ret := dp_splitIntoFibonacci(S,i,num2,cur_num,nums)
			if ret{
				return true
			}
			*nums = (*nums)[:len(*nums) - 1]
		}
	}
	return false
}

func SplitIntoFibonacci(S string) []int {
	var l int = len(S)
	var nums []int
	for i := 1;i <= l/2;i++{
		first_num,_ := strconv.Atoi(S[:i])
		if S[0] == '0' && i > 1{
			break
		}
		if first_num > 2147483647{
			break
		}
		nums = append(nums,first_num)
		for j := i + 1;j <= i + l/3 && j < l;j++{
			if S[i] == '0' && j > i + 1{
				break
			}
			second_num,_ := strconv.Atoi(S[i:j])
			if second_num > 2147483647{
				break
			}
			nums = append(nums,second_num)
			//var nums []int = []int{first_num,second_num}
			ret := dp_splitIntoFibonacci(S,j,first_num,second_num,&nums)
			if ret{
				return nums
			}
			nums = nums[:len(nums) - 1]
		}
		nums = nums[:len(nums) - 1]
	}
	return []int{}
}