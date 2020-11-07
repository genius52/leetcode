package number

import (
	"sort"
	"strconv"
)

func ReorderedPowerOf2(N int) bool {
	var n int = 1
	var record map[string]bool = make(map[string]bool)
	for n <= 2147483647{
		var arr []int
		var cur_num int = n
		for cur_num > 0{
			arr = append(arr,cur_num % 10)
			cur_num = cur_num / 10
		}
		sort.Ints(arr)
		var s string
		for _,b:= range arr{
			s += strconv.Itoa(b)
		}
		record[s] = true
		n = n << 1
	}
	var arr2 []int
	for N > 0{
		arr2 = append(arr2,N % 10)
		N = N / 10
	}
	sort.Ints(arr2)
	var s2 string
	for _,b:= range arr2{
		s2 += strconv.Itoa(b)
	}
	if _,ok := record[s2];ok{
		return true
	}
	return false
}