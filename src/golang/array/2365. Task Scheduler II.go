package array

func taskSchedulerII(tasks []int, space int) int64 {
	var l int = len(tasks)
	var record map[int]int64 = make(map[int]int64)
	var days int64 = 0
	for i := 0;i < l;i++{
		if preday,ok := record[tasks[i]];ok{
			cur_space := days - preday - 1
			if cur_space < int64(space){
				days += int64(space) - cur_space
			}
			record[tasks[i]] = days
			if i == l - 1{
				break
			}
			days++
		} else{
			record[tasks[i]] = days
			if i == l - 1{
				break
			}
			days++
		}
	}
	return days
}