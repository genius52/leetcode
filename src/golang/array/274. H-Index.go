package array

import "sort"

//Input: citations = [3,0,6,1,5]
//Output: 3
func HIndex(citations []int) int {
	sort.Ints(citations)
	var l int  = len(citations)
	i := 0
	for i < l{
		if(citations[i] >= (l - i)){
			break;
		}
		i++
	}
	return l - i
}