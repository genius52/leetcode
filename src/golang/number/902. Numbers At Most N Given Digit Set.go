package number

import (
	"sort"
	"strconv"
)

//Input: digits = ["1","3","5","7"], n = 100
//Output: 20
//Explanation:
//The 20 numbers that can be written are:
//1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
func dfs_atMostNGivenDigitSet(available_nums []int,l1 int,origin []int,pos int,l2 int,pre_equal bool)int{
	if pos >= l2{
		return 0
	}
	if pre_equal{
		//cur_num := origin[pos]
		var i int = l1 - 1
		for i >= 0{
			if available_nums[i] <= origin[pos]{
				break
			}
			i--
		}
		if i < 0{
			return 0
		}
		var res int = 0
		if available_nums[i] == origin[pos]{
			if i != 0{
				//存在比当前最高位小的数字
				if pos == (l2 - 1){
					res += i + 1
					return res
				}else{
					res += i * dfs_atMostNGivenDigitSet(available_nums,l1,origin,pos + 1,l2,false)
				}
			}
			if pos == (l2 - 1){
				res += 1
			}else{
				res += dfs_atMostNGivenDigitSet(available_nums,l1,origin,pos + 1,l2,true)
			}
			return res
		}else{
			if pos == (l2 - 1){
				return (i + 1)
			}else{
				return (i + 1) * dfs_atMostNGivenDigitSet(available_nums,l1,origin,pos + 1,l2,false)
			}
		}
	}else{
		var res int = 1
		for i := pos ;i < l2;i++{
			res *= l1
		}
		return res
	}
}

func AtMostNGivenDigitSet(digits []string, n int) int {
	var choose_len int = len(digits)
	var available_nums []int = make([]int,choose_len)
	for i := 0;i < choose_len;i++{
		n,_ := strconv.Atoi(digits[i])
		available_nums[i] = n
	}
	sort.Ints(available_nums)

	var origin []int
	for n > 0{
		origin = append([]int{n % 10},origin...)
		n = n/10
	}
	var res int = 0
	var target_len int = len(origin)
	res += dfs_atMostNGivenDigitSet(available_nums,choose_len,origin,0,target_len,true)
	for i := 1;i < target_len;i++{
		cnt := 1
		for j := i;j < target_len;j++{
			cnt *= choose_len
		}
		res += cnt
	}
	return res
}