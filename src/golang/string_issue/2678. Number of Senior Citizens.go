package string_issue

import "strconv"

func countSeniors(details []string) int {
	var cnt int = 0
	for _, s := range details {
		age, err := strconv.Atoi(s[11:13])
		if err == nil {
			if age > 60 {
				cnt++
			}
		}
	}
	return cnt
}
