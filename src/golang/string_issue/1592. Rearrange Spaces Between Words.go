package string_issue

import "strings"

func reorderSpaces(text string) string{
	var l int = len(text)
	var total_space int = 0
	var words []string = strings.Fields(text)
	var word_cnt int = len(words)
	for i := 0;i < l;i++{
		if text[i] == ' '{
			total_space++
		}
	}
	var each_space int = 0
	var mod int = 0
	if word_cnt > 1{
		each_space = total_space / (word_cnt - 1)
		mod = total_space % (word_cnt - 1)
	}else{
		each_space = 0
		mod = total_space
	}
	var span string
	for i := 0;i < each_space;i++{
		span += " "
	}
	var res string
	for i := 0;i < word_cnt;i++{
		if len(res) != 0{
			res += span
		}
		res += words[i]
	}
	for i := 0;i < mod;i++{
		res += " "
	}
	return res
}

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