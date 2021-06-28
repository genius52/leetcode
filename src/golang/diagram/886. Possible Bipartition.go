package diagram

//Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
//Output: true
func dfs_possibleBipartition(cur_node int,val int,relation map[int]map[int]bool,group []int)bool{
	if group[cur_node] == -val{
		return false
	}
	if group[cur_node] == val{
		return true
	}
	group[cur_node] = val
	for dislike,_ := range relation[cur_node]{
		if !dfs_possibleBipartition(dislike,-val,relation,group){
			return false
		}
	}
	return true
}

func PossibleBipartition(N int, dislikes [][]int) bool{
	var relation map[int]map[int]bool = make(map[int]map[int]bool)
	for _,r := range dislikes{
		if _,ok := relation[r[0]];!ok{
			relation[r[0]] = make(map[int]bool)
		}
		relation[r[0]][r[1]] = true
		if _,ok := relation[r[1]];!ok{
			relation[r[1]] = make(map[int]bool)
		}
		relation[r[1]][r[0]] = true
	}
	var group []int = make([]int,N + 1)
	for i := 1;i <= N;i++{
		if group[i] != 0{
			continue
		}
		group[i] = 1
		for dislike,_ := range relation[i]{
			if !dfs_possibleBipartition(dislike,-1,relation,group) {
				return false
			}
		}
	}
	return true
}

//func dfs_PossibleBipartition(last int,cur int,N int,relation [][]int,group []int)bool{
//	if group[cur] == group[last] {
//		return false
//	}
//	if group[last] == 1{
//		group[cur] = 2
//	}else if group[last] == 2{
//		group[cur] = 1
//	}
//	for j := 1;j <= N;j++{
//		if j == cur || j == last ||relation[cur][j] == 0{
//			continue
//		}
//		if group[j] == 0{
//			res := dfs_PossibleBipartition(cur,j,N,relation,group)
//			if !res{
//				return false
//			}
//		}else{
//			if group[cur] == group[j]{
//				return false
//			}
//		}
//	}
//	return true
//}
//
//func PossibleBipartition(N int, dislikes [][]int) bool {
//	var relation [][]int = make([][]int,N + 1)
//	for i := 0;i <= N;i++{
//		relation[i] = make([]int,N + 1)
//	}
//	var l int = len(dislikes)
//	for i := 0;i < l;i++{
//		relation[dislikes[i][0]][dislikes[i][1]] = 1
//		relation[dislikes[i][1]][dislikes[i][0]] = 1
//	}
//	var group []int = make([]int,N + 1)
//	for i := 1;i <= N;i++{
//		var cnt int = 0
//		for j := 1;j <= N;j++{
//			if relation[i][j] == 0{
//				continue
//			}
//			cnt++
//			if group[i] != 0{
//				res := dfs_PossibleBipartition(i,j,N,relation,group)
//				if !res{
//					return false
//				}
//			}else{
//				group[i] = 1
//				res1 := dfs_PossibleBipartition(i,j,N,relation,group)
//				if !res1{
//					group[i] = 2
//					res2 := dfs_PossibleBipartition(i,j,N,relation,group)
//					if !res1 && !res2{
//						return false
//					}
//				}
//			}
//		}
//		if cnt == N{
//			return false
//		}
//	}
//	return true
//}