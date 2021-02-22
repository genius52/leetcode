package array

//func getRow(rowIndex int) []int {
//	if rowIndex < 0{
//		return []int{}
//	}
//	var triangle [][]int = make([][]int,rowIndex + 1)
//	triangle[0] = make([]int,1)
//	triangle[0][0] = 1
//	level := 1
//	for level <= rowIndex{
//		triangle[level] = make([]int,level + 1)
//		triangle[level][0] = 1
//		triangle[level][level] = 1
//		pos := 1
//		for pos <= level/2{
//			triangle[level][pos] = triangle[level-1][pos] + triangle[level-1][pos - 1]
//			triangle[level][level - pos] = triangle[level-1][level - pos] + triangle[level-1][level - pos - 1]
//			pos++
//		}
//		level++
//	}
//	return triangle[level-1]
//}

//Example 1:
//
//Input: rowIndex = 3
//Output: [1,3,3,1]
//Example 2:
//
//Input: rowIndex = 0
//Output: [1]
//Example 3:
//
//Input: rowIndex = 1
//Output: [1,1]

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