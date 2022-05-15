package string_issue

import "container/list"
//func is_similar(first string,second string)bool{
//	var difference int = 0
//	for i,_ := range first{
//		if second[i] != first[i]{
//			difference++
//		}
//		if difference >= 2{
//			return false
//		}
//	}
//	return difference == 1
//}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var find bool = false
	for _,s := range wordList{
		if endWord == s{
			find = true
			break
		}
	}
	if !find{
		return 0
	}
	var trace map[string]bool = make(map[string]bool) //store unvisited words
	for _,word := range wordList{
		trace[word] = true
	}
	var q list.List
	q.PushBack(beginWord)
	var steps int = 1
	for q.Len() > 0{
		cur_step_size := q.Len()
		for i := 0;i < cur_step_size;i++{
			ele := q.Front()
			cur_word := ele.Value.(string)
			if cur_word == endWord{
				return steps
			}
			q.Remove(ele)
			delete(trace,cur_word)
			for j := 0;j < len(cur_word);j ++{
				for k := 0;k < 26;k++{
					search_word := cur_word[0:j] + string(k + 'a') + cur_word[j+1:]
					if _,ok := trace[search_word];ok{
						q.PushBack(search_word)
					}
				}
			}
		}
		steps++
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	var dict map[string]bool = make(map[string]bool)
	var find bool = false
	for _,s := range wordList{
		if s == endWord{
			find = true
		}
		dict[s] = true
	}
	if !find{
		return 0
	}
	var q list.List
	q.PushBack(beginWord)
	var steps int = 1
	for q.Len() > 0{
		var cur_len int= q.Len()
		for i := 0;i < cur_len;i++{
			ele := q.Front()
			s := ele.Value.(string)
			if s == endWord{
				return steps
			}
			q.Remove(ele)
			for j := 0;j < len(s);j ++{
				for k := 0;k < 26;k++{
					search_word := s[0:j] + string(k + 'a') + s[j+1:]
					if _,ok := dict[search_word];ok{
						q.PushBack(search_word)
						delete(dict,search_word)
					}
				}
			}
		}
		steps++
	}
	return steps
}