package diagram

//Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
//Output: [2,4,5,6]
const (
	cycle int = 1
	nocycle int  = 2
)

func hascycle_eventualSafeNodes(record [][]int,cur int,visited []int)bool{
	if visited[cur] == cycle{
		return true
	}
	if visited[cur] == nocycle{
		return false
	}
	visited[cur] = cycle
	var ret bool = false
	for i := 0;i < len(record[cur]);i++{
		next := record[cur][i]
		ret = hascycle_eventualSafeNodes(record,next,visited)
		if ret{
			break
		}
	}
	if !ret{
		visited[cur] = nocycle
	}
	return ret
}


func EventualSafeNodes(graph [][]int) []int {
	var l int = len(graph)
	var res []int
	var visited []int = make([]int,l)
	for i := 0;i < l;i++{
		if visited[i] == cycle{
			continue
		}else if visited[i] == nocycle{
			res = append(res,i)
			continue
		}
		hascycle := hascycle_eventualSafeNodes(graph,i,visited)
		if !hascycle{
			res = append(res,i)
			//visited[i] = nocycle
		}
	}
	return res
}