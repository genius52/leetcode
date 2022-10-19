package array

func countDistinctIntegers(nums []int) int {
	var record map[int]bool = make(map[int]bool)
	for _,n := range nums{
		record[n] = true
		var n2 int = 0
		for n > 0{
			n2 *= 10
			n2 += n % 10
			n /= 10
		}
		record[n2] = true
	}
	return len(record)
}