package number

//Input: [0,1,0,1,0,1,99]
//Output: 99
func SingleNumber(nums []int) int {
	var record []int = make([]int,32)
	for _,n := range nums{
		var bit int = 1
		for i := 0;i < 32;i++{
			res := n | bit
			if n == res{
				record[i]++
				record[i] = record[i] % 3
			}
			bit = bit << 1
		}
	}
	var res int = 0;
	var b int = 1 << 31
	var is_negtive bool = false
	for i := 31;i >= 0;i--{
		if i == 31{
			if record[i] == 1{
				is_negtive = true
			}
		}
		if is_negtive{
			if record[i] == 0{
				res = res | b
			}
		}else{
			if record[i] == 1{
				res = res | b
			}
		}
		b = b >> 1
	}
	if is_negtive{
		return -(res + 1)
	}
	return res
}