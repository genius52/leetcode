package number

func numIdenticalPairs(nums []int) int {
	var record [101]int
	for _,n := range nums{
		record[n]++
	}
	var res int = 0
	for i := 1;i <= 100;i++{
		if record[i] > 1{
			res += record[i] * (record[i] - 1)/2
		}
	}
	return res
}