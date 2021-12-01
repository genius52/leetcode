package diagram

//Input: source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
//Output: 1
//Explanation: source can be transformed the following way:
//- Swap indices 0 and 1: source = [2,1,3,4]
//- Swap indices 2 and 3: source = [2,1,4,3]
//The Hamming distance of source and target is 1 as they differ in 1 position: index 3.
func MinimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	type funcType func (groups []int,node int)int
	var get_parent funcType
	get_parent = func (groups []int,node int)int{
		if groups[node] != node{
			groups[node] = get_parent(groups,groups[node])
		}
		return groups[node]
	}
	var l int = len(source)
	var groups []int = make([]int,l)
	for i := 0;i < l;i++{
		groups[i] = i
	}
	for i := 0;i < len(allowedSwaps);i++{
		group1 := get_parent(groups,allowedSwaps[i][0])
		group2 := get_parent(groups,allowedSwaps[i][1])
		if group1 != group2{
			if group1 < group2{
				groups[group2] = group1
			}else{
				groups[group1] = group2
			}
		}
	}
	var record map[int]map[int]int = make(map[int]map[int]int)//record[i] : group[i]'s number and count

	for i := 0;i < l;i++{
		group := get_parent(groups,i)
		if record[group] == nil{
			record[group] = make(map[int]int)
		}
		record[group][source[i]]++
	}
	var same_cnt int = 0
	for i := 0;i < l;i++{
		group := get_parent(groups,i)
		if record[group][target[i]] > 0{
			record[group][target[i]]--
			same_cnt++
		}
	}
	return l - same_cnt
}

//func dfs_minimumHammingDistance(index int,group *[]int,graph map[int][]int,visited []bool){
//	if visited[index]{
//		return
//	}
//	*group = append(*group,index)
//	visited[index] = true
//	if _,ok := graph[index];!ok{
//		return
//	}
//	for _,next := range graph[index]{
//		dfs_minimumHammingDistance(next,group,graph,visited)
//	}
//}
//
//func MinimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
//	var graph map[int][]int = make(map[int][]int)
//	for _,item := range allowedSwaps{
//		if _,ok := graph[item[0]];ok{
//			graph[item[0]] = append(graph[item[0]],item[1])
//		}else{
//			graph[item[0]] = []int{item[1]}
//		}
//		if _,ok := graph[item[1]];ok{
//			graph[item[1]] = append(graph[item[1]],item[0])
//		}else{
//			graph[item[1]] = []int{item[0]}
//		}
//	}
//	var l int = len(source)
//	var visited []bool = make([]bool,l)
//	var res int = 0
//	for i := 0;i < l;i++{
//		if visited[i]{
//			continue
//		}
//		if _,ok := graph[i];!ok{
//			if source[i] != target[i]{
//				res++
//			}
//		}else{
//			var group []int
//			dfs_minimumHammingDistance(i,&group,graph,visited)
//			var source_map map[int]int = make(map[int]int)
//			var target_map map[int]int = make(map[int]int)
//			for _,pos := range group{
//				if _,ok := source_map[source[pos]];ok{
//					source_map[source[pos]]++
//				}else{
//					source_map[source[pos]] = 1
//				}
//				if _,ok := target_map[target[pos]];ok{
//					target_map[target[pos]]++
//				}else{
//					target_map[target[pos]] = 1
//				}
//			}
//			for n,cnt := range source_map{
//				if cnt2,ok := target_map[n];ok{
//					if cnt > cnt2{
//						res += (cnt - cnt2)
//					}
//				}else{
//					res += cnt
//				}
//			}
//		}
//	}
//	return res
//}