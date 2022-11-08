package array

func MaximumSubarraySum(nums []int, k int) int64 {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)
	var res int64 = 0
	var sum int64 = 0
	var dup_cnt int = 0
	for i := 0;i < k;i++{
		sum += int64(nums[i])
		record[nums[i]]++
		if record[nums[i]] == 2{
			dup_cnt++
		}
	}
	if dup_cnt == 0{
		res = sum
	}
	for i := k;i < l;i++{
		sum -= int64(nums[i - k])
		sum += int64(nums[i])
		record[nums[i - k]]--
		record[nums[i]]++
		if nums[i - k] != nums[i]{
			if record[nums[i - k]] == 0{
				delete(record,nums[i - k])
			}
			if record[nums[i - k]] == 1{
				dup_cnt--
			}
			if record[nums[i]] == 2{
				dup_cnt++
			}
		}
		if dup_cnt == 0{
			if sum > res{
				res = sum
			}
		}
	}
	return res
}