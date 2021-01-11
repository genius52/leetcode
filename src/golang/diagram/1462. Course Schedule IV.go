package diagram

//Input: n = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
//Output: [true,true]

//Input: n = 3, prerequisites = [[1,0],[2,0]], queries = [[0,1],[2,0]]
//Output: [false,true]
func dfs_checkIfPrerequisite(cur int,target int,graph map[int]map[int]bool,memo [][]int)bool{
	if cur == target{
		return true
	}
	if memo[cur][target] > 0{
		if memo[cur][target] == 2{
			return true
		}else {
			return false
		}
	}
	var res bool = false
	if sub,ok := graph[cur];ok{
		for pre,_ := range sub{
			res = dfs_checkIfPrerequisite(pre,target,graph,memo)
			if res{
				break
			}
		}
	}else{
		res = false
	}
	if res{
		memo[cur][target] = 2
	}else{
		memo[cur][target] = 1
	}
	return res
}

func CheckIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	for _,q := range prerequisites{
		if _,ok := graph[q[1]];!ok{
			graph[q[1]] = make(map[int]bool)
		}
		graph[q[1]][q[0]] = true
	}
	var memo [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		memo[i] = make([]int,n)
	}
	var l int = len(queries)
	var res []bool = make([]bool,l)
	var i int = 0
	for _,q := range queries{
		ret := dfs_checkIfPrerequisite(q[1],q[0],graph,memo)
		res[i] = ret
		i++
	}
	return res
}