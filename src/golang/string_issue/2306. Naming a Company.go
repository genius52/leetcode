package string_issue

func DistinctNames(ideas []string) int64 {
	var record [26]map[string]bool
	for i := 0;i < 26;i++{
		record[i] = make(map[string]bool)
	}
	var l int = len(ideas)
	for i := 0;i < l;i++{
		record[ideas[i][0] - 'a'][ideas[i][1:]] = true
	}
	var res int64 = 0
	for i := 0;i < 26;i++{
		for j := i + 1;j < 26;j++{
			var dup_cnt int64 = 0
			for k1,_ := range record[i]{
				if _,ok := record[j][k1];ok{
					dup_cnt++
				}
			}
			res += 2 * (int64(len(record[i])) - dup_cnt) * (int64(len(record[j])) - dup_cnt)
		}
	}
	return res
}