package string_issue

import "strings"

func maximumTime(time string) string {
	var res []string = strings.Split(time, "")
	if res[0] == "?" {
		if res[1] >= "4" && res[1] <= "9" {
			res[0] = "1"
		} else {
			res[0] = "2"
		}
	}
	if res[1] == "?" {
		if res[0] == "2" {
			res[1] = "3"
		} else {
			res[1] = "9"
		}
	}
	if res[3] == "?" {
		res[3] = "5"
	}
	if res[4] == "?" {
		res[4] = "9"
	}
	return strings.Join(res, "")
}
