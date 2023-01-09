package array

func CaptureForts(forts []int) int {
	var l int = len(forts)
	var res int = 0
	var left int = 0
	for left < l{
		for left < l && forts[left] == 0{
			left++
		}
		var right int = left + 1
		for right < l && forts[right] == 0{
			right++
		}
		if right < l && forts[right] == -forts[left]{
			cur := right - left - 1
			if cur > res{
				res = cur
			}
		}
		left = right
	}
	return res
}