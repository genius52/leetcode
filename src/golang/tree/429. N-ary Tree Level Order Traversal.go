package tree

import "container/list"
type Node2 struct {
	Val int
	Children []*Node2
}

func levelOrder429(root *Node2) [][]int {
	var res [][]int
	if root == nil{
		return res
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var cur []int
		for i := 0;i < l;i++{
			var node *Node2 = q.Front().Value.(*Node2)
			q.Remove(q.Front())
			cur = append(cur,node.Val)
			for _,child := range node.Children{
				q.PushBack(child)
			}
		}
		res = append(res,cur)
	}
	return res
}