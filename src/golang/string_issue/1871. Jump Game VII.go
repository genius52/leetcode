package string_issue

import "container/list"

//BFS
func CanReach(s string, minJump int, maxJump int) bool{
	var l int = len(s)
	if s[l - 1] == '1'{
		return false
	}
	var visited []bool = make([]bool,l)
	var q list.List
	q.PushBack(0)
	visited[0] = true
	var fastest int = 0
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur int = q.Front().Value.(int)
			if cur == l - 1{
				return true
			}
			q.Remove(q.Front())
			for next := max_int(cur + minJump,fastest + 1);next <= cur + maxJump;next++{
				if next >= l{
					continue
				}
				if visited[next]{
					continue
				}
				if s[next] == '1'{
					continue
				}
				visited[next] = true
				q.PushBack(next)
			}
			fastest = max_int(fastest,cur + maxJump)
		}
	}
	return false
}
//TLE by DP Solution
//func dfs_canReach(s string,pos int,minJump int, maxJump int,visited []bool)bool{
//	if pos == (len(s) - 1){
//		return true
//	}
//	if s[pos] == '1'{
//		return false
//	}
//	if visited[pos]{
//		return false
//	}
//	var res bool = false
//	for i := min_int(maxJump,len(s) - 1 - pos);i >= minJump;i--{
//		if s[pos + i] == '1'{
//			continue
//		}
//		if dfs_canReach(s,pos + i,minJump,maxJump,visited){
//			res = true
//			break
//		}
//	}
//	visited[pos] = true
//	return res
//}
//
//func CanReach(s string, minJump int, maxJump int) bool {
//	var l int = len(s)
//	if s[l - 1] == '1'{
//		return false
//	}
//	var visited []bool = make([]bool,l)
//	return dfs_canReach(s,0,minJump,maxJump,visited)
//}