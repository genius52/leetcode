package diagram

import "container/list"

//There are n cities connected by some number of flights.
//You are given an array flights where flights[i] = [fromi, toi, pricei] indicates
//that there is a flight from city fromi to city toi with cost pricei.
//
//You are also given three integers src, dst, and k,
//return the cheapest price from src to dst with at most k stops. If there is no such route, return -1.


//Input:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
//Output: 200
type city_cost struct{
	city int
	money int
}

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var graph [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		graph[i] = make([]int,n)
	}
	for _,f := range flights{
		graph[f[0]][f[1]] = f[2]
	}
	var visited map[int]int = make(map[int]int)//visited[n] : cost of from src to n
	var q list.List
	var start city_cost
	start.city = src
	q.PushBack(start)
	var res int = 2147483647
	for q.Len() > 0 && k >= -1{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur city_cost = q.Front().Value.(city_cost)
			visited[cur.city] = cur.money
			q.Remove(q.Front())
			if cur.city == dst{
				res = min_int_number(res,cur.money)
				continue
			}
			for j := 0;j < n;j++{
				if graph[cur.city][j] == 0{
					continue
				}
				var next city_cost
				next.city = j
				next.money = cur.money + graph[cur.city][j]
				if c,ok := visited[j];ok{
					if next.money < c{
						visited[j] = next.money
						q.PushBack(next)
					}
				}else{
					q.PushBack(next)
				}
			}
		}
		k--
	}
	if res == 2147483647{
		return -1
	}
	return res
}


//func dfs_findCheapestPrice(n int,graph [][]int,src int,dst int,K int,memo map[string]int,visited []bool)int{
//	if K < 0{
//		return math.MaxInt32
//	}
//	k := string(src) + "," + string(K)
//	if val,ok := memo[k];ok{
//		return val
//	}
//	if src == dst{
//		return 0
//	}
//	visited[src] = true
//	var min_cost int = math.MaxInt32
//	for i := 0;i < n;i++{
//		if graph[src][i] == 0{
//			continue
//		}
//		if visited[i]{
//			continue
//		}
//		costs := dfs_findCheapestPrice(n,graph,i,dst,K - 1,memo,visited)
//		if costs != math.MaxInt32{
//			if graph[src][i] + costs < min_cost{
//				min_cost = graph[src][i] + costs
//			}
//		}
//	}
//	visited[src] = false
//	if min_cost != math.MaxInt32{
//		memo[k] = min_cost
//	}
//	return min_cost
//}
//
//func FindCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
//	var graph [][]int = make([][]int,n)
//	for i := 0;i < n;i++{
//		graph[i] = make([]int,n)
//	}
//	for _,f := range flights{
//		graph[f[0]][f[1]] = f[2]
//	}
//	var memo map[string]int = make(map[string]int)
//	var visited []bool = make([]bool,n)
//	res := dfs_findCheapestPrice(n,graph,src,dst,K + 1,memo,visited)
//	if res == math.MaxInt32{
//		return -1
//	}
//	return res
//}