package string_issue

import "math"

//Input: s = "00110", k = 2
//Output: true
//Input: s = "00110110", k = 2
//Output: true
//Explanation: The binary codes of length 2 are "00", "01", "10" and "11".
//They can be all found as substrings at indicies 0, 1, 3 and 2 respectively.
func hasAllCodes(s string, k int) bool {
	var record map[string]bool = make(map[string]bool)
	var l int = len(s)
	for i := 0;i + k <= l;i++{
		record[s[i:i + k]] = true
	}
	return len(record) == int(math.Pow(2,float64(k)))
}