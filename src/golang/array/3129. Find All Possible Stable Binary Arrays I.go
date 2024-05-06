package array

//binary array
// arr is called stable if:
//
//The number of occurrences of 0 in arr is exactly zero.
//The number of occurrences of 1 in arr is exactly one.
//Each
//subarray
// of arr with a size greater than limit must contain both 0 and 1.

func dp_numberOfStableArrays(zero_cnt int, one_cnt int, prev_num int, limit int, memo *[2][201][201]int) int {
	//if zero_cnt == 0 && one_cnt == 0 {
	//	return 1
	//}
	if prev_num == 0 {
		if zero_cnt == 0 {
			if one_cnt <= limit {
				return 1
			} else {
				return 0
			}
		}
	}
	if prev_num == 1 {
		if one_cnt == 0 {
			if zero_cnt <= limit {
				return 1
			} else {
				return 0
			}
		}
	}
	if zero_cnt < 0 || one_cnt < 0 {
		return 0
	}
	if memo[prev_num][zero_cnt][one_cnt] != -1 {
		return memo[prev_num][zero_cnt][one_cnt]
	}
	var res int = 0
	var MOD int = 1e9 + 7
	if prev_num == 1 {
		for i := 1; i <= limit && i <= zero_cnt; i++ {
			res += dp_numberOfStableArrays(zero_cnt-i, one_cnt, 0, limit, memo)
			res %= MOD
		}
	} else if prev_num == 0 {
		for i := 1; i <= limit && i <= one_cnt; i++ {
			res += dp_numberOfStableArrays(zero_cnt, one_cnt-i, 1, limit, memo)
			res %= MOD
		}
	}
	memo[prev_num][zero_cnt][one_cnt] = res
	return res
}

func NumberOfStableArrays(zero int, one int, limit int) int {
	var memo [2][201][201]int
	for i := 0; i <= 200; i++ {
		for j := 0; j <= 200; j++ {
			memo[0][i][j] = -1
			memo[1][i][j] = -1
		}
	}
	//var res int = 0
	var MOD int = 1e9 + 7
	//for i := 1; i <= limit && i <= zero; i++ {
	//	res += dp_numberOfStableArrays(zero-i, one, 0, limit, memo)
	//	res %= MOD
	//}
	//for i := 1; i <= limit && i <= one; i++ {
	//	res += dp_numberOfStableArrays(zero, one-i, 1, limit, memo)
	//	res %= MOD
	//}
	return (dp_numberOfStableArrays(zero, one, 0, limit, &memo) + dp_numberOfStableArrays(zero, one, 1, limit, &memo)) % MOD
}
