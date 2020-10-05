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