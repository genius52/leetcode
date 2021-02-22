package tree

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	var record [][]int
	if root == nil{
		return record
	}
	var level int = 0
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var node *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if len(record) <= level{
				record = append(record,[]int{node.Val})
			}else{
				record[level] = append(record[level],node.Val)
			}
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
		level++
	}
	return record
}