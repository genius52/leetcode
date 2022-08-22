package tree

import "container/list"

func creategraph(node *TreeNode,parent int,graph map[int][]int){
	if node == nil{
		return
	}
	if parent != -1{
		graph[node.Val] = append(graph[node.Val],parent)
		graph[parent] = append(graph[parent],node.Val)
	}
	creategraph(node.Left,node.Val,graph)
	creategraph(node.Right,node.Val,graph)
}

func AmountOfTime(root *TreeNode, start int) int{
	var graph map[int][]int = make(map[int][]int)
	creategraph(root,-1,graph)
	var visited map[int]bool = make(map[int]bool)
	var q list.List
	q.PushBack(start)
	visited[start] = true
	var steps int = 0
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var node int = q.Front().Value.(int)
			q.Remove(q.Front())
			for _,neighbour := range graph[node]{
				if visited[neighbour]{
					continue
				}
				visited[neighbour] = true
				q.PushBack(neighbour)
			}
		}
		if q.Len() == 0{
			break
		}
		steps++
	}
	return steps
}

//func find_start(node *TreeNode,start int)(int,*TreeNode){
//	if node == nil{
//		return -1,nil
//	}
//	if node.Val == start{
//		return 0,node
//	}
//	res1,n := find_start(node.Left,start)
//	if res1 >= 0{
//		return res1 + 1,n
//	}
//	res2,n := find_start(node.Right,start)
//	if res2 >= 0{
//		return res2 + 1,n
//	}
//	return -1,nil
//}
//
//func calc_depth(node *TreeNode)int{
//	if node == nil{
//		return 0
//	}
//	return max_int(1 + calc_depth(node.Left),1 + calc_depth(node.Right))
//}
//
//func AmountOfTime(root *TreeNode, start int) int {
//	var start_root_dis int = 0
//	var root_another_dis int = 0
//	var node *TreeNode = nil
//	if root.Val != start{
//		start_root_dis,node = find_start(root.Left,start)
//		if start_root_dis == -1{
//			start_root_dis,node = find_start(root.Right,start)
//			root_another_dis = calc_depth(root.Left)
//		}else{
//			root_another_dis = calc_depth(root.Right)
//		}
//		start_end_dis := calc_depth(node) - 1
//		return max_int(start_end_dis,1 + start_root_dis + root_another_dis)
//	}else{
//		return calc_depth(root) - 1
//	}
//}