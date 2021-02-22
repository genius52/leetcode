package tree

import "container/list"

func levelOrderBottom(root *TreeNode) [][]int {
	var record [][]int
	if root == nil{
		return record
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var data []int
		for i := 0;i < l;i++{
			var node *TreeNode = q.Front().Value.(*TreeNode)
				data = append(data,node.Val)
			q.Remove(q.Front())
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
		record = append([][]int{data},record...)
	}
	return record
}