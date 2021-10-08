package number

func findLucky(arr []int) int {
	var record [501]int
	for _,n := range arr{
		record[n]++
	}
	var res int = -1
	for i := 500;i >= 1;i--{
		if record[i] == i{
			return i
		}
	}
	return res
}