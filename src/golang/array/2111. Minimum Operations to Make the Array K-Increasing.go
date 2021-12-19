package array

func lis(data []int) int {
	var l int = len(data)
	var dp []int //dp[i] 长度为i + 1的最后一个元素
	for i := 0; i < l; i++ {
		if len(dp) == 0 {
			dp = append(dp, data[i])
		} else {
			var left int = 0
			var right int = len(dp) - 1
			if data[i] >= dp[right] {
				dp = append(dp, data[i])
			} else {
				for left < right {
					mid := (left + right) / 2
					if data[i] < dp[mid] {
						right = mid
					} else if data[i] >= dp[mid] {
						left = mid + 1
					}
				}
				dp[left] = data[i]
			}
		}
	}
	return l - len(dp)
}

func KIncreasing(arr []int, k int) int {
	if k == 0 {
		return 0
	}
	var l int = len(arr)
	var res int = 0
	for i := 0; i < k; i++ {
		var data []int
		for j := i; j < l; j += k {
			data = append(data, arr[j])
		}
		res += lis(data)
	}
	return res
}

//func kIncreasing(arr []int, k int) int {
//	if k == 0 {
//		return 0
//	}
//	var l int = len(arr)
//	var dp []int = make([]int, l)
//	for i := 0; i < l; i++ {
//		dp[i] = 1
//		for j := i - k; j >= 0; j -= k {
//			if arr[i] >= arr[j] {
//				dp[i] = max_int(dp[i], 1+dp[j])
//			}
//		}
//	}
//	var cnt int = 0
//	for i := l - 1; i >= 0 && i > l-1-k; i-- {
//		var cur_max int = 0
//		for j := i; j >= 0; j -= k {
//			cur_max = max_int(cur_max, dp[j])
//		}
//		cnt += cur_max
//	}
//	return l - cnt
//}
