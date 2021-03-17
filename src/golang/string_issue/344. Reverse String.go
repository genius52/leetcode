package string_issue

func reverseString(s []byte)  {
	var l int = len(s)
	for i := 0;i < l/2;i++{
		s[i],s[l - 1 - i] = s[l - 1 - i],s[i]
	}
}