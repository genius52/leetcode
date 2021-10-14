package array

import "container/list"

//输入：nums = [10,2,-10,5,20], k = 2
//输出：37
//解释：子序列为 [10, 2, 5, 20]
func ConstrainedSubsetSum(nums []int, k int) int{
	var l int = len(nums)
	var q list.List //q.PushBack(nums[0])//store idx from "i - k" to "i - 1",decrease sort
	var dp []int = make([]int,l)//dp[i]:the biggest sum from nums[0] to nums[i]
	var res int = nums[0]
	for i := 0;i < l;i++{
		dp[i] = nums[i]
		if q.Len() > 0{
			dp[i] += dp[q.Front().Value.(int)]
		}
		res = max_int(res,dp[i])
		//remove the idx which dp[idx] smaller than dp[i],from small to big
		for q.Len() > 0 && dp[q.Back().Value.(int)] < dp[i]{
			q.Remove(q.Back())
		}
		//remove the idx which i - idx >= k
		if q.Len() > 0 && (i - q.Front().Value.(int)) >= k{
			q.Remove(q.Front())
		}
		if dp[i] > 0{
			q.PushBack(i)
		}

	}
	return res
}

//TLE
//func dfs_constrainedSubsetSum(nums []int,l int,pos int,last_pos int,k int,memo *[]int)int{
//	if pos >= l{
//		return 0
//	}
//	if last_pos != -1 && (pos - last_pos) > k{
//		return 0
//	}
//	if (*memo)[pos] != -2147483648{
//		return (*memo)[pos]
//	}
//	//skip current
//	var res1 int = dfs_constrainedSubsetSum(nums,l,pos + 1,last_pos,k,memo)
//	//choose current
//	var res2 int = nums[pos]
//	for i := 1;i <= k;i++{
//		if (pos + i) >= l{
//			break
//		}
//		cur := nums[pos] + dfs_constrainedSubsetSum(nums,l,pos + i,pos,k,memo)
//		if cur > res2{
//			res2 = cur
//		}
//	}
//	if res2 >= res1{
//		(*memo)[pos] = res2
//		return (*memo)[pos]
//	}
//	return res1
//}
//
//func ConstrainedSubsetSum(nums []int, k int) int {
//	var l int = len(nums)
//	var memo []int = make([]int,l)
//	for i := 0;i < l;i++{
//		memo[i] = -2147483648
//	}
//	return dfs_constrainedSubsetSum(nums,l,0,-1,k,&memo)
//}