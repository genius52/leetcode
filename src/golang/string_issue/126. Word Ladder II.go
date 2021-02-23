package string_issue

import "container/list"

//Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
//Output: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]

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

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	var neighbours map[string]map[string]bool = make(map[string]map[string]bool)//
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
			if _,ok := neighbours[wordList[i]];!ok{
				neighbours[wordList[i]] = make(map[string]bool)
			}
			neighbours[wordList[i]][wordList[j]] = true
			if _,ok := neighbours[wordList[j]];!ok{
				neighbours[wordList[j]] = make(map[string]bool)
			}
			neighbours[wordList[j]][wordList[i]] = true
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
		var visited map[string]bool = make(map[string]bool)
		for i := 0;i < l;i++{
			var node *word_layer = q.Front().Value.(*word_layer)
			q.Remove(q.Front())
			for w,_ := range neighbours[node.word]{
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
						visited[w] = true
					}
				}
			}
		}
		if over{
			return res
		}
		for v,_ := range visited{
			delete(dict,v)
		}
	}
	return res
}