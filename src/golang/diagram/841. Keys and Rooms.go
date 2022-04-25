package diagram

import "container/list"

// [[1],[2],[3],[]]
//  [[1,3,2],[2,3],[2,1,3,1],[]]
//  [[1,3],[3,0,1],[2],[0]]
//bfs solution
func CanVisitAllRooms(rooms [][]int) bool {
	var l int = len(rooms)
	var visited map[int]bool = make(map[int]bool)
	var q list.List
	q.PushBack(0)
	visited[0] = true
	for q.Len() != 0{
		var next_len int = q.Len()
		for i := 0;i < next_len;i++{
			next := q.Front().Value.(int)
			q.Remove(q.Front())
			for i := 0;i < len(rooms[next]);i++{
				if _,ok := visited[rooms[next][i]];!ok{
					q.PushBack(rooms[next][i])
					visited[rooms[next][i]] = true
				}
			}
		}
	}
	return len(visited) == l
}

//dfs solution
func dfs_canVisitAllRooms(rooms [][]int,cur int,visited []bool){
	if visited[cur]{
		return
	}
	visited[cur] = true
	var next_len int = len(rooms[cur])
	for i := 0;i < next_len;i++{
		dfs_canVisitAllRooms(rooms,rooms[cur][i],visited)
	}
}

func canVisitAllRooms2(rooms [][]int) bool {
	var l int = len(rooms)
	var visited []bool = make([]bool,l)
	dfs_canVisitAllRooms(rooms,0,visited)
	for i := 0;i < l;i++{
		if !visited[i]{
			return false
		}
	}
	return true
}