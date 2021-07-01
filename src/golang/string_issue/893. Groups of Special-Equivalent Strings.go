package string_issue

import "strconv"

func NumSpecialEquivGroups(A []string) int {
	var l int = len(A)
	var s_len int = len(A[0])
	var record map[string]bool = make(map[string]bool)
	for i := 0;i < l;i++{
		var even [26]int
		var odd [26]int
		for j := 0;j < s_len;j++{
			if (j | 1) == j{
				odd[A[i][j] - 'a']++
			}else{
				even[A[i][j] - 'a']++
			}
		}
		var even_str string
		var odd_str string
		for k := 0;k < 26;k++{
			if even[k] != 0{
				even_str += string('a' + k) + strconv.Itoa(even[k])
			}
			if odd[k] != 0{
				odd_str += string('a' + k) + strconv.Itoa(odd[k])
			}
		}
		record[even_str + "," + odd_str] = true
	}
	return len(record)
}