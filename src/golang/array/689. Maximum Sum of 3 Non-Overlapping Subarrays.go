package array

//Input: nums = [1,2,1,2,6,7,5,1], k = 2
//Output: [0,3,5]
//Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
//We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.
type val_index struct{
	val int
	index []int
}

func dp_maxSumOfThreeSubarrays(nums []int,pos int,k int,pre_sum []int,total_arr int,memo [][4]val_index)(int,[]int){
	var l int = len(nums)
	if pos + k > l || total_arr == 0{
		return 0,[]int{}
	}
	if memo[pos][total_arr].val != 0{
		return memo[pos][total_arr].val,memo[pos][total_arr].index
	}
	var max_sum int = 0
	var res []int
	for i := pos;i + k  + (total_arr - 1) * k <= l;i++{
		cur_sum := pre_sum[i + k] - pre_sum[i]
		next_sum,index_arr := dp_maxSumOfThreeSubarrays(nums,i + k,k,pre_sum,total_arr - 1,memo)
		if cur_sum + next_sum > max_sum{
			max_sum = cur_sum + next_sum
			res = append([]int{i},index_arr...)
		}
	}
	memo[pos][total_arr].val = max_sum
	memo[pos][total_arr].index = res
	return max_sum,res
}

func MaxSumOfThreeSubarrays(nums []int, k int) []int {
	var l int = len(nums)
	var pre_sum []int = make([]int,l + 1)
	for i := 1;i <= l;i++{
		pre_sum[i] = nums[i - 1] + pre_sum[i - 1]
	}
	var memo [][4]val_index = make([][4]val_index,l)
	_,res := dp_maxSumOfThreeSubarrays(nums,0,k,pre_sum,3,memo)
	return res
}