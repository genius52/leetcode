package tree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//BFS solution
func cnt_isCompleteTree(node *TreeNode) int{
	if node == nil{
		return 0
	}
	return 1 + cnt_isCompleteTree(node.Left) + cnt_isCompleteTree(node.Right)
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil{
		return true
	}
	var q list.List
	q.PushBack(root)
	var empty_child bool = false
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if cur.Left == nil && cur.Right != nil{
				return false
			}
			if !empty_child{
				if cur.Left == nil || cur.Right == nil{ //都为空，右边为空
					empty_child = true
				}
				if cur.Left != nil{
					q.PushBack(cur.Left)
				}
				if cur.Right != nil{
					q.PushBack(cur.Right)
				}
			}else{
				if cur.Left != nil || cur.Right != nil{
					return false
				}
			}
		}
	}
	return true
}



//func visit_isCompleteTree(node *TreeNode,parent int,cur int,record map[int][]int){
//	if node == nil{
//		return
//	}
//	record[parent] = append(record[parent],cur)
//	visit_isCompleteTree(node.Left,cur,cur * 2,record)
//	visit_isCompleteTree(node.Right,cur,cur * 2 + 1,record)
//}
//
//func isCompleteTree(root *TreeNode) bool {
//	if root == nil{
//		return true
//	}
//	var record map[int][]int = make(map[int][]int)
//	visit_isCompleteTree(root,0,1,record)
//	var total int = 0
//	var max_num int = 0
//	for _,nums := range record{
//		for _,n := range nums{
//			total++
//			if n > max_num{
//				max_num = n
//			}
//		}
//	}
//	return total == max_num
//}