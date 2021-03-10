package string_issue

import "strings"

//Input: pattern = "abba", s = "dog cat cat dog"
//Output: true

//Input: pattern = "abba", s = "dog cat cat fish"
//Output: false
func WordPattern(pattern string, s string) bool {
	words := strings.Split(s," ")
	var l1 int = len(pattern)
	var l2 int = len(words)
	if l1 != l2{
		return false
	}
	var char_graph map[uint8][]int = make(map[uint8][]int)
	var word_graph map[string][]int = make(map[string][]int)
	var char_word map[uint8]string = make(map[uint8]string)
	for i := 0;i < l1;i++{
		char_word[pattern[i]] = words[i]
		if _,ok := char_graph[pattern[i]];ok{
			char_graph[pattern[i]] = append(char_graph[pattern[i]],i)
		}else{
			char_graph[pattern[i]] = []int{i}
		}
		if _,ok := word_graph[words[i]];ok{
			word_graph[words[i]] = append(word_graph[words[i]],i)
		}else{
			word_graph[words[i]] = []int{i}
		}
	}

	for c, pos1 := range char_graph{
		var cur_s string = char_word[c]
		var pos2 []int = word_graph[cur_s]
		var len_pos1 int = len(pos1)
		var len_pos2 int = len(pos2)
		if len_pos1 != len_pos2{
			return false
		}
		for i := 0;i < len_pos1;i++{
			if pos1[i] != pos2[i]{
				return false
			}
		}
	}
	return true
}