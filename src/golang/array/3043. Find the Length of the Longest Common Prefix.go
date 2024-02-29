package array

import (
	"container/list"
	"strconv"
)

type Num_trie struct {
	//val  int
	next [10]*Num_trie //point to zero and one
}

func LongestCommonPrefix(arr1 []int, arr2 []int) int {
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	var t1 *Num_trie = new(Num_trie)
	var t2 *Num_trie = new(Num_trie)
	for i := 0; i < l1; i++ {
		var s string = strconv.Itoa(arr1[i])
		var cur_len int = len(s)
		var visit *Num_trie = t1
		for j := 0; j < cur_len; j++ {
			if visit.next[s[j]-'0'] == nil {
				visit.next[s[j]-'0'] = new(Num_trie)
			}
			visit = visit.next[s[j]-'0']
		}
	}
	for i := 0; i < l2; i++ {
		var s string = strconv.Itoa(arr2[i])
		var cur_len int = len(s)
		var visit *Num_trie = t2
		for j := 0; j < cur_len; j++ {
			if visit.next[s[j]-'0'] == nil {
				visit.next[s[j]-'0'] = new(Num_trie)
			}
			visit = visit.next[s[j]-'0']
		}
	}
	//var max_len int = 0
	type Pair struct {
		visit1 *Num_trie
		visit2 *Num_trie
	}
	var steps int = 0
	var q list.List
	var first Pair
	first.visit1 = t1
	first.visit2 = t2
	q.PushBack(first)
	for q.Len() > 0 {
		var l int = q.Len()
		for i := 0; i < l; i++ {
			cur := q.Front().Value.(Pair)
			q.Remove(q.Front())
			for j := 0; j < 10; j++ {
				if cur.visit1.next[j] != nil && cur.visit2.next[j] != nil {
					var next Pair
					next.visit1 = cur.visit1.next[j]
					next.visit2 = cur.visit2.next[j]
					q.PushBack(next)
				}
			}
		}
		if q.Len() == 0 {
			break
		}
		steps++
	}
	return steps
}
