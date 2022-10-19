package array

func findMaxK(nums []int) int {
	var max_val int = -1
	var exist map[int]bool = make(map[int]bool)
	for _,n := range nums{
		exist[n] = true
		if n > 0{
			if n > max_val{
				if _,ok := exist[-n];ok{
					max_val = n
				}
			}
		}else if n < 0{
			if -n > max_val{
				if _,ok := exist[-n];ok{
					max_val = -n
				}
			}
		}
	}
	return max_val
}