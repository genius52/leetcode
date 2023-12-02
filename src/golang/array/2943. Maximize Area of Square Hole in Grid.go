package array

import "sort"

func MaximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	//hBars = append(hBars, 1)
	//hBars = append(hBars, n+2)
	//vBars = append(vBars, 1)
	//vBars = append(vBars, m+2)
	sort.Ints(hBars)
	sort.Ints(vBars)
	var res int = 1
	//loop:
	//var edge_len int = 1
	for left := 0; left < len(hBars); left++ {
		var right int = left + 1
		for right < len(hBars) {
			if hBars[right]-1 != hBars[right-1] {
				break
			}
			right++
		}
	}
	return res
}
