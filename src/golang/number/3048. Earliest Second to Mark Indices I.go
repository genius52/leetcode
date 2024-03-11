package number

func check_earliestSecondToMarkIndices(nums []int, changeIndices []int, target int) bool {
	//var l1 int = len(nums)
	//var l2 int = len(changeIndices)
	//var visited []bool = make([]bool, l1)
	//for i := l2 - 1; i >= 0; i-- {
	//
	//}
	return true
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	var l1 int = len(nums)
	//var visited []bool = make([]bool, l1)
	var data_idx map[int][]int = make(map[int][]int)
	var l2 int = len(changeIndices)
	var low int = 1
	var high int = l1 //用于标记
	for i := 0; i < l2; i++ {
		if changeIndices[i] > high {
			high = changeIndices[i]
		}
		data_idx[changeIndices[i]] = append(data_idx[changeIndices[i]], i)
	}
	if l2 < high {
		return -1
	}
	if len(data_idx) < l1 {
		return -1
	}
	//var over int = high + 1
	var res int = -1
	for low < high {
		mid := (low + high) / 2
		ret := check_earliestSecondToMarkIndices(nums, changeIndices, mid)
		if !ret {
			high = mid - 1
		} else {
			res = mid
			low = mid + 1
		}
	}
	if res == -1 {
		return -1
	}
	return res
}
