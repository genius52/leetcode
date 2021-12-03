package number

func TupleSameProduct(nums []int) int {
	var l int = len(nums)
	if l < 4{
		return 0
	}
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			product := nums[i] * nums[j]
			record[product]++
		}
	}
	var res int = 0
	for _,count := range record{
		if count == 1{
			continue
		}
		res +=  4 * count * (count - 1)
	}
	return res
}