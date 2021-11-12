package number

//输入：nums = [3,1,4,2], p = 6
//输出：1
//解释：nums 中元素和为 10，不能被 p 整除。我们可以移除子数组 [4] ，剩余元素的和为 6 。
func MinSubarray(nums []int, p int) int {
	var l int = len(nums)
	var total int = 0
	for i := 0;i < l;i++{
		total += nums[i]
	}
	if total < p{
		return -1
	}
	if total == p{
		return 0
	}
	//I need a subarray which sum % p == target
	var target_mod int = total % p
	if target_mod == 0{
		return 0
	}
	var res int = l
	var cur_sum int = 0
	var record map[int]int = make(map[int]int)
	record[0] = -1
	for i := 0;i < l;i++{
		cur_sum += nums[i]
		var cur_mod int = cur_sum % p
		var need_mod int = (cur_mod + p - target_mod) % p
		if cur_mod >= target_mod{
			need_mod = cur_mod - target_mod
		}else{
			need_mod = cur_mod + p - target_mod
		}
		if pos,ok := record[need_mod];ok{
			if i - pos < res{
				res = i - pos
				if res == 1{
					break
				}
			}
		}
		record[need_mod] = i
	}
	if res == l{
		return -1
	}
	return res
}