package tree

import "container/list"

//103
//[1,2,3,4,null,null,5]
//[3,9,20,null,null,15,7]
//func level_visit103(node *TreeNode,level int,res *[][]int){
//	if nil == node{
//		return
//	}
//	l := len(*res)
//	if l <= level{
//		*res = append(*res,[]int{node.Val})
//	}else{
//		(*res)[level] = append((*res)[level], node.Val)
//	}
//	level_visit103(node.Left,level + 1,res)
//	level_visit103(node.Right,level + 1,res)
//}
//
//func zigzagLevelOrder(root *TreeNode) [][]int {
//	var res [][]int
//	level_visit103(root,0,&res)
//	for i := 0; i < len(res);i++{
//		if i % 2 == 1{
//			for begin, end := 0, len(res[i])-1; begin < end; begin, end = begin+1, end-1 {
//				res[i][begin], res[i][end] = res[i][end], res[i][begin]
//			}
//		}
//	}
//	return res
//}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var record [][]int
	if root == nil{
		return record
	}
	var odd bool = false
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var data []int
		for i := 0;i < l;i++{
			var node *TreeNode = q.Front().Value.(*TreeNode)
			if odd{
				data = append([]int{node.Val},data...)
			}else{
				data = append(data,node.Val)
			}
			q.Remove(q.Front())
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
		record = append(record,data)
		odd = !odd
	}
	return record
}