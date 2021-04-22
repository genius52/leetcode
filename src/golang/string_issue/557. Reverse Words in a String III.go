package string_issue

func reverseWords(s string) string {
	var l int = len(s)
	var left int = 0
	var res string
	for left < l{
		var visit int = left
		for visit != l && s[visit] != ' '{
			visit++
		}
		var right int = visit - 1
		for i := right;i >= left;i--{
			res += string(s[i])
		}
		if visit != l{
			res += " "
		}
		left = visit + 1
	}
	return res
}