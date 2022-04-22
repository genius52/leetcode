package array

//Input: nums = [2, 2, 3, 3, 3, 4]
//Output: 9
func dp_deleteAndEarn(record []int,val int,max_val int,memo map[int]int)int{
	if val > max_val{
		return 0
	}
	if n,ok := memo[val];ok{
		return n
	}
	select_cur := record[val] * val + dp_deleteAndEarn(record,val + 2,max_val,memo)
	skip_cur := dp_deleteAndEarn(record,val + 1,max_val,memo)
	if select_cur > skip_cur{
		memo[val] = select_cur
	}else{
		memo[val] = skip_cur
	}
	return memo[val]
}

func DeleteAndEarn(nums []int) int {
	var record []int = make([]int,10001)
	var max_val int = 0
	var min_val int = 2147483647
	for _,n := range nums{
		if n > max_val{
			max_val = n
		}
		if n < min_val{
			min_val = n
		}
		record[n]++
	}
	var memo map[int]int = make(map[int]int)
	return dp_deleteAndEarn(record,min_val,max_val,memo)
}

func DeleteAndEarn2(nums []int) int {
	var record []int = make([]int,10001)
	for _,n := range nums{
		record[n]++
	}
	var skip []int = make([]int,10001)
	var choose []int = make([]int,10001)
	var res int = 0
	for i := 1;i <= 10000;i++{
		if record[i] == 0{
			if choose[i - 1] > skip[i - 1]{
				skip[i] = choose[i - 1]
			}else{
				skip[i] = skip[i - 1]
			}
			choose[i] = choose[i - 1]
			continue
		}
		choose[i] = record[i] * i + skip[i - 1]
		if choose[i - 1] > skip[i - 1]{
			skip[i] = choose[i - 1]
		}else{
			skip[i] = skip[i - 1]
		}
		if choose[i] > skip[i]{
			res = choose[i]
		}else{
			res = skip[i]
		}
	}
	return res
}