package string_issue

func addMinimum(word string) int {
	var l int = len(word)
	var status int = 0 //a : 1, ab : 2, abc = 0
	var res int = 0
	for i := 0; i < l; i++ {
		if word[i] == 'a' {
			if status == 1 {
				res += 2
			}
			if status == 2 {
				res += 1
			}
			status = 1
		} else if word[i] == 'b' {
			if status == 2 {
				res += 2
			} else if status == 0 {
				res += 1
			}
			status = 2
		} else if word[i] == 'c' {
			if status == 0 {
				res += 2
			} else if status == 1 {
				res += 1
			}
			status = 0
		}
	}
	if status == 1 {
		res += 2
	} else if status == 2 {
		res += 1
	}
	return res
}
