package number

//We have an array k of first n ugly number. We only know, at the beginning, the first one, which is 1. Then
//
//k[1] = min( k[0]x2, k[0]x3, k[0]x5). The answer is k[0]x2. So we move 2's pointer to 1. Then we test:
//
//k[2] = min( k[1]x2, k[0]x3, k[0]x5). And so on. Be careful about the cases such as 6, in which we need to forward both pointers of 2 and 3.
func NthUglyNumber(n int) int {
	if n <= 0{
		return 0
	}
	var dp []int = make([]int,n)
	dp[0] = 1
	var i,j,k int = 0,0,0
	var cnt int = 1
	for cnt < n{
		dp[cnt] = min_int_number(dp[i] * 2,dp[j] * 3,dp[k] * 5)
		if dp[cnt] == dp[i] * 2 {
			i++
		}
		if dp[cnt] ==  dp[j] * 3 {
			j++
		}
		if dp[cnt] ==  dp[k] * 5{
			k++
		}
		cnt++
	}
	return dp[cnt - 1]
}

func NthUglyNumber2(n int) int{
	if n <= 0{
		return 0
	}
	var dp []int = make([]int,n)
	dp[0] = 1
	var last_multiple_two int = 0
	var last_multiple_three int = 0
	var last_multiple_five int = 0
	for i := 1;i < n;i++{
		dp[i] = min_int_number(dp[last_multiple_two] * 2,dp[last_multiple_three] * 3,dp[last_multiple_five] * 5)
		if dp[i] == dp[last_multiple_two] * 2{
			last_multiple_two++
		}
		if dp[i] == dp[last_multiple_three] * 3{
			last_multiple_three++
		}
		if dp[i] == dp[last_multiple_five] * 5{
			last_multiple_five++
		}
	}
	return dp[n - 1]
}