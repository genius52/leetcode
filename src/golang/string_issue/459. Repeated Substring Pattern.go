package string_issue

func repeatedSubstringPattern(s string) bool {
	var l int = len(s)
	for sub_len := 1;sub_len <= l/2;{
		ss := s[:sub_len]
		var equal bool = true
		for i := sub_len;i < l;i += sub_len{
			if s[i:i + sub_len] != ss{
				equal = false
				break
			}
		}
		if equal{
			return true
		}
		sub_len++
		for sub_len <= l/2 && l % sub_len != 0{
			sub_len++
		}
	}
	return false
}