package number

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	var res []int = make([]int, k)
	var record map[int]map[int]bool = make(map[int]map[int]bool) //user id:
	var l int = len(logs)
	for i := 0; i < l; i++ {
		if record[logs[i][0]] == nil {
			record[logs[i][0]] = make(map[int]bool)
		}
		record[logs[i][0]][logs[i][1]] = true
	}
	for _, v := range record {
		res[len(v)-1]++
	}
	return res
}
