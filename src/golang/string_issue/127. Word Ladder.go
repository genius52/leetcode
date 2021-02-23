package string_issue

import "container/list"

type word_layer struct {
	word string
	last *word_layer
}

func is_similar(w1 string,w2 string)bool{
	var l int = len(w1)
	var diff_cnt int = 0
	for i := 0;i < l;i++{
		if w1[i] != w2[i]{
			diff_cnt++
			if diff_cnt > 1{
				return false
			}
		}
	}
	return true
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var neighbours map[string][]string = make(map[string][]string)
	var dict map[string]bool = make(map[string]bool)
	for _,w := range wordList{
		dict[w] = true
	}
	delete(dict,beginWord)
	wordList = append(wordList,beginWord)
	var word_len int = len(wordList)
	for i := 0;i < word_len;i++{
		for j := 0;j < word_len;j++{
			if i == j{
				continue
			}
			if !is_similar(wordList[i],wordList[j]){
				continue
			}
			if _,ok := neighbours[wordList[i]];ok{
				neighbours[wordList[i]] = append(neighbours[wordList[i]],wordList[j])
			}else{
				neighbours[wordList[i]] = []string{wordList[j]}
			}
			if _,ok := neighbours[wordList[j]];ok{
				neighbours[wordList[j]] = append(neighbours[wordList[j]],wordList[i])
			}else{
				neighbours[wordList[j]] = []string{wordList[i]}
			}
		}
	}
	var res [][]string
	var q list.List
	var begin *word_layer = new(word_layer)
	begin.word = beginWord
	begin.last = nil
	var over bool = false
	q.PushBack(begin)
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var node *word_layer = q.Front().Value.(*word_layer)
			q.Remove(q.Front())
			for _,w := range neighbours[node.word]{
				if _,ok := dict[w];ok{
					if w == endWord{
						var cur []string = []string{endWord}
						var visit *word_layer = node
						for visit != nil{
							cur = append([]string{visit.word},cur...)
							visit = visit.last
						}
						res = append(res,cur)
						over = true
						break
					}else{
						var next *word_layer = new(word_layer)
						next.word = w
						next.last = node
						q.PushBack(next)
						delete(dict,w)
					}
				}
			}
		}
		if over{
			return res
		}
	}
	return res
}