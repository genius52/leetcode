package string_issue

func longestDecomposition(text string) int {
	var l int = len(text)
	var left int = 0
	var right int = l - 1
	last_left := 0
	last_right := l - 1
	var res int = 0
	var head int = 0
	var tail int = 0
	var sub_len int = 0
	var find bool = false
	base := 1
	for left < right{
		head = (head * 26 + int(text[left] - 'a')) % 19260817
		tail = (int(text[right] - 'a') * base + tail) % 19260817
		if head == tail{
			head = 0
			tail = 0
			sub_len = 0
			last_left = left
			last_right = right
			base = 1
			find = true
			res++
		}else{
			sub_len++
			base = base * 26 % 19260817
		}
		left++
		right--
	}
	if !find{
		return 1
	}else if last_left + 1 == last_right{
		return res * 2
	}else{
		return res * 2 + 1
	}
}