package number

import "sort"

func max_int64(a,b int64)int64{
	if a > b {
		return a
	}else{
		return b
	}
}

func MaximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	var l int = len(flowers)
	var data []int
	var full_cnt int = 0
	for i := 0;i < l;i++{
		if flowers[i] >= target{
			full_cnt++
		}else{
			data = append(data,flowers[i])
		}
	}
	sort.Ints(data)
	var newFlowers2 int64 = newFlowers
	var add_full_cnt int = 0
	var min_flowers int = 2147483647
	for j := 0;j < len(data);j++{
		if data[j] < min_flowers{
			min_flowers = data[j]
		}
	}
	i := len(data) - 1
	for ;i >= 0;i--{
		if int64(target - data[i]) > newFlowers2{
			break
		}
		newFlowers2 = newFlowers2 - int64(target - data[i])
		add_full_cnt++
	}
	if i == -1{
		return max_int64(int64(full) * int64(l),int64(full) * int64(full_cnt) + int64(partial * min_flowers))
	}
	var max_val int64 = int64(full) * int64(full_cnt + add_full_cnt) + int64(partial * min_flowers)
	var left int = data[0]
	var right int = data[len(data) - 1]
	var guess_min_flowers int64 = int64(data[0])
	for left <= right{
		mid := (left + right)/2
		var sum int64 = 0
		for j := 0;j < len(data);j++{
			if data[j] < mid{
				sum += int64(mid - data[j])
			}
		}
		if sum == newFlowers{
			guess_min_flowers = int64(mid)
			break
		}else if sum < newFlowers{
			left = mid + 1
			guess_min_flowers = int64(left)
		}else if sum > newFlowers2{
			right = mid - 1
		}
	}
	add_low := int64(full * full_cnt) + int64(partial) * guess_min_flowers
	res := max_int64(add_low,max_val)
	return res
}