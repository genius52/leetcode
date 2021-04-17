package number

import "strconv"

//Input: "1+-1i", "1+-1i"
//Output: "0+-2i"
//Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
func getab(s string)(a int,b int){
	var l int = len(s)
	var index int = 0
	for s[index] != '+'{
		index++
	}
	a,err1 := strconv.Atoi(s[:index])
	if err1 != nil{
		return 0,0
	}
	b,err2 := strconv.Atoi(s[index + 1:l - 1])
	if err2 != nil{
		return 0,0
	}
	return a,b
}

func complexNumberMultiply(a string, b string) string {
	a1,b1 := getab(a)
	a2,b2 := getab(b)
	c := a1 * a2 - b1 * b2
	d := a1 * b2 + a2 * b1
	res := strconv.Itoa(c) + "+" + strconv.Itoa(d) + "i"
	return res
}