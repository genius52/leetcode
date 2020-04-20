package array

//Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.
//
//Example:
//
//Input: s = 7, nums = [2,3,1,2,,43]
//Output: 2
//Explanation: the subarray [4,3] has the minimal length under the problem constraint.
func MinSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	var dp []int = make([]int,l + 1)
	dp[0] = 0
	for i := 1;i <= l;i++{
		dp[i] = dp[i - 1] + nums[i - 1]//dp[i] = sum from 0 to i - 1
		if nums[i - 1] >= s{
			return 1
		}
	}
	if dp[l] < s{
		return 0
	}
	var begin int = 0
	var end int = 1
	var min int = l
	for end <= l{
		if dp[end] - dp[begin] >= s{
			if end - begin < min{
				min = end - begin
			}
			begin++
		}else{
			end++
		}
	}
	return min
}

func MinSubArrayLen2(s int, nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	var begin int = 0
	var end int = 0
	var min int = l
	sum := nums[begin]
	var meet bool = false
	for end < l{
		if sum >= s{
			if end - begin + 1< min{
				min = end - begin + 1
			}
			sum -= nums[begin]
			begin++
			meet = true
		}else{
			end++
			if end < l{
				sum += nums[end]
			}
		}
	}
	if !meet{
		return 0
	}
	return min
}