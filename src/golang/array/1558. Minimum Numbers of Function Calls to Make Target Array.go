package array

import "math"

func MinOperations2(nums []int) int {
	var record map[int]int = make(map[int]int)
	for _,n := range nums{
		record[n]++
	}
	var res int = 0
	var max_div int = 0
	for k,cnt := range record{
		var div_cnt int = 0
		for k != 0{
			if k % 2 == 1{
				res += cnt
			}
			k /= 2
			if k > 0{
				div_cnt++
			}
		}
		if max_div < div_cnt{
			max_div = div_cnt
		}
	}
	res += max_div
	return res
}

func minOperations2(nums []int) int{
	var max_num int = 0
	var res int = 0
	for _,n := range nums{
		if n > max_num{
			max_num = n
		}
		for n > 0{
			if n | 1 == n{
				res++
			}
			n /= 2
		}
	}
	if max_num > 0{
		res += int(math.Log2(float64(max_num)))
	}
	return res
}