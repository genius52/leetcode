package array

//		target := []int{6,4,8,1,3,2}
//		arr := []int{4,7,6,2,3,8,6,1}
//arr2 = [4,6,2,3,8,6,1]
func MinOperations1713(target []int, arr []int) int{
	var char_idx map[int]int = make(map[int]int)//char:index
	for idx,n := range target{
		char_idx[n] = idx
	}
	var arr2 []int
	for _,n := range arr{
		if _,ok := char_idx[n];ok{
			arr2 = append(arr2,n)
		}
	}
	var l1 int = len(target)
	var l2 int = len(arr2)
	if l2 == 0{
		return l1
	}
	var tail []int //tail[i] 长度为i + 1的最后一个数字
	for i := 0;i < l2;i++{
		if len(tail) == 0 || char_idx[arr2[i]] > char_idx[tail[len(tail) - 1]]{
			tail = append(tail,arr2[i])
		}else{
			var left int = 0
			var right int = len(tail) - 1
			var res int = 0 //找到第一个出现在arr2[i]前面的数字
			for left < right{
				mid := (left + right)/2
				if char_idx[arr2[i]] == char_idx[tail[mid]]{
					res = mid
					break
				}else if char_idx[arr2[i]] > char_idx[tail[mid]]{
					left = mid + 1
					res = left
				}else if char_idx[arr2[i]] < char_idx[tail[mid]]{
					right = mid
					res = right
				}
			}
			tail[res] = arr2[i]
		}
	}
	return l1 - len(tail)
}

//Classic DP solution.  TLE
//func MinOperations1713(target []int, arr []int) int{
//	var record map[int]int = make(map[int]int)//char:index
//	for idx,n := range target{
//		record[n] = idx
//	}
//	var arr2 []int
//	for _,n := range arr{
//		if _,ok := record[n];ok{
//			arr2 = append(arr2,n)
//		}
//	}
//	var l1 int = len(target)
//	var l2 int = len(arr2)
//	if l2 == 0{
//		return l1
//	}
//	var dp []int = make([]int,l2)//dp[i]: 在位置i,最长的子序列长度
//	dp[0] = 1
//	for i := 0;i < l2;i++{
//		dp[i] = 1
//		for j := 0;j < i;j++{
//			if arr2[i] != arr2[j] && record[arr2[i]] > record[arr2[j]]{
//				dp[i] = max_int(dp[i],dp[j] + 1)
//			}
//		}
//	}
//	var max_val int = 0
//	for i := 0;i < l2;i++{
//		if dp[i] > max_val{
//			max_val = dp[i]
//		}
//	}
//	return l1 - max_val
//}

//MLE
//func MinOperations1713(target []int, arr []int) int {
//	var l1 int = len(target)
//	var l2 int = len(arr)
//	var dp [][]int = make([][]int,l1 + 1)//dp[i][j]:target[0:i], arr[0:j] 最长相同子序列的长度
//	for i := 0;i <= l1;i++{
//		dp[i] = make([]int,l2 + 1)
//	}
//	for i := 1;i <= l1;i++{
//		for j := 1;j <= l2;j++{
//			if target[i - 1] == arr[j - 1]{
//				dp[i][j] = max_int(dp[i][j],dp[i - 1][j - 1] + 1)
//			}else{
//				dp[i][j] = max_int(dp[i - 1][j],dp[i][j - 1])
//			}
//		}
//	}
//	return l1 - dp[l1][l2]
//}