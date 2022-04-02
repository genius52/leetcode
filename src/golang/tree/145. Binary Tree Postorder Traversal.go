package tree

import "container/list"

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil{
		return res
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var node *TreeNode = q.Back().Value.(*TreeNode)
		q.Remove(q.Back())
		res = append([]int{node.Val},res...)
		if node.Left != nil{
			q.PushBack(node.Left)
		}
		if node.Right != nil{
			q.PushBack(node.Right)
		}
	}
	return res
}

func recursive_postorderTraversal(node *TreeNode,res *[]int){
	if node == nil{
		return
	}
	recursive_postorderTraversal(node.Left,res)
	recursive_postorderTraversal(node.Right,res)
	*res = append(*res,node.Val)
}

func postorderTraversal2(root *TreeNode) []int {
	var res []int
	recursive_postorderTraversal(root,&res)
	return res
}