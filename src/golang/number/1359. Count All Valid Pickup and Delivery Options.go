package number

//func jieCheng(n int) int {
//	result := 1
//	for i := 2; i <= n; i++ {
//		result *= i
//		result %= 1000000007
//	}
//	return result
//}

//输入：n = 2
//输出：6
//解释：所有可能的序列包括：
//(P1,P2,D1,D2)，(P1,P2,D2,D1)，(P1,D1,P2,D2)，(P2,P1,D1,D2)，(P2,P1,D2,D1) 和 (P2,D2,P1,D1)。
//(P1,D2,P2,D1) 是一个无效的序列，因为物品 2 的收件服务（P2）不应在物品 2 的配送服务（D2）之后。
func CountOrders(n int) int {
	var res int = 1
	var MOD int =  1e9 + 7
	for i := n;i >= 1;i--{
		res *= i
		res %= MOD
	}
	var space int = 1
	for i := 1;i <= n;i++{
		res *= space
		res %= MOD
		space += 2
	}
	return res
}