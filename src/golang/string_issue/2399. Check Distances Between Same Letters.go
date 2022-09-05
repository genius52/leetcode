package string_issue

func checkDistances(s string, distance []int) bool {
	var record [26][]int
	for i,c := range s{
		idx := int(c - 'a')
		record[idx] = append(record[idx],i)
	}
	for i := 0;i < 26;i++{
		if len(record[i]) == 0{
			continue
		}
		if record[i][1] - record[i][0] - 1 != distance[i]{
			return false
		}
	}
	return true
}