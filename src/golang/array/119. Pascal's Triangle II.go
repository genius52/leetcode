package array

//1
//11
//121
//1331
//14641
func getRow(rowIndex int) []int {
	if rowIndex == 0{
		return []int{1}
	}
	if rowIndex == 1{
		return []int{1,1}
	}
	var last []int = []int{1,1}
	var last_len int = 2
	for i := 2;i <= rowIndex;i++{
		var cur []int = make([]int,last_len + 1)
		cur[0] = 1
		cur[last_len] = 1
		for i := 1;i < last_len;i++{
			cur[i] = last[i - 1] + last[i]
		}
		last = cur
		last_len++
	}
	return last
}