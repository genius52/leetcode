package string_issue

func FinalString(s string) string {
	var data []uint8
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			var l2 int = len(data)
			for j := 0; j < l2/2; j++ {
				data[j], data[l2-1-j] = data[l2-1-j], data[j]
			}
		} else {
			data = append(data, s[i])
		}
	}
	var s2 string
	for i := 0; i < len(data); i++ {
		s2 += string(data[i])
	}
	return s2
}
