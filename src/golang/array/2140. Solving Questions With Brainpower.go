package array

func dp_mostPoints(questions [][]int,l int,idx int,memo []int64)int64{
	if idx >= l{
		return 0
	}
	if memo[idx] != -1{
		return memo[idx]
	}
	var choose int64 = int64(questions[idx][0]) + dp_mostPoints(questions,l,idx + questions[idx][1] + 1,memo)
	var skip int64 = dp_mostPoints(questions,l,idx + 1,memo)
	memo[idx] = max_int64(choose,skip)
	return memo[idx]
}

func mostPoints(questions [][]int) int64 {
	var l int = len(questions)
	var memo []int64 = make([]int64,l)
	for i := 0;i < l;i++{
		memo[i] = -1
	}
	return dp_mostPoints(questions,l,0,memo)
}
