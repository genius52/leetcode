package string_issue

//Input: s = "abcdddeeeeaabbbcd"
//Output: [[3,5],[6,9],[12,14]]
//Explanation: The large groups are "ddd", "eeee", and "bbb".
func largeGroupPositions(s string) [][]int {
	var l int = len(s)
	var res [][]int
	var left int = 0
	for left < l{
		var right int = left
		for right < l && s[left] == s[right]{
			right++
		}
		distance := right - left
		if distance >= 3{
			res = append(res,[]int{left,right - 1})
		}
		left = right
	}
	return res
}