package number

import "strconv"

//9923
//Input: 2736
//Output: 7236
func MaximumSwap2(num int) int{
	var s string = strconv.Itoa(num)
	var l int = len(s)
	var max_digit int = -1
	var max_index int = -1
	var left_index int = -1
	var right_index int = -1
	for i := l - 1;i >= 0;i--{
		cur_digit := int(s[i] - '0')
		if cur_digit > max_digit{
			max_digit = cur_digit
			max_index = i
		}
		if cur_digit < max_digit{
			left_index = i
			right_index = max_index
		}
	}
	if left_index == -1{
		return num
	}
	c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	c[left_index],c[right_index] = c[right_index],c[left_index]
	s = string(c)  // 再转换回 string 类型
	res,_ := strconv.Atoi(s)
	return res
}


func MaximumSwap(num int) int {
	if num <= 9{
		return num
	}
	var data []int
	var back int = num
	for num > 0{
		rest := num % 10
		data = append(data,rest)
		num = num / 10
	}
	var l int = len(data)
	for i := 0;i < l/2;i++{
		data[i],data[l - 1 - i] = data[l - 1 - i],data[i]
	}
	var dp []int = make([]int,l)
	dp[l - 1] = data[l - 1]
	var max int = data[l - 1]
	for i := l - 2;i >= 0;i--{
		if data[i] > max{
			dp[i] = data[i]
			max = data[i]
		}else{
			dp[i] = max
		}
	}
	for i := 0;i < l - 1;i++{
		if data[i] >= dp[i]{
			continue
		}
		target := dp[i + 1]
		for j := l - 1; j > i;j--{
			if data[j] == target{
				data[i],data[j] = data[j],data[i]
				var swap_num int = 0
				for k := 0;k < l;k++{
					swap_num *= 10
					swap_num += data[k]
				}
				return swap_num
			}
		}
	}
	return back
}