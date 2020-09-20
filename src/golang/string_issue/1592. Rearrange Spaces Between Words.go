package string_issue

import "strings"

func ReorderSpaces(text string) string {
	var l int = len(text)
	var space_cnt int = 0
	for i := 0;i < l;i++{
		if text[i] == ' '{
			space_cnt++
		}
	}
	var words []string = strings.Fields(text)
	var word_cnt int = len(words)
	if word_cnt <= 1{
		var tail string
		for i := 0;i < space_cnt;i++{
			tail += " "
		}
		return words[0] + tail
	}
	var span int = space_cnt / (word_cnt - 1)
	var tail_space string
	if space_cnt % (word_cnt - 1) != 0{
		//span++
		for i := 0;i < space_cnt - span * (word_cnt - 1);i++{
			tail_space += " "
		}
	}
	var space string
	for i := 0;i < span;i++{
		space += " "
	}
	var res string
	for i := 0;i < word_cnt;i++{
		if len(res) != 0{
			res += space
		}
		res += words[i]
		if i == word_cnt - 1{
			res += tail_space
			break
		}
	}
	return res
}