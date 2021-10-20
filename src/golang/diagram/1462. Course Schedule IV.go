package diagram

//Input: n = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
//Output: [true,true]

//Input: n = 3, prerequisites = [[1,0],[2,0]], queries = [[0,1],[2,0]]
//Output: [false,true]
//请判断 queries[i][0] 是否是 queries[i][1] 的先修课程
//Floyd solution
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool{
	var graph [][]bool = make([][]bool,numCourses)
	for i := 0;i < numCourses;i++{
		graph[i] = make([]bool,numCourses)
	}
	for _,pre := range prerequisites{
		graph[pre[0]][pre[1]] = true
	}
	for k := 0;k < numCourses;k++{
		for i := 0;i < numCourses;i++{
			for j := 0;j < numCourses;j++{
				graph[i][j] = graph[i][j] || (graph[i][k] && graph[k][j])
			}
		}
	}
	var q_len int = len(queries)
	var res []bool = make([]bool,q_len)
	for i := 0;i < q_len;i++{
		res[i] = graph[queries[i][0]][queries[i][1]]
	}
	return res
}

//Optimized dfs solution
func dfs_checkIfPrerequisite2(graph *[][]int,source int,target int,memo *[][]int)bool{
	if source == target{
		return true
	}
	if (*memo)[source][target] != 0{
		return (*memo)[source][target] == 1
	}
	for i := 0;i < len((*graph)[source]);i++{
		ret := dfs_checkIfPrerequisite2(graph,(*graph)[source][i],target,memo)
		if ret{
			(*memo)[source][target] = 1
			return true
		}else{
			(*memo)[source][target] = 2
		}
	}
	return false
}

func CheckIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	var graph [][]int = make([][]int,numCourses)
	for _,pre := range prerequisites{
		graph[pre[0]] = append(graph[pre[0]],pre[1])
	}
	var q_len int = len(queries)
	var res []bool = make([]bool,q_len)
	var memo [][]int = make([][]int,numCourses)
	for i := 0;i < numCourses;i++{
		memo[i] = make([]int,numCourses)
	}
	for i := 0;i < q_len;i++{
		res[i] = dfs_checkIfPrerequisite2(&graph,queries[i][0],queries[i][1],&memo)
	}
	return res
}

//func dfs_checkIfPrerequisite(cur int,target int,graph map[int]map[int]bool,memo [][]int)bool{
//	if cur == target{
//		return true
//	}
//	if memo[cur][target] > 0{
//		if memo[cur][target] == 2{
//			return true
//		}else {
//			return false
//		}
//	}
//	var res bool = false
//	if sub,ok := graph[cur];ok{
//		for pre,_ := range sub{
//			res = dfs_checkIfPrerequisite(pre,target,graph,memo)
//			if res{
//				break
//			}
//		}
//	}else{
//		res = false
//	}
//	if res{
//		memo[cur][target] = 2
//	}else{
//		memo[cur][target] = 1
//	}
//	return res
//}
//
//func CheckIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
//	var graph map[int]map[int]bool = make(map[int]map[int]bool)
//	for _,q := range prerequisites{
//		if _,ok := graph[q[1]];!ok{
//			graph[q[1]] = make(map[int]bool)
//		}
//		graph[q[1]][q[0]] = true
//	}
//	var memo [][]int = make([][]int,n)
//	for i := 0;i < n;i++{
//		memo[i] = make([]int,n)
//	}
//	var l int = len(queries)
//	var res []bool = make([]bool,l)
//	var i int = 0
//	for _,q := range queries{
//		ret := dfs_checkIfPrerequisite(q[1],q[0],graph,memo)
//		res[i] = ret
//		i++
//	}
//	return res
//}