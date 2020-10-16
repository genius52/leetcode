package diagram

import "container/list"

//Input: [[1,3], [0,2], [1,3], [0,2]]
//Output: true
//Explanation:
//The graph looks like this:
//0----1
//|    |
//|    |
//3----2
//We can divide the vertices into two groups: {0, 2} and {1, 3}.
//func IsBipartite(graph [][]int) bool {
//	var l int = len(graph)
//	var relation [][]bool = make([][]bool,l)
//	for i := 0;i < l;i++{
//		relation[i] = make([]bool,l)
//		for _,j := range graph[i]{
//			relation[i][j] = true
//		}
//	}
//	var first_group map[int]bool = make(map[int]bool)
//	first_group[0] = true
//	for i := 1;i < l;i++{
//		var conflict bool = false
//		for n,_ := range first_group{
//			if relation[i][n]{
//				conflict = true
//				break
//			}
//		}
//		if conflict{
//			continue
//		}
//		first_group[i] = true
//	}
//	var second_group []int
//	for i := 0;i < l;i++{
//		if _,ok := first_group[i];!ok{
//			second_group = append(second_group,i)
//		}
//	}
//	var second_conflict bool = false
//	for i := 0;i < len(second_group);i++{
//		for j := i + 1;j < len(second_group);j++{
//			if relation[second_group[i]][second_group[j]]{
//				second_conflict = true
//				break
//			}
//		}
//	}
//	if second_conflict{
//		return false
//	}
//	return true
//}

func IsBipartite(graph [][]int) bool {
	var l int = len(graph)
	var relation [][]bool = make([][]bool,l)
	for i := 0;i < l;i++{
		relation[i] = make([]bool,l)
		for _,j := range graph[i]{
			relation[i][j] = true
		}
	}
	var visited []int = make([]int,l)//-1 : not viisted, 1: 1st group, 2: 2nd group
	for i := 0;i < l;i++{
		visited[i] = -1
	}

	for i := 0;i < l;i++{
		if visited[i] != -1{
			continue
		}
		var q list.List
		q.PushBack(i)
		for q.Len() > 0{
			var n int = q.Front().Value.(int)
			q.Remove(q.Front())
			for i := 0;i < l;i++{
				if relation[n][i]{
					if visited[i] != -1{
						if visited[n] == visited[i]{
							return false
						}
					}else{
						if visited[n] == 1{
							visited[i] = 2
						}else{
							visited[i] = 1
						}
						q.PushBack(i)
					}
				}
			}
		}
	}

	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if relation[i][j]{
				if visited[i] == visited[j]{
					return false
				}
			}
		}
	}
	return true
}