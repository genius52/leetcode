package diagram

import "container/list"

//func dfs_minimumCost(start []int, target []int, specialRoads [][]int, visited map[int64]int) int {
//	if start[0] == target[0] && start[1] == target[1] {
//		return 0
//	}
//	var key int64 = int64(start[0])*100000 + int64(start[1])
//	if _, ok := visited[key]; ok {
//		return visited[key]
//	}
//	visited[key] = 1 << 30
//	var res int = abs_int(start[0]-target[0]) + abs_int(start[1]-target[1])
//	for _, road := range specialRoads {
//		start_x := road[0]
//		start_y := road[1]
//		end_x := road[2]
//		end_y := road[3]
//		cost := road[4]
//		ret := dfs_minimumCost([]int{end_x, end_y}, target, specialRoads, visited)
//		if ret != -1 {
//			cur := ret + abs_int(start[0]-start_x) + abs_int(start[1]-start_y) + cost
//			if cur < res {
//				res = cur
//			}
//		}
//	}
//	visited[key] = res
//	return res
//}
//
//func MinimumCost(start []int, target []int, specialRoads [][]int) int {
//	var visited map[int64]int = make(map[int64]int)
//	return dfs_minimumCost(start, target, specialRoads, visited)
//}

type pos_cost struct {
	x    int
	y    int
	cost int
}

func MinimumCost(start []int, target []int, specialRoads [][]int) int {
	//var visited [100001][100001]bool
	var l int = len(specialRoads)
	var from_start []int = make([]int, l)
	for i := 0; i < l; i++ {
		from_start[i] = abs_int(start[0]-specialRoads[i][0]) + abs_int(start[1]-specialRoads[i][1])
	}
	var q list.List
	var p pos_cost
	p.x = start[0]
	p.y = start[1]
	var res int = abs_int(start[0]-target[0]) + abs_int(start[1]-target[1])
	q.PushBack(p)
	for q.Len() > 0 {

	}
	return res
}
