package string_issue

import "strings"

//Input: number = "--17-5 229 35-39475 "
//Output: "175-229-353-94-75"
func ReformatNumber(number string) string {
	number2 := strings.ReplaceAll(number,"-","")
	number2 = strings.ReplaceAll(number2," ","")
	var index int = 0
	var l int = len(number2)
	if l <= 3{
		return number2
	}
	var res string
	if l % 3 == 0{
		for index < l{
			if len(res) != 0 && index % 3 == 0{
				res += "-"
			}
			res += string(number2[index])
			index++
		}
	}else if l % 3 == 1{
		for index < l{
			if len(res) != 0 && index % 3 == 0{
				res += "-"
			}
			if l - index <= 4{
				break
			}
			res += string(number2[index])
			index++
		}
		res += number2[l - 4:l - 2] + "-" + number2[l - 2:l]
	}else if l % 3 == 2{
		for index < l{
			if len(res) != 0 && index % 3 == 0{
				res += "-"
			}
			if l - index <= 2{
				break
			}
			res += string(number2[index])
			index++
		}
		res += number2[l - 2:l]
	}
	return res
}