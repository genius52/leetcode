package array

//输入：nums = [1,2,2,2,5,0]
//输出：3
//解释：nums 总共有 3 种好的分割方案：
//[1] [2] [2,2,5,0]
//[1] [2,2] [2,5,0]
//[1,2] [2,2] [5,0]
func WaysToSplit(nums []int) int{
	var l int = len(nums)
	var prefix []int = make([]int,l)
	prefix[0] = nums[0]
	for i := 1;i < l;i++{
		prefix[i] = prefix[i - 1] + nums[i]
	}
	var MOD int64 = 1e9 + 7
	var res int64 = 0
	for i := 0;i < l - 2;i++{
		sum1 := prefix[i]
		rest := prefix[l - 1] - sum1
		if sum1 > rest/2{
			break
		}
		var left int = i + 1
		var right int = l - 1
		var pos1 int = left
		for left < right{
			mid := (left + right)/2
			if (prefix[mid] - prefix[i]) >= prefix[i]{
				pos1 = mid
				right = mid
			}else{
				left = mid + 1
				pos1 = left
			}
		}
		left = i + 1
		right = l - 1
		var pos2 int = left
		for left < right{
			mid := (left + right)/2
			if (prefix[l - 1] - prefix[mid]) >= (prefix[mid] - prefix[i]){
				pos2 = mid
				left = mid + 1
			}else{
				right = mid
				//pos2 = mid
			}
		}
		if pos1 > pos2{
			continue
		}
		res += int64(pos2 - pos1 + 1)
	}
	return int(res % MOD)
}
//TLE
//func waysToSplit(nums []int) int {
//	var l int = len(nums)
//	var prefix []int = make([]int,l + 1)
//	for i := 0;i < l;i++{
//		prefix[i + 1] = prefix[i] + nums[i]
//	}
//	var res int = 0
//	for i := 1;i < l - 2;i++{
//		sum1 := prefix[i]
//		rest := prefix[l] - sum1
//		if sum1 > rest/2{
//			break
//		}
//		var left_idx int = -1
//		var right_idx int = -1
//		for j := i + 1;j < l - 1;j++{
//			sum2 := prefix[j] - sum1
//			if sum2 < sum1{
//				continue
//			}
//			sum3 := prefix[l] - prefix[j]
//			if sum2 > sum3{
//				break
//			}
//			if sum1 <= sum2 && sum2 <= sum3{
//				if left_idx == -1{
//					left_idx = j
//				}else{
//					right_idx = j
//				}
//			}
//		}
//		res += right_idx - left_idx + 1
//	}
//	return res
//}