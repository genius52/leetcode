package list_queue

func FindTheWinner(n int, k int) int {
	var record []int = make([]int,n)
	for i := 0;i < n;i++{
		record[i] = i + 1
	}
	cur_pos := 0
	for n > 1{
		next := (cur_pos + k - 1) % n
		cur_pos = next % (n - 1)
		record = append(record[:next],record[next + 1:]...)
		n--
	}
	return record[0]
}