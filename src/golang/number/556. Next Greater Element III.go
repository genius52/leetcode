package number

import (
	"math"
	"sort"
)

func NextGreaterElement2(n int) int{
	var num []int
	for n > 0{
		num = append([]int{n%10},num...)
		n = n / 10
	}
	var find bool = false
	var l int = len(num)
	for i := l - 2;i >= 0;i--{
		var min_num int = math.MaxInt32
		var min_index int = 0
		for j := i + 1;j < l;j++{
			if num[i] < num[j]{
				if num[j] < min_num{
					min_num = num[j]
					min_index = j
					find = true
				}
			}
		}
		if find{
			num[i],num[min_index] = num[min_index],num[i]
			find = true
			sort.Ints(num[i+1:])
			break
		}
	}
	if !find{
		return -1
	}
	var res int64 = 0
	for _,n := range num{
		res *= 10
		res += int64(n)
	}
	if res > math.MaxInt32{
		return -1
	}
	return int(res)
}

func NextGreaterElement(n int) int {
	var num []int
	for n > 0{
		num = append(num,n % 10)
		n = n / 10
	}
	var l int = len(num)
	for i, j := 0, l - 1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
	for i := l - 2;i >= 0;i--{
		var bigger_than_current int = -1
		for j := l - 1;j > i;j--{
			if num[j] > num[i]{
				if bigger_than_current == -1{
					bigger_than_current = j
				}else{
					if num[j] < num[bigger_than_current]{
						bigger_than_current = j
					}
				}
			}
		}
		if bigger_than_current != -1{
			num[i] ,num[bigger_than_current] = num[bigger_than_current],num[i]
			for m := i + 1;m < l;m++{
				var small_index int = m
				for n := m;n < l;n++{
					if num[n] < num[small_index]{
						small_index = n
					}
				}
				num[m],num[small_index] = num[small_index],num[m]
			}
			var res int64 = 0
			for k := 0;k < l;k++{
				res *= 10
				res += int64(num[k])
			}
			if res > math.MaxInt32{
				return -1
			}
			return int(res)
		}
	}
	return -1
}