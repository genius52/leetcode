package string_issue

import (
	"sort"
	"strings"
)

//Input: dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
//Output: "a a a a a a a a bbb baba a"
func ReplaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})
	words := strings.Split(sentence," ")
	var res string
	for _,word := range words{
		l := len(word)
		var find bool = false
		for _,d := range dictionary{
			var l_d int = len(d)
			if l_d >= l{
				continue
			}
			if strings.HasPrefix(word,d){
				if len(res) != 0{
					res += " "
				}
				find = true
				res += d
				break
			}
		}
		if !find{
			if len(res) != 0{
				res += " "
			}
			res += word
		}
	}
	return res
}