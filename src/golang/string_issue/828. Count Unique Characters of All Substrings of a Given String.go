package string_issue

//abcdbe
func UniqueLetterString(s string) int {
	var l int = len(s)
	var char_pos [26][]int
	for i := 0;i < l;i++{
		index := s[i] - 'A'
		char_pos[index] = append(char_pos[index],i)
	}
	var res int = 0
	for i := 0;i < 26;i++{
		var cur_len int = len(char_pos[i])
		if cur_len == 0{
			continue
		}
		if cur_len == 1{
			res += (char_pos[i][0] + 1) * (l - char_pos[i][0])
			continue
		}
		for j := 0;j < cur_len;j++{
			if j == 0{
				right := char_pos[i][j + 1]
				res += (char_pos[i][j] + 1) * (right - char_pos[i][j])
				res %= 1000000007
			}else if j == (cur_len - 1){
				left := char_pos[i][j - 1]
				res += (char_pos[i][j] - left) * (l - char_pos[i][j])
				res %= 1000000007
			}else{
				left := char_pos[i][j - 1]
				right := char_pos[i][j + 1]
				res += (char_pos[i][j] - left) * (right - char_pos[i][j])
				res %= 1000000007
			}
		}
	}
	return res
}