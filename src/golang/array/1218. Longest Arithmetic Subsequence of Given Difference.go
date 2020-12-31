package array

//Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2
//Output: 4
//Explanation: The longest arithmetic subsequence is [7,5,3,1].
//[4,12,10,0,-2,7,-8,9,-9,-12,-12,8,8]
//0
//{4,12,10,0,-2,7,-8,9,-9,-12,-12,8,8}
func LongestSubsequence(arr []int, difference int) int{
	var l int = len(arr)
	var dp map[int]int = make(map[int]int)//dp[i]: 从0到i符合要求的的最大长度
	var max_len int = 1
	for i := 0;i < l;i++ {
		if n,ok := dp[arr[i] - difference];ok{
			dp[arr[i]] = 1 + n
		}
		if _,ok := dp[arr[i]];!ok{
			dp[arr[i]] = 1
		}
		if dp[arr[i]] > max_len{
			max_len = dp[arr[i]]
		}
	}
	return max_len
}

//func LongestSubsequence(arr []int, difference int) int {
//	var l int = len(arr)
//	var dp []int = make([]int,l)//dp[i]: 从0到i符合要求的的最大长度
//	for i := 0;i < l;i++{
//		dp[i] = 1
//	}
//	var max_len int = 1
//	for i := 1;i < l;i++{
//		for j := 0;j < i;j++{
//			if arr[i] - arr[j] == difference{
//				dp[i] = dp[j] + 1
//			}
//			if dp[i] > max_len{
//				max_len = dp[i]
//			}
//		}
//	}
//	return max_len
//}