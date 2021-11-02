package array

//输入：target = [1,2,3,2,1]
//输出：3
//解释：我们需要至少 3 次操作从 intial 数组得到 target 数组。
//[0,0,0,0,0] 将下标为 0 到 4 的元素（包含二者）加 1 。
//[1,1,1,1,1] 将下标为 1 到 3 的元素（包含二者）加 1 。
//[1,2,2,2,1] 将下表为 2 的元素增加 1 。
//[1,2,3,2,1] 得到了目标数组。
func minNumberOperations(target []int) int {
	var l int = len(target)
	var dp []int = make([]int,l)
	dp[0] = target[0]
	for i := 1;i < l;i++{
		if target[i] <= target[i - 1]{
			dp[i] = dp[i - 1]
		}else{
			dp[i] = dp[i - 1] + target[i] - target[i - 1]
		}
	}
	return dp[l - 1]
}