package number

//Input: steps = 3, arrLen = 2
//Output: 4
//Explanation: There are 4 differents ways to stay at index 0 after 3 steps.
//Right, Left, Stay
//Stay, Right, Left
//Right, Stay, Left
//Stay, Stay, Stay
func recursive_numWays(steps int,arrLen int,cur_pos int,memo *[][]int)int{
	if steps == 0 {
		if cur_pos == 0{
			return 1
		}else{
			return 0
		}
	}
	if cur_pos > steps{
		return 0
	}
	if (*memo)[cur_pos][steps] != 0{
		return (*memo)[cur_pos][steps]
	}
	var res int = 0
	if cur_pos > 0 {
		res += recursive_numWays(steps - 1,arrLen,cur_pos - 1,memo)
	}
	if cur_pos < (arrLen - 1) && cur_pos <= steps{
		res += recursive_numWays(steps - 1,arrLen,cur_pos + 1,memo)
	}
	res += recursive_numWays(steps - 1,arrLen,cur_pos,memo)
	res %= 1000000007
	(*memo)[cur_pos][steps] = res
	return res
}

func NumWays(steps int, arrLen int) int {
	var memo [][]int = make([][]int,steps/2 + 1)
	for i := 0;i < arrLen;i++{
		memo[i] = make([]int,steps + 1)
	}
	return recursive_numWays(steps,arrLen,0,&memo)
}