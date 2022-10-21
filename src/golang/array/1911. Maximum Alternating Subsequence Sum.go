package array

//Input: nums = [4,2,5,3]
//Output: 7
//Explanation: It is optimal to choose the subsequence [4,2,5] with alternating sum (4 + 5) - 2 = 7.

//Input: nums = [5,6,7,8]
//Output: 8
//Explanation: It is optimal to choose the subsequence [8] with alternating sum 8.

//Input: nums = [6,2,1,2,4,5]
//Output: 10
//Explanation: It is optimal to choose the subsequence [6,1,5] with alternating sum (6 + 5) - 1 = 10.

func dp_maxAlternatingSum(nums []int,l int,pos int,is_even int,memo [][2]int64)int64{
	if pos >= l{
		return 0
	}
	if memo[pos][is_even] != 0{
		return memo[pos][is_even]
	}
	var res int64 = 0
	if is_even == 1{
		//add current && skip current
		res = max_int64(int64(nums[pos]) + dp_maxAlternatingSum(nums,l,pos + 1,0,memo),dp_maxAlternatingSum(nums,l,pos + 1,1,memo))
	}else{
		//minus current && not sell
		res = max_int64(int64(-nums[pos]) + dp_maxAlternatingSum(nums,l,pos + 1,1,memo),
			dp_maxAlternatingSum(nums,l,pos + 1,0,memo))
	}
	memo[pos][is_even] = res
	return res;
}

func maxAlternatingSum(nums []int) int64 {
	var l int = len(nums)
	var memo [][2]int64 = make([][2]int64,l)
	return dp_maxAlternatingSum(nums,l,0,1,memo)
}

func maxAlternatingSum2(nums []int) int64 {
	var l int = len(nums)
	var profit int64 = 0
	var pre int = 0
	for i := 0;i < l;i++{
		if nums[i] > pre{
			profit += int64(nums[i] - pre)
		}
		pre = nums[i]
	}
	return profit
}