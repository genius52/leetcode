package string_issue

//Input: s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
//Output: [true,false,false,true,true]
func CanMakePaliQueries(s string, queries [][]int) []bool {
	var l int = len(s)
	var record [][26]int = make([][26]int,l + 1)
	for i := 0;i < l;i++{
		for j := 0;j < 26;j++{
			record[i + 1][j] = record[i][j]
		}
		record[i + 1][s[i] - 'a'] = record[i][s[i] - 'a'] + 1
	}

	var q_len int = len(queries)
	var res []bool = make([]bool,q_len)
	for i := 0;i < q_len;i++{
		left := queries[i][0]
		right := queries[i][1]
		allow_odd_cnt := queries[i][2]
		if allow_odd_cnt >= (right - left + 1) / 2{
			res[i] = true
			continue
		}
		var diff_cnt int = 0
		for j := 0;j < 26;j++{
			diff := record[right + 1][j] - record[left][j]
			if diff | 1 == diff{//odd
				diff_cnt++
			}
		}
		if diff_cnt / 2 > allow_odd_cnt{
			res[i] = false
		}else{
			res[i] = true
		}
	}
	return res
}