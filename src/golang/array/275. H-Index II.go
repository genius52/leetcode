package array

//Input: citations = [0,1,3,5,6]
//Output: 3
func HIndex2(citations []int) int {
	var l int = len(citations)
	begin := 0
	end := l - 1
	for begin <= end{
		var mid int = begin + (end - begin)/2
		if citations[mid] >= (l - mid){
			if mid - 1 >= 0 && citations[mid - 1] < (l - mid + 1){
				return l - mid
			}
			end = mid - 1
		}else{
			if (mid + 1) < l && citations[mid + 1] >= (l - mid - 1){
				return l - mid - 1
			}
			begin = mid + 1
		}
	}
	return l - begin
}