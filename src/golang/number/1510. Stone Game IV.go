package number

//给你正整数 n ，且已知两个人都采取最优策略。如果 Alice 会赢得比赛，那么返回 True ，否则返回 False 。
//输入：n = 7
//输出：false
//解释：当 Bob 采取最优策略时，Alice 无法赢得比赛。
//如果 Alice 一开始拿走 4 个石子， Bob 会拿走 1 个石子，然后 Alice 只能拿走 1 个石子，Bob 拿走最后一个石子并赢得胜利（7 -> 3 -> 2 -> 1 -> 0）。
//如果 Alice 一开始拿走 1 个石子， Bob 会拿走 4 个石子，然后 Alice 只能拿走 1 个石子，Bob 拿走最后一个石子并赢得胜利（7 -> 6 -> 2 -> 1 -> 0）。
func dp_winnerSquareGame(n int,memo map[int]bool)bool{
	if n == 0{
		return false //current player lose
	}
	if _,ok := memo[n];ok{
		return memo[n]
	}
	for i := 1;i * i <= n;i++{
		ret := dp_winnerSquareGame(n - i * i,memo)
		if !ret{//enemy lose = I win
			memo[n] = true
			return true
		}
	}
	memo[n] = false
	return false
}

func WinnerSquareGame(n int) bool {
	var memo map[int]bool = make(map[int]bool)
	return dp_winnerSquareGame(n,memo)
}