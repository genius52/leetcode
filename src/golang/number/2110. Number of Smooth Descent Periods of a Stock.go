package number

//输入：prices = [3,2,1,4]
//输出：7
//解释：总共有 7 个平滑下降阶段：
//[3], [2], [1], [4], [3,2], [2,1] 和 [3,2,1]
//注意，仅一天按照定义也是平滑下降阶段。
func getDescentPeriods(prices []int) int64 {
	var l int = len(prices)
	var res int64 = 0
	var cnt int64 = 0
	for i := 0; i < l; i++ {
		if i != 0 {
			if prices[i] == prices[i-1]-1 {
				cnt++
			} else {
				cnt = 0
			}
		}
		res += (cnt + 1)
	}
	return res
}
