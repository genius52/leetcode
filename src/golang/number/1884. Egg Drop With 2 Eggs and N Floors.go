package number

//输入：n = 100
//输出：14
//解释：
//一种最优的策略是：
//- 将第一枚鸡蛋从 9 楼扔下。如果碎了，那么 f 在 0 和 8 之间。将第二枚从 1 楼扔下，然后每扔一次上一层楼，在 8 次内找到 f 。总操作次数 = 1 + 8 = 9 。
//- 如果第一枚鸡蛋没有碎，那么再把第一枚鸡蛋从 22 层扔下。如果碎了，那么 f 在 9 和 21 之间。将第二枚鸡蛋从 10 楼扔下，然后每扔一次上一层楼，在 12 次内找到 f 。总操作次数 = 2 + 12 = 14 。
//- 如果第一枚鸡蛋没有再次碎掉，则按照类似的方法从 34, 45, 55, 64, 72, 79, 85, 90, 94, 97, 99 和 100 楼分别扔下第一枚鸡蛋。
//不管结果如何，最多需要扔 14 次来确定 f 。
func dp_twoEggDrop(n int, memo []int) int {
	if n == 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	var res int = n
	for i := 2; i < n; i++ {
		cur := max_int(i, 1+dp_twoEggDrop(n-i, memo))
		if cur < res {
			res = cur
		}
	}
	memo[n] = res
	return res
}

func twoEggDrop(n int) int {
	//times := int(math.Sqrt(float64(n)))
	var memo []int = make([]int, n+1)
	return dp_twoEggDrop(n, memo)
}
