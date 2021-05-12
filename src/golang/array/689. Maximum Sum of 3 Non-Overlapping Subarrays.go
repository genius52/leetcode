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

type val_start struct{
	val int
	start int
}

func MaxSumOfThreeSubarrays2(nums []int, k int) []int{
	var l int = len(nums)
	var left_sum int = 0
	for i := 0;i < k;i++{
		left_sum += nums[i]
	}
	var left_max []val_start = make([]val_start,l)
	left_max[k - 1].val = left_sum
	left_max[k - 1].start = 0
	for i := k;i + k * 2 < l;i++{
		cur_sum := left_sum - nums[i - k] + nums[i]
		left_sum = cur_sum
		if left_sum > left_max[i - 1].val{
			left_max[i].val = left_sum
			left_max[i].start = i - k + 1
		}else{
			left_max[i] = left_max[i - 1]
		}
	}
	var right_max []val_start = make([]val_start,l)
	var right_sum int = 0
	for i := 0;i < k;i++{
		right_sum += nums[l - 1 - i]
	}
	right_max[l - k].val = right_sum
	right_max[l - k].start = l - k
	for i := l - 1 - k;i - k * 2 >= 0;i--{
		cur_sum := right_sum - nums[i + k] + nums[i]
		right_sum = cur_sum
		if right_sum >= right_max[i + 1].val{
			right_max[i].val = right_sum
			right_max[i].start = i
		}else{
			right_max[i] = right_max[i + 1]
		}
	}

	var mid_sum int = 0
	for i := k;i < k * 2;i++{
		mid_sum += nums[i]
	}
	var max_sum int = left_max[k - 1].val + mid_sum + right_max[k * 2].val
	var res []int = []int{0,k,right_max[k * 2].start}
	for i := k + 1;i <= l - k * 2;i++{
		mid_sum = mid_sum - nums[i - 1] + nums[i - 1 + k]
		cur_sum := left_max[i - 1].val + mid_sum  + right_max[i + k].val
		if cur_sum > max_sum{
			max_sum = cur_sum
			res = []int{left_max[i - 1].start,i,right_max[i + k].start}
		}
	}
	return res
}