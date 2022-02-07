package number

import "sort"

func minimumSum(num int) int {
	var data []int
	for num > 0{
	data = append(data,num%10)
	num /= 10
	}
	sort.Ints(data)
	var l int = len(data)
	var res int = 0
	for i := 0;i < l/2;i++{
	res += data[i] * 10 + data[l - 1 - i]
	}
	return res
}