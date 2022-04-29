package diagram

import "container/list"

//Input: blocked = [], source = [0,0], target = [999999,999999]
//Output: true

func check_isEscapePossible(forbid map[int64]bool, source []int, target []int)bool{
	var visited map[int64]bool = make(map[int64]bool)
	var q list.List
	var start Point
	start.x = source[0]
	start.y = source[1]
	q.PushBack(start)
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur Point = q.Front().Value.(Point)
			if cur.x == target[0] && cur.y == target[1]{
				return true
			}
			q.Remove(q.Front())
			for _,dir := range dirs{
				var next Point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.x >= 1000000 || next.y < 0 || next.y >= 1000000{
					continue
				}
				key := int64(next.x * 1000000 + next.y)
				if _,ok := visited[key];ok{
					continue
				}
				if _,ok := forbid[key];ok{
					continue
				}
				visited[int64(next.x * 1000000 + next.y)] = true
				if len(visited) > 20000{
					return true
				}
				q.PushBack(next)
			}
		}
	}
	return false
}

func IsEscapePossible(blocked [][]int, source []int, target []int) bool {
	var forbid map[int64]bool = make(map[int64]bool)
	for _,block := range blocked{
		forbid[int64(block[0] * 1000000 + block[1])] = true
	}
	if len(forbid) == 0{
		return true
	}
	return check_isEscapePossible(forbid,source,target) && check_isEscapePossible(forbid,target,source)
}