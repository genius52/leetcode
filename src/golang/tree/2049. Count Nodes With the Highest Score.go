package tree

func dfs_countHighestScoreNodes(graph [][]int,idx int,total int,record map[int]int,max_val *int)int{
	var l int = len(graph[idx])
	if l == 0{
		record[total - 1]++
		*max_val = max_int(*max_val,total - 1)
		return 1
	}
	if l == 1{
		children := dfs_countHighestScoreNodes(graph,graph[idx][0],total,record,max_val)
		parents := total - children - 1
		scores := parents * children
		if scores != 0{
			record[scores]++
			*max_val = max_int(*max_val,scores)
		}else{
			record[children]++
			*max_val = max_int(*max_val,children)
		}
		return 1 + children
	}else{
		left_children := dfs_countHighestScoreNodes(graph,graph[idx][0],total,record,max_val)
		right_children := dfs_countHighestScoreNodes(graph,graph[idx][1],total,record,max_val)
		parents := total - left_children - right_children - 1
		scores := parents * left_children * right_children
		if scores != 0{
			record[scores]++
			*max_val = max_int(*max_val,scores)
		}else{
			record[left_children * right_children]++
			*max_val = max_int(*max_val,left_children * right_children)
		}
		return 1 + left_children + right_children
	}
}

func CountHighestScoreNodes(parents []int) int{
	var l int = len(parents)
	var graph [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		if parents[i] == -1{
			continue
		}
		graph[parents[i]] = append(graph[parents[i]],i)
	}
	var total int = len(graph)
	var record map[int]int = make(map[int]int)
	var max_val int = 0
	dfs_countHighestScoreNodes(graph,0,total,record,&max_val)
	return record[max_val]
}

//func post_countHighestScoreNodes(graph [][]int,cur int,node_children *[][2]int)int{
//	var l int = len(graph[cur])
//	var total int = 0
//	if l == 0{ //leaf
//		total = 1
//	}else if l == 1{
//		children := post_countHighestScoreNodes(graph,graph[cur][0],node_children)
//		(*node_children)[cur][0] = children
//		total = children + 1
//	}else if l == 2{
//		left_children := post_countHighestScoreNodes(graph,graph[cur][0],node_children)
//		right_children := post_countHighestScoreNodes(graph,graph[cur][1],node_children)
//		(*node_children)[cur][0] = left_children
//		(*node_children)[cur][1] = right_children
//		total = left_children + right_children + 1
//	}
//	return total
//}
//
//func CountHighestScoreNodes(parents []int) int {
//	var l int = len(parents)
//	var graph [][]int = make([][]int,l)
//	for i := 0;i < l;i++{
//		if parents[i] == -1{
//			continue
//		}
//		graph[parents[i]] = append(graph[parents[i]],i)
//	}
//	var node_children [][2]int = make([][2]int,l)
//	for i := 0;i < l;i++{
//		if node_children[i][0] > 0 || node_children[i][1] > 0{
//			continue
//		}
//		post_countHighestScoreNodes(graph,i,&node_children)
//
//	}
//	var record map[int]int = make(map[int]int)
//	var max_val int = 0
//	for i := 0;i < l;i++{
//		left_children := node_children[i][0]
//		right_children := node_children[i][1]
//		rest := l - left_children - right_children - 1
//		var product int = 0
//		if left_children == 0 && right_children == 0{
//			product = rest
//		}else if left_children == 0 && rest == 0{
//			product = right_children
//		}else if right_children == 0 && rest == 0{
//			product = left_children
//		}else if left_children == 0{
//			product = right_children * rest
//		}else if right_children == 0{
//			product = left_children * rest
//		}else if rest == 0{
//			product = left_children * right_children
//		}else{
//			product = left_children * right_children * rest
//		}
//		if product > max_val{
//			max_val = product
//		}
//		record[product]++
//	}
//	return record[max_val]
//}