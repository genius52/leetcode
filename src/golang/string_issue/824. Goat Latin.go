package string_issue

import "strings"

//Input: sentence = "The quick brown fox jumped over the lazy dog"
//Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
func toGoatLatin(S string) string {
	var data []string = strings.Split(S," ")
	var l int = len(data)
	var res string
	var a string = "aa"
	for i := 0;i < l;i++{
		if len(res) > 0{
			res += " "
		}
		if data[i][0] == 'a' || data[i][0] == 'e' || data[i][0] == 'i'|| data[i][0] == 'o'|| data[i][0] == 'u' ||
			data[i][0] == 'A' || data[i][0] == 'E' || data[i][0] == 'I'|| data[i][0] == 'O'|| data[i][0] == 'U'{
			res += data[i] + "m" + a

		}else{
			res += data[i][1:] + string(data[i][0]) + "m" + a
		}
		a += "a"
	}
	return res
}