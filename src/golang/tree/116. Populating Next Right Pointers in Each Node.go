package tree

import "container/list"

func connect(root *Node) *Node {
	if root == nil{
		return nil
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var pre *Node = nil
		for i := 0;i < l;i++{
			var node *Node = q.Front().Value.(*Node)
			q.Remove(q.Front())
			if pre != nil{
				pre.Next = node
			}
			pre = node
			pre.Next = nil
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
	}
	return root
}