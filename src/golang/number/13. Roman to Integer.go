package number

import "math"

//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
func romanToInt(s string) int {
	var table map[string]int = make(map[string]int)
	table["I"] = 1
	table["V"] = 5
	table["X"] = 10
	table["L"] = 50
	table["C"] = 100
	table["D"] = 500
	table["M"] = 1000
	var next map[string]string = make(map[string]string)
	next["I"] = "V"
	next["V"] = "X"
	next["X"] = "L"
	next["L"] = "C"
	next["C"] = "D"
	next["D"] = "M"
	next["M"] = ""

	var total int = 0
	for i := len(s) - 1;i >= 0;i--{
		v := table[string(s[i])]
		upper := math.MaxInt32
		if string(s[i]) != "M"{
			upper = table[next[string(s[i])]]
		}
		if total >= upper{
			total -= v
		}else{
			total += v
		}
	}
	return total
}
