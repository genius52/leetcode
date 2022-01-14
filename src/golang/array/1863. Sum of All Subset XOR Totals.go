package array

//Input: nums = [1,3]
//Output: 6
//Explanation: The 4 subsets of [1,3] are:
//- The empty subset has an XOR total of 0.
//- [1] has an XOR total of 1.
//- [3] has an XOR total of 3.
//- [1,3] has an XOR total of 1 XOR 3 = 2.
//0 + 1 + 3 + 2 = 6
func dfs_subsetXORSum(nums []int, l int, pos int, sum int) int {
	if pos >= l {
		return sum
	}
	return dfs_subsetXORSum(nums, l, pos+1, sum^nums[pos]) + dfs_subsetXORSum(nums, l, pos+1, sum)
}

func SubsetXORSum(nums []int) int {
	return dfs_subsetXORSum(nums, len(nums), 0, 0)
}

//func combine_subsetXORSum(nums []int,pos int,target_len int,pre_sum int)int{
//	var l int = len(nums)
//	if target_len == 0{
//		return pre_sum
//	}
//	if pos >= l{
//		return 0
//	}
//	return combine_subsetXORSum(nums,pos + 1,target_len - 1,pre_sum ^ nums[pos]) +
//			combine_subsetXORSum(nums,pos + 1,target_len,pre_sum)
//}
//
//func subsetXORSum(nums []int) int {
//	var l int = len(nums)
//	var res int = 0
//	for i := 1;i <= l;i++{
//		res += combine_subsetXORSum(nums,0,i,0)
//	}
//	return res
//}
