package number

//Input: nums = [1,2,3,2]
//Output: 4
//Explanation: The unique elements are [1,3], and the sum is 4.
func SumOfUnique(nums []int) int {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)
	var sum int = 0
	var dup_sum int = 0
	for i := 0;i < l;i++{
		if times,ok := record[nums[i]];ok{
			if(times == 1){
				dup_sum += nums[i] * 2
			}else{
				dup_sum += nums[i]
			}
			record[nums[i]]++
		}else{
			record[nums[i]] = 1
		}
		sum += nums[i]
	}
	return sum - dup_sum
}