package array

//Input: ranges = [[1,2],[3,4],[5,6]], left = 2, right = 5
//Output: true

func IsCovered(ranges [][]int, left int, right int) bool {
	var l int = len(ranges)
	var record [52]int
	for i := 0;i < l;i++{
		record[ranges[i][0]]++
		record[ranges[i][1] + 1]--
	}
	var cur_total [51]int
	for i := 1;i <= 50;i++{
		cur_total[i] = cur_total[i - 1] + record[i]
	}
	for i := left;i <= right;i++{
		if cur_total[i] <= 0{
			return false
		}
	}
	return true
}