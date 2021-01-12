package array

func GetLastMoment(n int, left []int, right []int) int {
	var to_left int = 0
	var to_right int = n
	for _,val := range left{
		if val > to_left{
			to_left = val
		}
	}
	for _,val := range right{
		if val < to_right{
			to_right = val
		}
	}
	var left_len int = len(left)
	var right_len int = len(right)
	if left_len == 0{
		return n - to_right
	}
	if right_len == 0{
		return to_left
	}
	if to_left > n - to_right{
		return to_left
	}else{
		return n - to_right
	}
}