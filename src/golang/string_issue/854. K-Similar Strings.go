package string_issue

import (
	"container/list"
	"strings"
)

func equal(data1 []string,data2 []string,l int)bool{
	for i := 0;i < l;i++{
		if data1[i] != data2[i]{
			return false
		}
	}
	return true
}

func KSimilarity(s1 string, s2 string) int {
	var l int = len(s1)
	var target []string = strings.Split(s2,"")
	var visited map[string]bool = make(map[string]bool)
	var steps int = 0
	var q list.List
	q.PushBack(s1)
	visited[s1] = true
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur string = q.Front().Value.(string)
			q.Remove(q.Front())
			if cur == s2{
				return steps
			}
			var data []string = strings.Split(cur,"")
			for j := 0;j < l;j++{
				if data[j] == target[j]{
					continue
				}
				for k := j + 1;k < l;k++{
					if data[k] == target[k]{
						continue
					}
					if data[k] == target[j]{
						data[j],data[k] = data[k],data[j]
						next := strings.Join(data,"")
						data[j],data[k] = data[k],data[j]
						if _,ok := visited[next];ok{
							continue
						}
						visited[next] = true
						q.PushBack(next)
					}
				}
				break;
			}
		}
		steps++
	}
	return -1
}