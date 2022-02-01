package number

//假设某一天，你访问i 号房间。
//如果算上本次访问，访问i 号房间的次数为 奇数 ，那么 第二天 需要访问nextVisit[i] 所指定的房间，其中 0 <= nextVisit[i] <= i 。
//如果算上本次访问，访问i 号房间的次数为 偶数 ，那么 第二天 需要访问(i + 1) mod n 号房间。

//输入：nextVisit = [0,1,2,0]
//输出：6
//解释：
//你每天访问房间的次序是 [0,0,1,1,2,2,3,...] 。
//第 6 天是你访问完所有房间的第一天。

func FirstDayBeenInAllRooms(nextVisit []int) int {
	var l int = len(nextVisit)
	var first []int64 = make([]int64,l)
	var second []int64 = make([]int64,l)
	var MOD int64 = 1e9 + 7
	first[0] = 1
	second[0] = 2
	for i := 1;i < l;i++{
		first[i] = (second[i - 1] + 1) % MOD
		second[i] = first[i] + first[i] - first[nextVisit[i]] + 1
	}
	return int((first[l - 1] + MOD - 1) % MOD)
}