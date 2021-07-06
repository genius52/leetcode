package string_issue

//Input: s = "Test1ng-Leet=code-Q!"
//Output: "Qedo1ct-eeLg=ntse-T!"
func ReverseOnlyLetters(s string) string {
	var l int = len(s)
	var data []uint8 = make([]uint8,l)
	var visit int = 0
	for visit < l{
		data[visit] = s[visit]
		visit++
	}
	var left int = 0
	var right int = l - 1
	for left < right{
		for left < l && (data[left] < 'a' || data[left] > 'z') && (data[left] < 'A' || data[left] > 'Z') {
			left++
		}
		for right >= left && (data[right] < 'a' || data[right] > 'z') && (data[right] < 'A' || data[right] > 'Z') {
			right--
		}
		if left < right {
			data[left],data[right] = data[right],data[left]
		}
		left++
		right--
	}
	var res string
	for i := 0;i < l;i++{
		res += string(data[i])
	}
	return res
}