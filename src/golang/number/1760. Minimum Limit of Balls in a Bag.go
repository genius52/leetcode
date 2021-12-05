package number

func MinimumSize(nums []int, maxOperations int) int {
	var l int = len(nums)
	var max_val int = 0
	for i := 0; i < l; i++ {
		if nums[i] > max_val {
			max_val = nums[i]
		}
	}
	var low int = 1
	var high int = max_val
	for low < high {
		mid := (low + high) / 2
		cnt := 0
		for i := 0; i < l; i++ {
			if nums[i] > mid {
				cnt += nums[i]/mid - 1
				if nums[i]%mid != 0 {
					cnt++
				}
			}
		}
		if cnt > maxOperations {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
