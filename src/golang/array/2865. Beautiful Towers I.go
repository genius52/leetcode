package array

func maximumSumOfHeights(maxHeights []int) int64 {
	var l int = len(maxHeights)
	var res int64 = 0
	for i := 0; i < l; i++ {
		var sum int64 = int64(maxHeights[i])
		var pre int = maxHeights[i]
		for left := i - 1; left >= 0; left-- {
			if maxHeights[left] >= pre {
				sum += int64(pre)
			} else {
				sum += int64(maxHeights[left])
				pre = maxHeights[left]
			}
		}
		pre = maxHeights[i]
		for right := i + 1; right < l; right++ {
			if maxHeights[right] >= pre {
				sum += int64(pre)
			} else {
				sum += int64(maxHeights[right])
				pre = maxHeights[right]
			}
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

//func maximumSumOfHeights(maxHeights []int) int64 {
//	var max_idx int = -1
//	var max_height int = 0
//	var l int = len(maxHeights)
//	for i := 0; i < l; i++ {
//		if maxHeights[i] > max_height {
//			max_height = maxHeights[i]
//			max_idx = i
//		}
//	}
//	var sum int64 = int64(maxHeights[max_idx])
//	var pre int = maxHeights[max_idx]
//	for i := max_idx - 1; i >= 0; i-- {
//		if maxHeights[i] >= pre {
//			sum += int64(pre)
//		} else {
//			sum += int64(maxHeights[i])
//			pre = maxHeights[i]
//		}
//	}
//	pre = maxHeights[max_idx]
//	for i := max_idx + 1; i < l; i++ {
//		if maxHeights[i] >= pre {
//			sum += int64(pre)
//		} else {
//			sum += int64(maxHeights[i])
//			pre = maxHeights[i]
//		}
//	}
//	return sum
//}
