package array

//输入：startPos = [1, 0], homePos = [2, 3], rowCosts = [5, 4, 3], colCosts = [8, 2, 6, 7]
//输出：18
//解释：一个最优路径为：
//从 (1, 0) 开始
//-> 往下走到 (2, 0) 。代价为 rowCosts[2] = 3 。
//-> 往右走到 (2, 1) 。代价为 colCosts[1] = 2 。
//-> 往右走到 (2, 2) 。代价为 colCosts[2] = 6 。
//-> 往右走到 (2, 3) 。代价为 colCosts[3] = 7 。
//总代价为 3 + 2 + 6 + 7 = 18
func MinCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	var res int = 0
	var step int = 1
	if startPos[0] > homePos[0]{
		step = -1
	}
	r := startPos[0]
	for r != homePos[0]{
		r += step
		res += rowCosts[r]
	}
	if startPos[1] > homePos[1]{
		step = -1
	}else{
		step = 1
	}
	c := startPos[1]
	for c != homePos[1]{
		c += step
		res += colCosts[c]
	}
	return res
}