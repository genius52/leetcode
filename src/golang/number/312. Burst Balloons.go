package number

//Input: [3,1,5,8]
//Output: 167
//Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
//func dp_maxCoins(nums []int,memo [][]int)int{
//	if len(nums) < 0 {
//		return 0
//	}
//	var res int = 0
//	for mid := 0; mid < len(nums);mid++{
//		var left_num int = 1
//		var right_num int = 1
//		if mid != 0{
//			left_num = nums[mid-1]
//		}
//		if mid != len(nums) - 1{
//			right_num = nums[mid + 1]
//		}
//		cur_val := left_num * nums[mid] * right_num
//		var new_nums []int = make([]int,len(nums))
//		copy(new_nums,nums)
//		new_nums = append(new_nums[:mid],new_nums[mid+1:]...)
//		rest_val := dp_maxCoins(new_nums,memo)
//		sum := rest_val + cur_val
//		if sum > res{
//			res = sum
//		}
//	}
//	return res
//}
//
//func maxCoins(nums []int) int {
//	var record [][]int = make([][]int,len(nums))
//	for i := 0;i < len(nums);i++{
//		record[i] = make([]int,len(nums))
//	}
//	return dp_maxCoins(nums,record)
//}

func MaxCoins312(nums []int) int {
	new_len := len(nums) + 2
	var nums2 []int = make([]int,new_len)
	for i,n := range nums{
		nums2[i + 1] = n
	}
	nums2[0] = 1
	nums2[new_len - 1] = 1
	var dp [][]int = make([][]int,new_len)	//dp[i][j] means  maximum coins if we burst ballons from postion i to j
	for i := 0;i < new_len;i++{
		dp[i] = make([]int,new_len)
	}
	for span := 1; span <= len(nums);span++{
		for left := 1;left <= len(nums) - span + 1;left++{//ignore left = 0
			right := left + span - 1
			for k := left;k <= right;k++{
				dp[left][right] = max_int(dp[left][right],nums2[left - 1] * nums2[k] * nums2[right + 1] + dp[left][k - 1] + dp[k + 1][right])
			}
		}
	}
	return dp[1][len(nums)]
}