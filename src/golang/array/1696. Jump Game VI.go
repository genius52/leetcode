package array

//Input: nums = [10,-5,-2,4,0,3], k = 3
//Output: 17
//Explanation: You can choose your jumps forming the subsequence [10,4,3]
//optimized dp solution
func dp_maxResult(nums []int,l int,pos int,k int,momo map[int]int)int{
	if pos >= l{
		return 0
	}
	if pos == l - 1{
		return nums[l - 1]
	}
	if res,ok := momo[pos];ok{
		return res
	}
	var max_score int = -2147483648
	for step := 1;step <= k && step + pos  < l;step++{
		cur_val := nums[pos] + dp_maxResult(nums,l,pos + step,k,momo)
		if cur_val > max_score{
			max_score = cur_val
		}
		if nums[pos + step] >= 0{
			break
		}
	}
	momo[pos] = max_score
	return max_score
}

func MaxResult(nums []int, k int) int {
	var l int = len(nums)
	var memo map[int]int = make(map[int]int)
	return dp_maxResult(nums,l,0,k,memo)
}