package number

import "strconv"

//Input: nums is [1, 1, 1, 1, 1], S is 3.
//Output: 5
//Explanation:
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
func dp_findTargetSumWays(nums []int,cur_pos int,cur_sum int,S int,memo map[string]int)int{
	key := strconv.Itoa(cur_sum) + "," + strconv.Itoa(cur_pos)
	if result,ok := memo[key];ok{
		return result
	}
	if cur_pos > len(nums){
		return 0
	}
	if cur_pos == len(nums){
		if cur_sum == S{
			return 1
		}else{
			return 0
		}
	}
	cnt := dp_findTargetSumWays(nums,cur_pos + 1,cur_sum + nums[cur_pos],S,memo)
	cnt += dp_findTargetSumWays(nums,cur_pos + 1,cur_sum - nums[cur_pos],S,memo)
	memo[key] = cnt
	return cnt
}

func FindTargetSumWays(nums []int, S int) int {
	var memo map[string]int = make(map[string]int)
	return dp_findTargetSumWays(nums,0,0,S,memo)
}