package diagram

import (
	"container/list"
	"sort"
)

//输入：watchedVideos = [["A","B"],["C"],["B","C"],["D"]], friends = [[1,2],[0,3],[0,3],[1,2]], id = 0, level = 1
//输出：["B","C"]
type Freq_name struct {
	freq int
	name string
}

func WatchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	var n int = len(friends)
	var graph [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		for j := 0;j < len(friends[i]);j++{
			graph[i] = append(graph[i],friends[i][j])
		}
	}
	var visited []bool = make([]bool,n)
	var q list.List
	q.PushBack(id)
	visited[id] = true
	for level > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			for j := 0;j < len(graph[cur]);j++{
				if graph[cur][j] == cur  || visited[graph[cur][j]]{
					continue
				}
				q.PushBack(graph[cur][j])
				visited[graph[cur][j]] = true
			}
		}
		level--
	}
	var record map[string]int = make(map[string]int)
	for q.Len() > 0{
		cur := q.Front().Value.(int)
		q.Remove(q.Front())
		for i := 0;i < len(watchedVideos[cur]);i++{
			record[watchedVideos[cur][i]]++
		}
	}
	tmp := []Freq_name{}
	for key,val := range record{
		tmp = append(tmp,Freq_name{
			freq:val,
			name:key,
		})
	}
	sort.Slice(tmp,func(i int,j int)bool{
		if tmp[i].freq == tmp[j].freq{
			return tmp[i].name < tmp[j].name
		}
		return tmp[i].freq < tmp[j].freq
	})
	var res []string
	for i := 0;i < len(tmp);i++{
		res = append(res,tmp[i].name)
	}
	return res
}