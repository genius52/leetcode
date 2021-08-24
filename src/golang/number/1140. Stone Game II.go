package number

import "math"

//On each player's turn, that player can take all the stones in the first X remaining piles,
//where 1 <= X <= 2M.  Then, we set M = max(M, X).
//Input: piles = [2,7,9,4,4]
//Output: 10
func dp_stoneGameII(piles []int,suffix_sum []int, l int,pos int,m int,memo [][]int )int{
	if (pos + m * 2) >= l{
		return suffix_sum[pos]
	}
	if memo[pos][m] != 0{
		return memo[pos][m]
	}
	var res int = 0
	for i := 1;i <= m * 2 && (pos + i) < l;i++{
		cur := suffix_sum[pos]
		next_m := int(math.Max(float64(i),float64(m)))
		res = max_int(res,cur - dp_stoneGameII(piles,suffix_sum,l,pos + i,next_m,memo))
	}
	memo[pos][m] = res
	return res
}

func StoneGameII(piles []int) int {
	var l int = len(piles)
	var suffix_sum []int = make([]int,l)
	suffix_sum[l - 1] = piles[l - 1]
	for i := l - 2;i >= 0;i--{
		suffix_sum[i] = piles[i] + suffix_sum[i + 1]
	}
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,l)
	}
	return dp_stoneGameII(piles,suffix_sum,l,0,1,memo)
}