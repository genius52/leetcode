package array

//Input: nums = [1,2,1,2,3], k = 2
//Output: 7
//Explanation: Subarrays formed with exactly 2 different integers:
//[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
func atmostk(nums []int,k int)int{
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	var left int = 0
	var res int = 0
	for right := 0;right < l;right++{
		if _,ok := record[nums[right]];ok{
			record[nums[right]]++
		}else{
			record[nums[right]] = 1
		}
		for len(record) > k{
			record[nums[left]]--
			if record[nums[left]] == 0{
				delete(record,nums[left])
			}
			left++
		}
		res += right + 1 - left//at least one
	}
	return res
}

func SubarraysWithKDistinct(nums []int, k int) int {
	return  atmostk(nums,k) - atmostk(nums,k - 1)
}