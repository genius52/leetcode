package array

func minDeletion(nums []int) int {
	var l int = len(nums)
	var cnt int = 0
	var idx int = 0
	for idx < l - 1{
		if nums[idx] == nums[idx + 1]{
			cnt++
			idx++
		}else{
			idx += 2
		}
	}
	if (l - idx) | 1 == l - idx{
		cnt++
	}
	return cnt
}