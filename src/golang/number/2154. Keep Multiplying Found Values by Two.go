package number

func findFinalValue(nums []int, original int) int {
	var record [1001]int
	for _,n := range nums{
		record[n]++
	}
	for original <= 1000 && record[original] > 0{
		original *= 2
	}
	return original
}