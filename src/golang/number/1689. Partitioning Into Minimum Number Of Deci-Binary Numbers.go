package number

import "strconv"

//Input: n = "32"
//Output: 3
//Explanation: 10 + 11 + 11 = 32
//Example 2:
//
//Input: n = "82734"
//Output: 8
func minPartitions2(n string) int{
	var res int = 0
	for _,c := range n{
		num := c - '0'
		if int(num) > res{
			res = int(num)
			if res == 9{
				break
			}
		}
	}
	return res
}

func minPartitions(n string) int {
	var l int = len(n)
	var res int = 0
	for i := 0;i < l;i++{
		if num,err := strconv.Atoi(string(n[i]));err == nil{
			if num > res{
				res = num
			}
		}
	}
	return res
}