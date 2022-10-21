package array

func containsNearbyDuplicate(nums []int, k int) bool {
	var record map[int]int = make(map[int]int)
	for index,v := range nums{
		if pos,ok := record[v];ok{
			if index - pos <= k{
				return true
			}else{
				record[v] = index
			}
		}else{
			record[v] = index
		}
	}
	return false
}