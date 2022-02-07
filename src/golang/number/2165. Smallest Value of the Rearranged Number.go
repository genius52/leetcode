package number

import "sort"

func SmallestNumber(num int64) int64 {
	var neg bool = false
	if num < 0{
		neg = true
		num = -num
	}
	var data []int64
	for num != 0{
		data = append(data,num % 10)
		num /= 10
	}
	sort.Slice(data, func(i, j int) bool {
		if neg{
			return data[i] > data[j]
		}else{
			return data[i] < data[j]
		}
	})
	var l int = len(data)
	if !neg{
		for i := 0;i < l;i++{
			if data[i] != 0{
				data[0],data[i] = data[i],data[0]
				break
			}
		}
	}
	var res int64 = 0
	for i:= 0;i < l;i++{
		res *= 10
		res += data[i]
	}
	if neg{
		res = -res
	}
	return res
}