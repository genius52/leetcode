package tree

import "container/list"

func recursive_inorderTraversal(node *TreeNode,res *[]int){
	if node == nil{
		return
	}
	recursive_inorderTraversal(node.Left,res)
	*res = append(*res,node.Val)
	recursive_inorderTraversal(node.Right,res)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	recursive_inorderTraversal(root,&res)
	return res
}

func inorderTraversal2(root *TreeNode) []int{
	var res []int
	if root == nil{
		return res
	}
	var visit *TreeNode = root
	var q list.List
	for q.Len() > 0 || visit != nil{
		if visit != nil{
			q.PushBack(visit)
			visit = visit.Left
		}else{
			cur := q.Back().Value.(*TreeNode)
			q.Remove(q.Back())
			res = append(res,cur.Val)
			visit = cur.Right
		}
	}
	return res
}