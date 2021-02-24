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