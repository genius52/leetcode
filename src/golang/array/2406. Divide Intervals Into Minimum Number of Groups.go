package array

func minGroups(intervals [][]int) int {
	var record [1000001]int
	var max_end int = 0
	for _,interval := range intervals{
		record[interval[0]]++
		record[interval[1] + 1]--
		if interval[1] > max_end{
			max_end = interval[1]
		}
	}
	//var prefix_sum []int = make([]int,max_end + 1)
	for i := 1;i <= max_end;i++{
		record[i] += record[i - 1]
	}
	var res int = 0
	for i := 0;i <= max_end;i++{
		if record[i] > res{
			res = record[i]
		}
	}
	return res
}