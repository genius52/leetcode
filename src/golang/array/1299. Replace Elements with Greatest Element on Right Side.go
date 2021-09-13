package array

//Input: arr = [17,18,5,4,6,1]
//Output: [18,6,6,6,1,-1]
func replaceElements(arr []int) []int {
	if len(arr) == 0{
		return arr
	}
	var biggest int = arr[len(arr) - 1]
	arr[len(arr) - 1] = -1
	for i := len(arr) - 2;i >= 0;i--{
		tmp := arr[i]
		arr[i] = biggest
		if tmp > biggest{
			biggest = tmp
		}
	}
	return arr
}