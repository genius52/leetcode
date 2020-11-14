package diagram

import "math"

//Input:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
//Output: 200
func dfs_findCheapestPrice(n int,graph [][]int,src int,dst int,K int,memo map[string]int,visited []bool)int{
	if K < 0{
		return math.MaxInt32
	}
	k := string(src) + "," + string(K)
	if val,ok := memo[k];ok{
		return val
	}
	if src == dst{
		return 0
	}
	visited[src] = true
	var min_cost int = math.MaxInt32
	for i := 0;i < n;i++{
		if graph[src][i] == 0{
			continue
		}
		if visited[i]{
			continue
		}
		costs := dfs_findCheapestPrice(n,graph,i,dst,K - 1,memo,visited)
		if costs != math.MaxInt32{
			if graph[src][i] + costs < min_cost{
				min_cost = graph[src][i] + costs
			}
		}
	}
	visited[src] = false
	if min_cost != math.MaxInt32{
		memo[k] = min_cost
	}
	return min_cost
}

func FindCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	var graph [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		graph[i] = make([]int,n)
	}
	for _,f := range flights{
		graph[f[0]][f[1]] = f[2]
	}
	var memo map[string]int = make(map[string]int)
	var visited []bool = make([]bool,n)
	res := dfs_findCheapestPrice(n,graph,src,dst,K + 1,memo,visited)
	if res == math.MaxInt32{
		return -1
	}
	return res
}