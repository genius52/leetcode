package array

func divideArray(nums []int) bool {
	var cnt [501]int
	for _,n := range nums{
		cnt[n]++
	}
	for i := 0;i <= 500;i++{
		if (cnt[i] | 1) == cnt[i]{
			return false
		}
	}
	return true
}