package tree

import "container/list"

func recursive_preorderTraversal(node *TreeNode,res *[]int){
	if node == nil{
		return
	}
	*res = append(*res,node.Val)
	recursive_preorderTraversal(node.Left,res)
	recursive_preorderTraversal(node.Right,res)
}

func preorderTraversal2(root *TreeNode) []int{
	var res []int
	recursive_preorderTraversal(root,&res)
	return res
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil{
		return res
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var node *TreeNode = q.Back().Value.(*TreeNode)
		q.Remove(q.Back())
		res = append(res,node.Val)
		if node.Right != nil{
			q.PushBack(node.Right)
		}
		if node.Left != nil{
			q.PushBack(node.Left)
		}
	}
	return res
}