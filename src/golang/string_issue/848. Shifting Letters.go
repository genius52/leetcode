package string_issue

//Input: S = "abc", shifts = [3,5,9]
//Output: "rpl"
//Explanation:
//We start with "abc".
//After shifting the first 1 letters of S by 3, we have "dbc".
//After shifting the first 2 letters of S by 5, we have "igc".
//After shifting the first 3 letters of S by 9, we have "rpl", the answer.
func ShiftingLetters(S string, shifts []int) string {
	var l int = len(shifts)
	for i := 1;i < l;i++{
		shifts[l - 1 - i] += shifts[l - i]
	}
	var res string
	for i := 0;i < l;i++{
		offset := (int(S[i]) - int('a') + shifts[i])% 26
		new_c := offset + 'a'
		c := new_c % 97 + 97
		res += string(c)
	}
	return res
}