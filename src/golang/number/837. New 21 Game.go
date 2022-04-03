package number

//爱丽丝参与一个大致基于纸牌游戏 “21点” 规则的游戏，描述如下：
//爱丽丝以 0 分开始，并在她的得分少于 K 分时抽取数字。 抽取时，她从 [1, W] 的范围中随机获得一个整数作为分数进行累计，
//其中 W 是整数。 每次抽取都是独立的，其结果具有相同的概率。
//当爱丽丝获得不少于 K 分时，她就停止抽取数字。 爱丽丝的分数不超过 N 的概率是多少？

//Input: n = 6, k = 1, maxPts = 10
//Output: 0.60000
//Input: n = 21, k = 17, maxPts = 10
//Output: 0.73278

func New21Game(N int, K int, maxPts int) float64{
	if N < K || K == 0 || K + maxPts <= N{
		return 1
	}
	var dp []float64 = make([]float64,K + maxPts)
	var sum float64 = 0
	for i := N;i >= K;i--{
		dp[i] = 1
		sum += 1
	}
	for i := K - 1;i >= 0;i--{
		dp[i] = sum / float64(maxPts)
		sum += dp[i] - dp[i + maxPts]
	}
	return dp[0]
}

//func New21Game(N int, K int, maxPts int) float64 {
//	var dp []float64 = make([]float64,30000)
//	dp[0] = 1.0
//	var rate float64 = 1/float64(maxPts)
//	var cur int = 1
//	for cur < K{
//		for i := 1;i <= maxPts;i++{
//			if (cur - i) < 0{
//				break
//			}
//			dp[cur] += (dp[cur - i] * rate)
//		}
//		cur++
//	}
//	for cur < (K + maxPts){
//		var i int = maxPts
//		if (cur - i) < 0{
//			i = cur
//		}
//		for ;i >= 1;i--{
//			if (cur - i) < 0{
//				continue
//			}
//			if (cur - i) >= K{
//				break
//			}
//			dp[cur] += (dp[cur - i] * rate)
//		}
//		cur++
//	}
//	var res float64
//	for i := N;i >= K;i--{
//		res += dp[i]
//	}
//	return res
//}