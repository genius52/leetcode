package tree

import "container/list"

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil{
		return res
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var add bool = false
		for i := 0;i < l;i++{
			var node *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if !add{
				res = append(res,node.Val)
				add = true
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
			if node.Left != nil{
				q.PushBack(node.Left)
			}
		}
	}
	return res
}