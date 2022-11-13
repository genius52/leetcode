package diagram

func find_bob_path(graph [][]int,cur int,bob int,bob_path []int,visited []bool)[]int{
	if cur == bob{
		return bob_path
	}
	if visited[cur]{
		return []int{}
	}
	visited[cur] = true
	for _,next := range graph[cur]{
		//var path2 []int = make([]int,len(bob_path))
		//copy(path2,bob_path)
		bob_path = append(bob_path,next)
		ret := find_bob_path(graph,next,bob,bob_path,visited)
		if len(ret) > 0{
			return ret
		}else{
			bob_path = bob_path[:len(bob_path) - 1]
		}
	}
	return []int{}
}

func dfs_alice_path(graph [][]int,cur int,steps int,node_steps map[int]int,amount []int,visited []bool)int{
	var res int = 0
	if _,ok := node_steps[cur];ok{
		if steps < node_steps[cur]{ //alice 比 bob 更早到达
			res += amount[cur]
		}else if steps == node_steps[cur]{ //同时到达
			res += amount[cur]/2
		}
	}else{
		res += amount[cur]
	}
	visited[cur] = true
	var ret int = -(1 << 31 - 1)
	//var id_leaf bool = true
	for _,next := range graph[cur]{
		if visited[next]{
			continue
		}
		//id_leaf = false
		ret = max_int(ret,dfs_alice_path(graph,next,steps + 1,node_steps,amount,visited))
	}
	if ret == -(1 << 31 - 1){
		return res
	}
	return res + ret
}

func MostProfitablePath(edges [][]int, bob int, amount []int) int {
	var l int = len(amount)
	var graph [][]int = make([][]int,l)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var visited []bool = make([]bool,l)
	var bob_path []int = find_bob_path(graph,0,bob,[]int{0},visited)
	var node_steps map[int]int = make(map[int]int)
	var l2 int = len(bob_path)
	for i := 0;i < l2;i++{
		node_steps[bob_path[i]] = l2 - 1 - i
	}
	var visited2 []bool = make([]bool,l)
	return dfs_alice_path(graph,0,0,node_steps,amount,visited2)
}