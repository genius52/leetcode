package number

//Input: jobDifficulty = [6,5,4,3,2,1], d = 2
//Output: 7
//Explanation: First day you can finish the first 5 jobs, total difficulty = 6.
//Second day you can finish the last job, total difficulty = 1.
//The difficulty of the schedule = 6 + 1 = 7
func dp_minDifficulty(jobDifficulty []int,l int,cur_pos int,d int,memo *[][]int)int{
	if d == 1{
		return max_int_number(jobDifficulty[cur_pos:]...)
	}
	if (*memo)[cur_pos][d] != 0{
		return (*memo)[cur_pos][d]
	}
	var total int = 2147483647
	var min_cost int = 0
	for i := cur_pos;i < l && (l - i) >= d;i++{
		min_cost = max_int(min_cost,jobDifficulty[i])
		cur := dp_minDifficulty(jobDifficulty,l,i + 1,d - 1,memo)
		total = min_int(total,min_cost + cur)
	}
	(*memo)[cur_pos][d] = total
	return total
}

func MinDifficulty(jobDifficulty []int, d int) int {
	var l int = len(jobDifficulty)
	if l < d{
		return -1
	}
	var memo [][]int = make([][]int,l + 1)
	for i := 0;i < l;i++{
		memo[i] = make([]int,d + 1)
	}
	return dp_minDifficulty(jobDifficulty,l,0,d,&memo)
}