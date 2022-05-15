package diagram

import (
	"container/list"
	"strings"
)

func OpenLock(deadends []string, target string) int {
	var visited map[string]bool = make(map[string]bool)
	for _,s := range deadends{
		if s == "0000"{
			return -1
		}
		visited[s] = true
	}
	var q list.List
	q.PushBack("0000")
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(string)
			q.Remove(q.Front())
			if cur == target{
				return steps
			}
			for i := 0;i < 4;i++{
				var add strings.Builder
				var minus strings.Builder
				for j := 0;j < 4;j++{
					if i == j{
						add.WriteByte((cur[j] - '0' + 1) % 10 + '0')
						minus.WriteByte((cur[j] - '0' + 10 - 1) % 10 + '0')
					}else{
						add.WriteByte(cur[j])
						minus.WriteByte(cur[j])
					}
				}
				next1 := add.String()
				if  _,ok := visited[next1];!ok{
					visited[next1] = true
					q.PushBack(next1)
				}
				next2 := minus.String()
				if  _,ok := visited[next2];!ok{
					visited[next2] = true
					q.PushBack(next2)
				}
			}
		}
		steps++
	}
	return -1
}