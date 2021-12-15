package array

//输入：n = 4, index = 2,  maxSum = 6
//输出：2
//解释：数组 [1,1,2,1] 和 [1,2,2,1] 满足所有条件。不存在其他在指定下标处具有更大值的有效数组。
func MaxValue(n int, index int, maxSum int) int {
	var low int = 1
	var high int = maxSum
	var res int = 1
	var right_len int = n - index - 1
	for low <= high {
		mid := (low + high) / 2
		var cur_sum int = 0
		if mid-index > 0 {
			cur_sum += (mid-index)*(index+1) + ((index+1)*index)/2
		} else {
			cur_sum += 1*mid + (mid*(mid-1))/2
			cur_sum += (index - mid + 1)
		}
		if right_len > 0 {
			if mid > right_len {
				cur_sum += (mid-right_len)*right_len + right_len*(right_len-1)/2
			} else {
				cur_sum += 1*(mid-1) + ((mid-1)*(mid-2))/2
				cur_sum += (right_len - mid + 1)
			}
		}
		if cur_sum > maxSum {
			high = mid - 1
			res = mid - 1
		} else if cur_sum < maxSum {
			low = mid + 1
			res = mid
		} else {
			res = mid
			break
		}
	}
	return res
}
