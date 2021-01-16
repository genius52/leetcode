package number

//1,2,3,4  p = 3 target = 1
//1 0,0,1

//total = 10
//Input: nums = [3,1,4,2], p = 6 //整个数组多出4
//Output: 1
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
	var target int = total % p
	if target == 0{
		return 0
	}
	var res int = l
	var cur_sum int = 0
	var record map[int]int = make(map[int]int)
	record[0] = -1
	for i := 0;i < l;i++{
		cur_sum += nums[i]
		var rest int = cur_sum % p
		var need int = (rest + p - target) % p
		if pos,ok := record[need];ok{
			if i - pos < res{
				res = i - pos
				if res == 1{
					break
				}
			}
		}
		record[rest] = i
	}
	if res == l{
		return -1
	}
	return res
}