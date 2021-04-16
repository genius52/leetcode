package array

//Input: [23, 2, 6, 4, 7],  k=6
//Output: True
//Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
func CheckSubarraySum(nums []int, k int) bool{
	var l int = len(nums)
	if l <= 1{
		return false
	}
	var record map[int]int = make(map[int]int)
	record[0] = -1
	var sum int = 0
	for i := 0;i < l;i++{
		sum += nums[i]
		mod := sum % k
		if i == 0{
			if _,ok := record[mod];!ok{
				record[mod] = i
			}
			continue
		}
		if sum == k {
			return true
		}
		if pre_index,ok := record[mod];ok{
			if i - pre_index > 1{
				return true
			}
		}else{
			record[mod] = i
		}
	}
	return false
}