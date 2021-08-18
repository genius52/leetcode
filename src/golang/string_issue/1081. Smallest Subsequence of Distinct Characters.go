package string_issue

import "container/list"

//Input: s = "cbacdcbc"
//Output: "acdb"

func SmallestSubsequence(s string) string{
	var char_cnt [26]int
	for _,c := range s{
		char_cnt[c - 'a']++
	}
	var visited [26]bool
	var q list.List
	var l int = len(s)
	for i := 0;i < l;i++{
		for q.Len() > 0 &&  s[i] < q.Back().Value.(uint8) && !visited[s[i] - 'a'] && char_cnt[q.Back().Value.(uint8) - 'a'] > 0{
			visited[q.Back().Value.(uint8) - 'a'] = false
			q.Remove(q.Back())
		}
		if !visited[s[i] - 'a']{
			q.PushBack(s[i])
		}
		char_cnt[s[i] - 'a']--
		visited[s[i] - 'a'] = true
	}
	var res string
	for q.Len() > 0{
		res += string(q.Front().Value.(uint8))
		q.Remove(q.Front())
	}
	return res
}

//func smallestSubsequence(s string) string {
//	var q list.List
//	var l int = len(s)
//	for i := 0;i < l;i++{
//		for q.Len() > 0 &&  s[i] < q.Back().Value.(uint8){
//			q.Remove(q.Back())
//		}
//		q.PushBack(s[i])
//	}
//	var res string
//	for q.Len() > 0{
//		res += string(q.Front().Value.(uint8))
//		q.Remove(q.Front())
//	}
//	return res
//}