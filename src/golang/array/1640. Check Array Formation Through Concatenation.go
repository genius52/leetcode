package array

func CanFormArray(arr []int, pieces [][]int) bool {
	var l1 int = len(arr)
	var l2 int = len(pieces)
	var index int = 0
	for index < l1{
		var find bool = false
		for j := 0;j < l2;j++{
			var k int = 0
			for k < len(pieces[j]){
				if pieces[j][k] == 0{
					k++
					continue
				}
				if pieces[j][k] == arr[index]{
					find = true
					pieces[j][k] = 0
				}
				break
			}
			if find{
				break
			}
		}
		if !find{
			return false
		}
		index++
	}
	return true
}