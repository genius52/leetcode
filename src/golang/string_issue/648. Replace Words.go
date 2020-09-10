package string_issue

import (
	"sort"
	"strings"
)

//Input: dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
//Output: "a a a a a a a a bbb baba a"
type sortString []string

func (s sortString) Less(i, j int) bool {
	if len(s[i]) < len(s[j]){
		return true
	}
	return false
}

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortString) Len() int {
	return len(s)
}

func ReplaceWords(dictionary []string, sentence string) string {
	sort.Sort(sortString(dictionary))
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