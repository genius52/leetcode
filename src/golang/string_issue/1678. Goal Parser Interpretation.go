package string_issue

import "strings"

//Input: command = "(al)G(al)()()G"
//Output: "alGalooG"
func interpret(command string) string {
	var res string = strings.ReplaceAll(command,"()","o")
	res = strings.ReplaceAll(res,"(al)","al")
	return res
}