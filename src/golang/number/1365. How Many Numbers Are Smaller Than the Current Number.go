package number

func smallerNumbersThanCurrent(nums []int) []int {
	var l int = len(nums)
	var record [101]int
	for i := 0;i < l;i++{
		record[nums[i]]++
	}
	for i := 1;i <= 100;i++{
		record[i] += record[i - 1]
	}
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		if nums[i] == 0{
			res[i] = 0
		}else{
			res[i] = record[nums[i] - 1]
		}
	}
	return res
}