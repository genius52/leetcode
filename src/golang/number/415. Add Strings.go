package number

import (
	"strconv"
)

//Input: num1 = "11", num2 = "123"
//Output: "134"
func AddStrings(num1 string, num2 string) string {
	var l1 int = len(num1)
	var l2 int = len(num2)
	var pos1 int = l1 - 1
	var pos2 int = l2 - 1
	var total string
	var upgrade bool = false
	for pos1 >= 0 || pos2 >= 0{
		var cur int = 0
		if pos1 >= 0{
			cur += int(num1[pos1] - '0')
		}
		if pos2 >= 0{
			cur += int(num2[pos2] - '0')
		}
		if upgrade{
			cur++
		}
		upgrade = cur >= 10
		total = strconv.Itoa(cur % 10) + total
		pos1--
		pos2--
	}
	if upgrade{
		total = "1" + total
	}
	return total
}