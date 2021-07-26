package number

func countTriplets(nums []int) int {
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			if _,ok := record[nums[i] & nums[j]];ok{
				record[nums[i] & nums[j]]++
			}else{
				record[nums[i] & nums[j]] = 1
			}
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		for k,v := range record{
			if (nums[i] & k) == 0{
				res += v
			}
		}
	}
	return res
}