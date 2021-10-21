package array

//Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
//
//Return the array in the form [x1,y1,x2,y2,...,xn,yn].

//Input: nums = [2,5,1,3,4,7], n = 3
//Output: [2,3,5,4,1,7]
func shuffle(nums []int, n int) []int {
	var res []int = make([]int,n * 2)
	var left int = 0
	var right int = n
	for i := 0;i < n * 2;i++{
		if (i | 1) != i{
			res[i] = nums[left]
			left++
		}else{
			res[i] = nums[right]
			right++
		}
	}
	return res
}