package array

func specialArray(nums []int) int {
	var l int = len(nums)
	for n := l;n >= 0;n--{
		var cnt int = 0
		for i := 0;i < l;i++{
			if nums[i] >= n{
				cnt++
				if cnt > n{
					break
				}
			}
		}
		if cnt == n{
			return n
		}
	}
	return -1
}