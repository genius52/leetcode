package string_issue

import (
	"strings"
)

func NumDifferentIntegers(word string) int {
	word += "a"
	var l int = len(word)
	var begin int = 0
	var end int = 0
	var record map[string]bool = make(map[string]bool)
	for begin < l && end < l{
		if word[begin] >= 'a' && word[begin] <= 'z'{
			begin++
		}else{
			end = begin + 1
			for end < l{
				if word[end] >= '0' && word[end] <= '9'{
					end++
				}else{
					break
				}
			}
			sub := word[begin:end]
			if sub != "0"{
				record[strings.TrimLeft(sub,"0")] = true
			}else{
				record["0"] = true
			}
			begin = end
		}
	}
	return len(record)
}