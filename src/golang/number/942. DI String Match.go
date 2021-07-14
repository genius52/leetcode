package number

//Input: s = "DDI"
//Output: [3,2,0,1]
func diStringMatch(S string) []int {
	var l int = len(S)
	var res []int = make([]int,l + 1)
	var low int = 0
	var high int = l
	for i := 0;i < l;i++{
		if S[i] == 'I'{
			res[i] = low
			low++
		}else{
			res[i] = high
			high--
		}
	}
	res[l] = low
	return res
}