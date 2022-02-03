package string_issue

func numberOfUniqueGoodSubsequences(binary string) int {
	var l int = len(binary)
	var lead_zero int = 0 //以0开头的个数
	var lead_one int = 0 //以1开头的个数
	var exist0 int = 0
	var MOD int = 1e9 + 7
	for i := l - 1;i >= 0;i--{
		if binary[i] == '0'{
			lead_zero = lead_zero + lead_one + 1
			lead_zero %= MOD
			exist0 = 1
		}else{
			lead_one = lead_zero + lead_one + 1
			lead_one %= MOD
		}
	}
	return (lead_one + exist0) % MOD
}