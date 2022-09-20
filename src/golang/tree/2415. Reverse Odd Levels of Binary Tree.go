package tree

import "container/list"

func ReverseOddLevels(root *TreeNode) *TreeNode {
	var q1 list.List
	q1.PushBack(root)
	var q2 list.List
	var is_odd bool = false
	for q1.Len() > 0{
		var cur_len int = q1.Len()
		for i := 0;i < cur_len;i++{
			var node *TreeNode = q1.Front().Value.(*TreeNode)
			q1.Remove(q1.Front())
			if node.Left != nil{
				q1.PushBack(node.Left)
			}
			if node.Right != nil{
				q1.PushBack(node.Right)
			}
			if is_odd{
				var val int = q2.Front().Value.(int)
				q2.Remove(q2.Front())
				node.Val = val
			}else{
				if node.Left != nil{
					q2.PushFront(node.Left.Val)
				}
				if node.Right != nil{
					q2.PushFront(node.Right.Val)
				}
			}
		}
		is_odd = !is_odd
	}
	return root
}