package string_issue

import "strings"

//Input: str1 = "ABABAB", str2 = "ABAB"
//Output: "AB"
func GcdOfStrings(str1 string, str2 string) string {
	var l1 int = len(str1)
	var l2 int = len(str2)
	var longstr string
	var shortstr string
	if l1 > l2{
		longstr = str1
		shortstr = str2
	}else if l1 < l2{
		longstr = str2
		shortstr = str1
	}else{
		if str1 == str2{
			return str1
		}
		return ""
	}
	if !strings.HasPrefix(longstr,shortstr){
		return ""
	}
	for strings.HasPrefix(longstr,shortstr){
		longstr = longstr[min_int(l1,l2):]
	}
	if longstr == ""{
		return shortstr
	}
	return GcdOfStrings(shortstr,longstr)
}