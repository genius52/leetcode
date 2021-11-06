package number

import "strconv"

//Input: n = 123456789
//Output: "123.456.789"
func thousandSeparator(n int) string{
	var res string
	s := strconv.Itoa(n)
	var l int = len(s)
	cnt := 0
	for i := l - 1;i >= 0;i--{
		res = string(s[i]) + res
		cnt++
		if i != 0 && cnt == 3{
			cnt = 0
			res = "." + res
		}
	}
	return res
}

func ThousandSeparator(n int) string {
	var res string
	s := strconv.Itoa(n)
	var l int = len(s)
	var bits int = 1
	for i := l - 1;i >= 0;i--{
		res = string(s[i]) + res
		if i != 0 && bits % 3 == 0{
			res = "." + res
		}
		bits++
	}
	return res
}