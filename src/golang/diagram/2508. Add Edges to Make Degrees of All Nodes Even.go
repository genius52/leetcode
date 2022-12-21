package diagram

func IsPossible(n int, edges [][]int) bool {
	var degree []int = make([]int,n + 1)
	var graph []map[int]bool = make([]map[int]bool,n + 1)
	for i := 1;i <= n;i++{
		graph[i] = make(map[int]bool)
	}
	for _,edge := range edges{
		graph[edge[0]][edge[1]] = true
		graph[edge[1]][edge[0]] = true
		degree[edge[0]]++
		degree[edge[1]]++
	}
	var odd_cnt int = 0
	var odd_node []int
	for i := 1;i <= n;i++{
		if (degree[i] | 1) == degree[i]{
			if degree[i] == n - 1{
				return false
			}
			odd_node = append(odd_node,i)
			odd_cnt++
			if odd_cnt > 4{
				return false
			}
		}
	}
	if odd_cnt == 0{
		return true
	}
	var even_cnt int = n - odd_cnt
	if odd_cnt == 2{
		if _,ok := graph[odd_node[0]][odd_node[1]];!ok{
			return true
		}
		if even_cnt == 0{
			return false
		}
		for i := 1;i <= n;i++{
			if odd_node[0] == i || odd_node[1] == i{
				continue
			}
			if _,ok1 := graph[odd_node[0]][i];!ok1{
				if _,ok2 := graph[odd_node[1]][i];!ok2{
					return true
				}
			}
		}
	}else if odd_cnt == 4{
		a := odd_node[0]
		b := odd_node[1]
		c := odd_node[2]
		d := odd_node[3]
		if _,ok1 := graph[a][b];!ok1{
			if _,ok2 := graph[c][d];!ok2{
				return true
			}
		}
		if _,ok1 := graph[a][c];!ok1{
			if _,ok2 := graph[b][d];!ok2{
				return true
			}
		}
		if _,ok1 := graph[a][d];!ok1{
			if _,ok2 := graph[b][c];!ok2{
				return true
			}
		}
	}
	return false
}