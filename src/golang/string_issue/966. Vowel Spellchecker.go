package string_issue

import "strings"

//Input: wordlist = ["KiTe","kite","hare","Hare"],
//queries = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"]
//Output: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]

//当查询完全匹配单词列表中的某个单词（区分大小写）时，应返回相同的单词。
//当查询匹配到大小写问题的单词时，您应该返回单词列表中的第一个这样的匹配项。
//当查询匹配到元音错误的单词时，您应该返回单词列表中的第一个这样的匹配项。
//如果该查询在单词列表中没有匹配项，则应返回空字符串。
func Spellchecker(wordlist []string, queries []string) []string {
	var origin map[string]bool = make(map[string]bool)
	var lower_origin map[string]string = make(map[string]string)
	var wrong_origin map[string]string = make(map[string]string)
	for _,word := range wordlist{
		origin[word] = true
		s2 := strings.ToLower(word)
		if _,ok := lower_origin[s2];!ok{
			lower_origin[s2] = word
		}
		var wrong string
		for _,c := range s2{
			if c == 'a' || c == 'e' || c == 'i' ||c == 'o' ||c == 'u' ||
				c == 'A' || c == 'E' || c == 'I' ||c == 'O' ||c == 'U'{
				wrong += "a"
			}else{
				wrong += string(c)
			}
		}
		if _,ok := wrong_origin[wrong];!ok{
			wrong_origin[wrong] = word
		}
	}
	var l int = len(queries)
	var res []string = make([]string,l)
	for i := 0;i < l;i++{
		if _,ok := origin[queries[i]];ok{
			res[i] = queries[i]
		}else{
			s2 := strings.ToLower(queries[i])
			if data,ok2 := lower_origin[s2];ok2{//相同字符，大小写不一致
				res[i] = data
			}else{
				//s1 := []rune(s)
				var wrong string
				for _,c := range s2{
					if c == 'a' || c == 'e' || c == 'i' ||c == 'o' ||c == 'u' ||
						c == 'A' || c == 'E' || c == 'I' ||c == 'O' ||c == 'U'{
						wrong += "a"
					}else{
						wrong += string(c)
					}
				}
				if _,ok := wrong_origin[wrong];ok{
					res[i] = wrong_origin[wrong]
				}
			}
		}
	}
	return res
}