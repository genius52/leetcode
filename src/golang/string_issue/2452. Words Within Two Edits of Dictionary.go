package string_issue

func twoEditWords(queries []string, dictionary []string) []string {
	var l1 int = len(queries)
	var l2 int = len(dictionary)
	var l3 int = len(queries[0])
	var res []string
	for i := 0;i < l1;i++{
		for j := 0;j < l2;j++{
			var diff_cnt int = 0
			for k := 0;k < l3;k++{
				if queries[i][k] != dictionary[j][k]{
					diff_cnt++
					if diff_cnt > 2{
						break
					}
				}
			}
			if diff_cnt <= 2{
				res = append(res,queries[i])
				break
			}
		}
	}
	return res
}