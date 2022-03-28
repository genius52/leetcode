package number

func maxValueOfCoins(piles [][]int, k int) int {
	var l int = len(piles)
	var prefix_sum [][]int = make([][]int,l)
	var res int = 0
	for i := 0;i < l;i++{
		for j := 0;j < len(piles[i]);j++{
			if j == 0{
				prefix_sum[i][j] = piles[i][j]
			}else{
				prefix_sum[i][j] = piles[i][j] + prefix_sum[i][j - 1]
			}
		}
	}
	var dp []int = make([]int,k)

	return res
}