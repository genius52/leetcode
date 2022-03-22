package string_issue


func MinMovesToMakePalindrome(s string) int {
	var l int = len(s)
	var data []uint8 = make([]uint8,l)
	for i := 0;i < l;i++{
		data[i] = s[i]
	}
	var res int = 0
	var left int = 0
	var right_start int = l - 1
	for left < l/2{
		var right int = right_start
		for right > left && data[right] != data[left]{
			right--
		}
		if left == right{
			res += l/2 - left
		}else{
			for j := right;j < right_start;j++{
				data[j] = data[j + 1]
				res++
			}
			data[right_start] = data[left]
			right_start--
		}
		left++
	}
	return res
}