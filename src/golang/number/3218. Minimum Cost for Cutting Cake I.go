package number

func dp_minimumCost(m int, n int, horizontalCut []int, verticalCut []int, left int, right int, top int, bottom int, memo [][][][]int) int {
	if left > right {
		sum := 0
		for i := top; i <= bottom; i++ {
			sum += horizontalCut[i]
		}
		return sum
	}
	if top > bottom {
		sum := 0
		for i := left; i <= right; i++ {
			sum += verticalCut[i]
		}
		return sum
	}
	if memo[left][right][top][bottom] != 0 {
		return memo[left][right][top][bottom]
	}
	var res int = 1<<31 - 1
	for i := left; i <= right; i++ {
		cur := dp_minimumCost(m, n, horizontalCut, verticalCut, left, i-1, top, bottom, memo) + verticalCut[i] +
			dp_minimumCost(m, n, horizontalCut, verticalCut, i+1, right, top, bottom, memo)
		res = min(res, cur)
	}
	for j := top; j <= bottom; j++ {
		cur := dp_minimumCost(m, n, horizontalCut, verticalCut, left, right, top, j-1, memo) + horizontalCut[j] +
			dp_minimumCost(m, n, horizontalCut, verticalCut, left, right, j+1, bottom, memo)
		res = min(res, cur)
	}
	memo[left][right][top][bottom] = res
	return res
}

func MinimumCost3218(m int, n int, horizontalCut []int, verticalCut []int) int {
	var memo [][][][]int = make([][][][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([][][]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = make([][]int, m+1)
			for k := 0; k <= m; k++ {
				memo[i][j][k] = make([]int, m+1)
			}
		}
	}
	return dp_minimumCost(m, n, horizontalCut, verticalCut, 0, n-2, 0, m-2, memo)
}
