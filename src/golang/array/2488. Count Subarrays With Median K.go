package array

//Input: nums = [3,2,1,4,5], k = 4
//Output: 3
//Explanation: The subarrays that have a median equal to 4 are: [4], [4,5] and [1,4,5].
func CountSubarrays2488(nums []int, k int) int {
	var l int = len(nums)
	var find_idx int = -1
	for i := 0;i < l;i++{
		if nums[i] == k{
			find_idx = i
			break
		}
	}
	var res int = 0
	var left_big_cnt int = 0
	var big_minus_small_cnt map[int]int = make(map[int]int)
	for i := find_idx - 1;i >= 0;i--{
		if nums[i] > k{
			left_big_cnt++
		}
		diff := left_big_cnt - ((find_idx - i) - left_big_cnt)
		if diff == 0 || diff == 1{
			res++
		}
		big_minus_small_cnt[diff]++
	}
	var right_big_cnt int = 0
	for i := find_idx + 1;i < l;i++{
		if nums[i] > k{
			right_big_cnt++
		}
		right_small_cnt := i - find_idx - right_big_cnt
		if right_big_cnt == right_small_cnt || right_big_cnt - right_small_cnt == 1{
			res++
		}
		require1 := right_small_cnt - right_big_cnt
		if val,ok := big_minus_small_cnt[require1];ok{
			res += val
		}
		require2 := right_small_cnt - right_big_cnt + 1
		if val,ok := big_minus_small_cnt[require2];ok{
			res += val
		}
	}
	return res + 1
}