package array

//输入：arr = [7,3,4,7], target = 7
//输出：2
//解释：尽管我们有 3 个互不重叠的子数组和为 7 （[7], [3,4] 和 [7]），但我们会选择第一个和第三个子数组，因为它们的长度和 2 是最小值。
func MinSumOfLengths(arr []int, target int) int {
	var l int = len(arr)
	var dp []int = make([]int,l)
	for i := 0;i < l;i++{
		dp[i] = 2147483647
	}
	var res int = 2147483647
	var sum int = 0
	var left int = 0
	for i := 0;i < l;i++{
		if i > 0{
			dp[i] = dp[i - 1]
		}
		sum += arr[i]
		for sum > target{
			sum -= arr[left]
			left++
		}
		if sum == target{
			if left > 0 && dp[left - 1] != 2147483647{
				cur := i - left + 1 + dp[left - 1]
				if cur < res{
					res = cur
				}
				dp[i] = min_int(i - left + 1,dp[i - 1])
			}else{
				dp[i] = i - left + 1
			}
		}
	}
	if res == 2147483647{
		return -1
	}
	return res
}

//TLE by prefix && suffix solution
//func MinSumOfLengths(arr []int, target int) int {
//	var l int = len(arr)
//	var prefix []int = make([]int,l + 1)//prefix[i]: sum from 0 to i - 1
//	for i := 0;i < l;i++{
//		prefix[i + 1] += prefix[i] + arr[i]
//	}
//	if prefix[l] < target{
//		return -1
//	}
//	var prefix_min []int = make([]int,l)//prefix_min[i]: the shortest length from 0 to i
//	for i := 0;i < l;i++{
//		prefix_min[i] = 2147483647
//	}
//	var suffix_min []int = make([]int,l)
//	for i := 0;i < l;i++{
//		suffix_min[i] = 2147483647
//	}
//	for i := 1;i <= l;i++{
//		if i > 1{
//			prefix_min[i - 1] = prefix_min[i - 2]
//		}
//		for j := i - 1;j >= 0;j--{
//			sum := prefix[i] - prefix[j]
//			if sum == target{
//				prefix_min[i - 1] = min_int(prefix_min[i - 1],i - j)
//				break
//			}
//			if sum > target{
//				break
//			}
//		}
//	}
//	for i := l - 1;i >= 0;i--{
//		if i < l - 1{
//			suffix_min[i] = suffix_min[i + 1]
//		}
//		for j := i + 1;j <= l;j++{
//			sum := prefix[j] - prefix[i]
//			if sum == target{
//				suffix_min[i] = min_int(suffix_min[i],j - i)
//				break
//			}
//			if sum > target{
//				break
//			}
//		}
//	}
//	var res int = 2147483647
//	for i := 0;i < l - 1;i++{
//		if prefix_min[i] != 2147483647 && suffix_min[i + 1] != 2147483647{
//			cur := prefix_min[i] + suffix_min[i + 1]
//			if cur < res{
//				res = cur
//			}
//		}
//	}
//	if res == 2147483647{
//		return -1
//	}
//	return res
//}