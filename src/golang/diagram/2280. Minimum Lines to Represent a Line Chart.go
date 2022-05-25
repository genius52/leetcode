package diagram

import "sort"

func minimumLines(stockPrices [][]int) int {
	var l int = len(stockPrices)
	if l <= 1{
		return 0
	}
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})
	var cnt int = 1
	for i := 1;i < l - 1;i++{
		x1 := stockPrices[i - 1][0]
		y1 := stockPrices[i - 1][1]
		x2 := stockPrices[i][0]
		y2 := stockPrices[i][1]
		x3 := stockPrices[i + 1][0]
		y3 := stockPrices[i + 1][1]
		v1 := (y2 - y1) * (x3 - x2)
		v2 := (y3 - y2) * (x2 - x1)
		if v1 != v2{
			cnt++
		}
	}
	return cnt
}