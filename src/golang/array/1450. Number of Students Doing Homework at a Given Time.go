package array

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var l int = len(startTime)
	var res int = 0
	for i := 0;i < l;i++{
		if startTime[i] <= queryTime && queryTime <= endTime[i]{
			res++
		}
	}
	return res
}