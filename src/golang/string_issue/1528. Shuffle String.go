package string_issue

func RestoreString(s string, indices []int) string {
	var l int = len(s)
	var chars []string = make([]string,l)
	for i := 0;i < l;i++{
		chars[indices[i]] = string(s[i])
	}
	var res string
	for _,c := range chars{
		res += c
	}
	return res
}
