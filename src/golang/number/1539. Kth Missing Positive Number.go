package number

//Input: arr = [2,3,4,7,11], k = 5
//Output: 9
func FindKthPositive(arr []int, k int) int {
	var l int = len(arr)
	var pos int = 0
	var number int = 1
	for pos < l{
		for arr[pos] != number{
			k--
			if(k == 0){
				return number
			}
			number++
		}
		number++
		pos++
	}
	return arr[l - 1] + k
}