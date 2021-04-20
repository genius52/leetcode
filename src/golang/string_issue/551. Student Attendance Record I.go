package string_issue

//Input: s = "PPALLP"
//Output: true
//Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.
func checkRecord(s string) bool {
	var a_cnt int = 0
	var later_cnt int = 0
	var l int = len(s)
	for i := 0;i < l;i++{
		if s[i] == 'A'{
			a_cnt++
			if a_cnt > 1{
				return false
			}
			later_cnt = 0
		}else if s[i] == 'L'{
			later_cnt++
			if later_cnt > 2{
				return false
			}
		}else{
			later_cnt = 0
		}
	}
	return true
}