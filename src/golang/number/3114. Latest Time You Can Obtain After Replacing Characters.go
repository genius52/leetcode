package number

func findLatestTime(s string) string {
	var res string
	if s[0] == '?' {
		if s[1] == '0' || s[1] == '1' || s[1] == '?' {
			res += "1"
		} else {
			res += "0"
		}
	} else {
		res += string(s[0])
	}

	if s[1] == '?' {
		if res[0] == '0' {
			res += "9"
		} else {
			res += "1"
		}
	} else {
		res += string(s[1])
	}
	res += ":"
	if s[3] == '?' {
		res += "5"
	} else {
		res += string(s[3])
	}
	if s[4] == '?' {
		res += "9"
	} else {
		res += string(s[4])
	}
	return res
}
