package array

//Input: nums = [1,3]
//Output: 6
//Explanation: The 4 subsets of [1,3] are:
//- The empty subset has an XOR total of 0.
//- [1] has an XOR total of 1.
//- [3] has an XOR total of 3.
//- [1,3] has an XOR total of 1 XOR 3 = 2.
//0 + 1 + 3 + 2 = 6
func combine_subsetXORSum(nums []int,pos int,target_len int,pre_sum int)int{
	var l int = len(nums)
	if target_len == 0{
		return pre_sum
	}
	if pos >= l{
		return 0
	}
	return combine_subsetXORSum(nums,pos + 1,target_len - 1,pre_sum ^ nums[pos]) +
			combine_subsetXORSum(nums,pos + 1,target_len,pre_sum)
}

func subsetXORSum(nums []int) int {
	var l int = len(nums)
	var res int = 0
	for i := 1;i <= l;i++{
		res += combine_subsetXORSum(nums,0,i,0)
	}
	return res
}

//func recrusive_subsetXORSum(nums []int,pos int)int{
//	var l int = len(nums)
//	if pos >= l{
//		return 0
//	}
//	var res int = 0
//	var sum int = 0
//	for right := pos;right < l;right++{
//		sum = sum ^ nums[right]
//		res += sum
//	}
//	return res
//}
//
//func subsetXORSum(nums []int) int {
//	var l int = len(nums)
//	var res int = 0
//	for i := 0;i < l;i++{
//		res += recrusive_subsetXORSum(nums,i)
//	}
//	return res
//}