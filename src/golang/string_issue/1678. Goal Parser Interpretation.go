package string_issue

import "strings"

//Input: command = "(al)G(al)()()G"
//Output: "alGalooG"
func interpret(command string) string {
	var res string = strings.ReplaceAll(command,"()","o")
	res = strings.ReplaceAll(res,"(al)","al")
	return res
}

func interpret2(command string) string{
	var res strings.Builder
	var l int = len(command)
	for i := 0;i < l;{
		if command[i] == 'G'{
			res.WriteByte('G')
			i++
		}else if command[i] == '(' && command[i + 1] == ')'{
			res.WriteString("o")
			i += 2
		}else{
			res.WriteString("al")
			i += 4
		}
	}
	return res.String()
}