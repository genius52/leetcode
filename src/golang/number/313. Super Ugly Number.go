package number

import (
	"math"
)

//Input: n = 12, primes = [2,7,13,19]
//Output: 32
//Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12
//             super ugly numbers given primes = [2,7,13,19] of size 4.
func NthSuperUglyNumber(n int, primes []int) int {
	l := len(primes)
	if n == 1{
		return 1
	}
	var dp []int = make([]int,1)
	dp[0] = 1
	var cnt int = 1
	for cnt < n {
		var min int = math.MaxInt32
		for i := 0;i < l;i++{
			for j := len(dp) - 1;j >= 0;j--{
				val := primes[i] * dp[j]
				if val < dp[len(dp) - 1]{
					break
				}
				if val > dp[len(dp) - 1] && val < min{
					min = val
				}
			}
		}
		dp = append(dp,min)
		cnt++
	}
	return dp[n - 1]
}

func NthSuperUglyNumber2(n int, primes []int) int {
	l := len(primes)
	if n == 1{
		return 1
	}
	var dp []int = make([]int,n)
	var index_record []int = make([]int,l)
	dp[0] = 1
	for i := 1;i < n;i++ {
		dp[i] = math.MaxInt32
		for j := 0;j < l;j++{
			val := primes[j] * dp[index_record[j]]
			if val < dp[i]{
				dp[i] = val
			}
		}
		for k := 0;k < l;k++{
			if primes[k] * dp[index_record[k]] == dp[i]{
				index_record[k]++
			}
		}
	}
	return dp[n - 1]
}