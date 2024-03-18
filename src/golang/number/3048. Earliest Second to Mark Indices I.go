package number

func check_earliestSecondToMarkIndices(nums []int, changeIndices []int, target int) bool {
	var free_days int = 0
	var l1 int = len(nums)
	var last_day []int = make([]int, l1)
	for i := 0; i < l1; i++ {
		last_day[i] = -1
	}
	for i := 0; i < target; i++ {
		last_day[changeIndices[i]-1] = i
	}
	for i := 0; i < l1; i++ {
		if last_day[i] == -1 {
			return false
		}
	}
	for i := 0; i < target; i++ {
		if i != last_day[changeIndices[i]-1] {
			free_days++
		} else {
			if free_days < nums[changeIndices[i]-1] {
				return false
			}
			free_days -= nums[changeIndices[i]-1]
		}
	}
	return true
}

func EarliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	var l1 int = len(nums)
	var data_idx map[int][]int = make(map[int][]int)
	var l2 int = len(changeIndices)
	var low int = 1
	var high int = len(changeIndices) //用于标记
	for i := 0; i < l2; i++ {
		data_idx[changeIndices[i]] = append(data_idx[changeIndices[i]], i)
	}
	if l2 < high {
		return -1
	}
	if len(data_idx) < l1 {
		return -1
	}
	var res int = -1
	for low <= high {
		mid := (low + high) / 2
		ret := check_earliestSecondToMarkIndices(nums, changeIndices, mid)
		if ret {
			if res == -1 || res > mid {
				res = mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if res == -1 {
		return -1
	}
	return res
}
