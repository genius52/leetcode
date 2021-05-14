package string_issue

func toLowerCase(str string) string {
	var res string
	for _,c := range str{
		if c >= 'A' && c <= 'Z'{
			res += string(c + 32)
		}else{
			res += string(c)
		}
	}
	return res
}