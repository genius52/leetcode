package array

func numberOfPoints(nums [][]int) int {
	var record [101]bool
	var cnt int = 0
	for _, position := range nums {
		for visit := position[0]; visit <= position[1]; visit++ {
			if !record[visit] {
				cnt++
				record[visit] = true
			}
		}
	}
	return cnt
}
