package array

//Input: nums = [2, 2, 3, 3, 3, 4]
//Output: 9
func dp_deleteAndEarn(record []int,pos int,memo map[int]int)int{
	if pos >= len(record){
		return 0
	}
	if n,ok := memo[pos];ok{
		return n
	}
	select_cur := record[pos] * pos + dp_deleteAndEarn(record,pos + 2,memo)
	skip_cur := dp_deleteAndEarn(record,pos + 1,memo)
	if select_cur > skip_cur{
		memo[pos] = select_cur
	}else{
		memo[pos] = skip_cur
	}
	return memo[pos]
}

func DeleteAndEarn(nums []int) int {
	var record []int = make([]int,10001)
	for _,n := range nums{
		record[n]++
	}
	var memo map[int]int = make(map[int]int)
	return dp_deleteAndEarn(record,0,memo)
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