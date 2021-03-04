package number

func containsDuplicate(nums []int) bool {
	var record map[int]bool = make(map[int]bool)
	for _,n := range nums{
		if _,ok := record[n];ok{
			return true
		}
		record[n] = true
	}
	return false
}