package array

//Input: nums = [3,1,3,4,3], k = 6
//Output: 1
//Explanation: Starting with nums = [3,1,3,4,3]:
//- Remove the first two 3's, then nums = [1,4,3]
//There are no more pairs that sum up to 6, hence a total of 1 operation.
func maxOperations(nums []int, k int) int {
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	for i := 0;i < l;i++{
		record[nums[i]]++
	}
	var res int = 0
	for i := 0;i < l;i++{
		target := k - nums[i]
		if cnt,ok := record[target];ok{
			if cnt <= 0{
				continue
			}
			if target == nums[i]{
				if cnt >= 2{
					record[target] -= 2
					res++
				}
			}else{
				if cnt2,ok2 := record[nums[i]];ok2{
					if cnt2 > 0{
						res++
						record[target]--
						record[nums[i]]--
					}
				}
			}
		}
	}
	return res
}