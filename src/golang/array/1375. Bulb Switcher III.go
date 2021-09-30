package array

//Input: light = [2,1,3,5,4]
//Output: 3
func numTimesAllBlue(light []int) int{
	var sum int64 = 0
	var max_val int = 0
	var res int = 0
	var l int = len(light)
	for i := 0;i < l;i++{
		sum += int64(light[i])
		max_val = max_int(max_val,light[i])
		if sum == int64((1 + max_val) * max_val / 2){
			res++
		}
	}
	return res
}

func numTimesAllBlue2(light []int) int{
	var max_val int = 0
	var res int = 0
	var l int = len(light)
	for i := 0;i < l;i++{
		max_val = max_int(max_val,light[i])
		if max_val == (i + 1){
			res++
		}
	}
	return res
}

func NumTimesAllBlue(light []int) int {
	var l int = len(light)
	var right int = light[0]
	var res int = 0
	if light[0] == 1{
		res++
	}
	for step := 1;step < l;step++{
		if light[step] > right {
			if light[step] == step + 1{
				res++
			}
			right = light[step]
		}else{
			if step + 1 == right{
				res++
			}
		}
	}
	return res
}