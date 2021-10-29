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

func dfs_winnerSquareGame(n int,is_alice bool,memo [][2]int)bool{
	if n == 0{
		return !is_alice //current player lose
	}
	if is_alice{
		if memo[n][1] != -1{
			return memo[n][1] == 1
		}
	}else{
		if memo[n][0] != -1{
			return memo[n][0] == 1
		}
	}
	var res bool = false
	for i := 1;i * i <= n;i++{
		res = dfs_winnerSquareGame(n - i * i,!is_alice,memo)
		if is_alice{
			if res{ //Alice can win
				break
			}
		}else{
			if !res{ //Bob will win = Alice will lose
				break
			}
		}
	}
	var val int = 0
	if res{
		val = 1
	}
	if is_alice{
		memo[n][1] = val
	}else{
		memo[n][0] = val
	}
	return res
}

func winnerSquareGame(n int) bool {
	var memo [][2]int = make([][2]int,n + 1)
	for i := 0;i <= n;i++{
		memo[i][0] = -1
		memo[i][1] = -1
	}
	return dfs_winnerSquareGame(n,true,memo)
}