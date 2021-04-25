package number

import "strconv"

//Input: n = 5
//Output: 5
//Explanation:
//Here are the non-negative integers <= 5 with their corresponding binary representations:
//0 : 0
//1 : 1
//2 : 10
//3 : 11
//4 : 100
//5 : 101
//Among them, only integer 3 disobeys the rule (two consecutive ones) and the other 5 satisfy the rule.
func FindIntegers(num int) int {
	var dp [32][2]int //dp[i][j]:
	dp[0][0] = 1
	dp[0][1] = 1
	dp[1][0] = 1
	dp[1][1] = 1
	for i := 2;i < 32;i++{
		dp[i][0] = dp[i - 1][1] + dp[i - 1][0]
		dp[i][1] = dp[i - 1][0]
	}
	var s string = strconv.Itoa(num)
	var s_len int = len(s)
	var res int = 0//dp[len(s) - 1][0] +  dp[len(s) - 1][1]
	for i := 0;i < s_len;i++{

	}
	return res
}