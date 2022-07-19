package diagram

import (
	"container/list"
)

//Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
//Output: 2
//Explanation: The best strategy is take the first bus to the bus stop 7,
//then take the second bus to the bus stop 6.

//type Route_stop struct{
//	route int
//	bus_stop int
//}
//
//func NumBusesToDestination(routes [][]int, source int, target int) int {
//	var l int = len(routes)
//	var bus_stop_route map[int][]int = make(map[int][]int) //key:bus stop, val:route1,route2...
//	for i := 0;i < l;i++{
//		var l1 int = len(routes[i])
//		for j := 0;j < l1;j++{
//			bus_stop_route[routes[i][j]] = append(bus_stop_route[routes[i][j]],i)
//		}
//	}
//	//visited route or bus stop + route?
//	var visited_route []bool = make([]bool,l)
//	var visited_busstop map[int]bool = make(map[int]bool)
//	var q list.List
//	for i := 0;i < l;i++{
//		var l1 int = len(routes[i])
//		for j := 0;j < l1;j++{
//			if routes[i][j] == source{
//				var r Route_stop
//				r.bus_stop = routes[i][j]
//				r.route = i
//				q.PushBack(r)
//				key := routes[i][j] * 1000 + i
//				visited_busstop[key] = true
//				break
//			}
//		}
//	}
//	var steps int = 0
//	for q.Len() > 0{
//		var cur_len int = q.Len()
//		for i := 0;i < cur_len;i++{
//			var cur Route_stop = q.Front().Value.(Route_stop)
//			if cur.bus_stop == target{
//				return steps
//			}
//			q.Remove(q.Front())
//			if visited_route[cur.route]{
//				continue
//			}
//			//选中当前路线的其他站点
//			//找到这些站点能到达的其他路线
//			//将新的站点 + 路线 加入队列
//			for _,other_bus_stop := range routes[cur.route]{
//				if other_bus_stop == cur.bus_stop{
//					continue
//				}
//				for _,next_route := range bus_stop_route[other_bus_stop]{
//					key := other_bus_stop * 1000 + next_route
//					if _,ok2 := visited_busstop[key];ok2{
//						continue
//					}
//					visited_busstop[key] = true
//					var r Route_stop
//					r.bus_stop = other_bus_stop
//					r.route = next_route
//					q.PushBack(r)
//				}
//			}
//			visited_route[cur.route] = true
//		}
//		steps++
//	}
//	return -1
//}

func NumBusesToDestination(routes [][]int, source int, target int) int{
	var busstop_route map[int][]int = make(map[int][]int)//车站:所有相关线路
	var l int = len(routes)
	for i := 0;i < l;i++{
		for j := 0;j < len(routes[i]);j++{
			busstop_route[routes[i][j]] = append(busstop_route[routes[i][j]],i)
		}
	}
	var visited_route []bool = make([]bool,l) //访问过的线路
	var visited_busstop map[int]bool = make(map[int]bool) //访问过的站点
	var steps int = 0
	var q list.List
	q.PushBack(source)
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			if cur == target{
				return steps
			}
			for _,route := range busstop_route[cur]{
				if visited_route[route]{
					continue
				}
				visited_route[route] = true
				for _,next := range routes[route]{
					if visited_busstop[next]{
						continue
					}
					visited_busstop[next] = true
					q.PushBack(next)
				}
			}
		}
		steps++
	}
	return -1
}