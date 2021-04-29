package tree

import "container/list"

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil{
		return res
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var sum int = 0
		var cnt int = q.Len()
		for i := 0;i < cnt;i++{
			node := q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			sum += node.Val
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
		res = append(res,float64(sum)/float64(cnt))
	}
	return res
}