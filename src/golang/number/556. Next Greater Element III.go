package number

import "math"

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