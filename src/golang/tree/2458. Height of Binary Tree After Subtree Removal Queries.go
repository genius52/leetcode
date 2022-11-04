package tree

func getheight_treeQueries(node *TreeNode,val_height map[*TreeNode]int)int{
	if node == nil{
		return 0
	}
	val_height[node] = 1 + max_int(getheight_treeQueries(node.Left,val_height),getheight_treeQueries(node.Right,val_height))
	return val_height[node]
}

func dfs_treeQueries(node *TreeNode,depth int,max_rest_height int,val_height map[*TreeNode]int,max_depth_without_cur []int){
	if node == nil{
		return
	}
	max_depth_without_cur[node.Val] = max_rest_height
	depth++
	if node.Left != nil{
		dfs_treeQueries(node.Right,depth,max_int(max_rest_height,depth + val_height[node.Left]),val_height,max_depth_without_cur)
	}else{
		dfs_treeQueries(node.Right,depth,max_int(max_rest_height,depth),val_height,max_depth_without_cur)
	}
	if node.Right != nil{
		dfs_treeQueries(node.Left,depth,max_int(max_rest_height,depth + val_height[node.Right]),val_height,max_depth_without_cur)
	}else{
		dfs_treeQueries(node.Left,depth,max_int(max_rest_height,depth),val_height,max_depth_without_cur)
	}
}

func treeQueries(root *TreeNode, queries []int) []int{
	//var val_node map[int]*TreeNode = make(map[int]*TreeNode)
	var val_height map[*TreeNode]int = map[*TreeNode]int{}
	getheight_treeQueries(root,val_height)//计算每个节点的高度
	var l int = len(queries)
	var max_depth_without_cur []int = make([]int,len(val_height) + 1)
	dfs_treeQueries(root,-1,0,val_height,max_depth_without_cur)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		res[i] = max_depth_without_cur[queries[i]]
	}
	return res
}

//func dfs_treeQueries(node *TreeNode,depth int,val_node map[int]*TreeNode,val_depth map[int]int,depth_cnt map[int]int){
//	if node == nil{
//		return
//	}
//	val_node[node.Val] = node
//	val_depth[node.Val] = depth
//	depth_cnt[depth]++
//	dfs_treeQueries(node.Left,depth + 1,val_node,val_depth,depth_cnt)
//	dfs_treeQueries(node.Right,depth + 1,val_node,val_depth,depth_cnt)
//}
//
//func dfs2_treeQueries(node *TreeNode,val_depth map[int]int,depth_cnt map[int]int,change map[int]int){
//	if node == nil{
//		return
//	}
//	depth_cnt[val_depth[node.Val]]--
//	change[val_depth[node.Val]]++
//	dfs2_treeQueries(node.Left,val_depth,depth_cnt,change)
//	dfs2_treeQueries(node.Right,val_depth,depth_cnt,change)
//}
//
//func treeQueries(root *TreeNode, queries []int) []int {
//	var val_node map[int]*TreeNode = make(map[int]*TreeNode)
//	var val_depth map[int]int = make(map[int]int)
//	var depth_cnt map[int]int = make(map[int]int)
//	dfs_treeQueries(root,0,val_node,val_depth,depth_cnt)
//	var res []int = make([]int,len(queries))
//	for i := 0;i < len(queries);i++{
//		var node int = queries[i]
//		//var depth_cnt2 map[int]int = depth_cnt
//		var change map[int]int = make(map[int]int)
//		dfs2_treeQueries(val_node[node],val_depth,depth_cnt,change)
//		var max_depth int = 0
//		for k,v := range depth_cnt{
//			if v == 0{
//				continue
//			}
//			if k > max_depth{
//				max_depth = k
//			}
//		}
//		res[i] = max_depth
//		for k,v := range change{
//			depth_cnt[k] += v
//		}
//	}
//	return res
//}