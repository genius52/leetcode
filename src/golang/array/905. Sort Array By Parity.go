package array

//Input: nums = [3,1,2,4]
//Output: [2,4,3,1]
//o(n) o(1)
func SortArrayByParity(nums []int) []int {
	var l int = len(nums)
	var even_idx int = 0
	var odd_idx int = 0
	for even_idx < l && odd_idx < l{
		//找到偶数
		for even_idx < l && (nums[even_idx] | 1) == nums[even_idx]{
			even_idx++
		}
		//找到奇数
		for odd_idx < l && (nums[odd_idx] | 1) != nums[odd_idx]{
			odd_idx++
		}
		if odd_idx < l && even_idx < l{
			if odd_idx < even_idx{			//如果奇数在偶数前面，则交换
				nums[even_idx],nums[odd_idx] = nums[odd_idx],nums[even_idx]
				odd_idx++
			}else{
				even_idx++
			}
		}
	}
	return nums
}
//o(n) o(n)
//func sortArrayByParity(nums []int) []int {
//	var l int = len(nums)
//	var even_idx int = 0
//	var odd_idx int = l - 1
//	var res []int = make([]int,l)
//	for i := 0;i < l;i++{
//		if (nums[i] | 1) == nums[i]{
//			res[odd_idx] = nums[i]
//			odd_idx--
//		}else{
//			res[even_idx] = nums[i]
//			even_idx++
//		}
//	}
//	return res
//}