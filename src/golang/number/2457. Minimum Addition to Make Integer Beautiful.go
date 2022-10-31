package number

import "math"

func MakeIntegerBeautiful(n int64, target int) int64 {
	var sum int64 = 0
	var data []int64
	var n2 int64 = n
	for n2 > 0{
		data = append(data,n2 % 10)
		sum += n2 % 10
		n2 /= 10
	}
	var res int64 = 0
	var offset int = 0
	for sum > int64(target){
		if data[0] == 0{
			data = data[1:]
			offset++
			continue
		}
		res += (10 - data[0]) * int64(math.Pow(10,float64(offset)))
		offset++
		var cur_sum int64 = 0
		var upgrade bool = true
		data = data[1:]
		for i := 0;i < len(data);i++{
			if upgrade{
				if data[i] == 9{
					data[i] = 0
					upgrade = true
				}else{
					data[i]++
					cur_sum += data[i]
					upgrade = false
				}
			}else{
				cur_sum += int64(data[i])
			}
		}
		if upgrade{
			cur_sum++
			data = append(data,1)
		}
		sum = cur_sum
	}
	return res
}