package number

import "strconv"

func LargestOddNumber(num string) string {
	var l int = len(num)
	i := l - 1
	for ;i >= 0;i--{
		val,_ := strconv.Atoi(string(num[i]))
		if (val | 1) == val{
			break
		}
	}
	if i == -1{
		return ""
	}
	return num[:i + 1]
}