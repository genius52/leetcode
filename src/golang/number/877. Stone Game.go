package number

func dp_stoneGame(piles []int,begin int,end int,my_round bool,memo map[int]map[int]int)int{
	if begin > end{
		return 0
	}
	if _,ok1 := memo[begin];ok1{
		if _,ok2 := memo[begin][end];ok2{
			return memo[begin][end]
		}
	}else{
		memo[begin] = make(map[int]int)
	}
	if my_round{
		choose_begin := piles[begin] + dp_stoneGame(piles,begin + 1,end,!my_round,memo)
		choose_end := piles[end] + dp_stoneGame(piles,begin,end - 1,!my_round,memo)
		memo[begin][end] = max_int(choose_begin,choose_end)
	}else{
		//Your opponent want to minimize your score
		choose_begin := - piles[begin] + dp_stoneGame(piles,begin + 1,end,!my_round,memo)
		choose_end := - piles[end] + dp_stoneGame(piles,begin,end - 1,!my_round,memo)
		memo[begin][end] = min_int(choose_begin,choose_end)
	}
	return memo[begin][end]
}

func StoneGame(piles []int) bool {
	var l int = len(piles)
	var memo map[int]map[int]int = make(map[int]map[int]int)
	val := dp_stoneGame(piles,0,l - 1,true,memo)
	return  val > 0
}