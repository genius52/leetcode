package array

func HardestWorker(n int, logs [][]int) int {
	var record []int = make([]int,n)
	var pre int = 0
	var max_val int = 0
	var id int = n
	for _,log := range logs{
		duration := log[1] - pre
		if duration > max_val{
			record[log[0]] = duration
			max_val = duration
			id = log[0]
		}else if duration == max_val{
			record[log[0]] = duration
			if log[0] < id{
				id = log[0]
			}
		}
		pre = log[1]
	}
	return id
}