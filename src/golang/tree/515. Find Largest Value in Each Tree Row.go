package tree

import "container/list"

//515
func preorder_search_max(node *TreeNode,record *[]int,level int){
	if nil == node{
		return
	}
	if level == len(*record){
		*record = append(*record, node.Val)
	}else{
		if node.Val > (*record)[level]{
			(*record)[level] = node.Val
		}
	}
	preorder_search_max(node.Left,record,level + 1)
	preorder_search_max(node.Right,record,level + 1)
}

func largestValues(root *TreeNode) []int {
	var res []int
	if nil == root{
		return res
	}
	preorder_search_max(root,&res,0)
	return res
}

func largestValues2(root *TreeNode) []int{
	var res []int
	if root == nil{
		return res
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var max_num int = -2147483648
		for i := 0;i < l;i++{
			node := q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if node.Val > max_num{
				max_num = node.Val
			}
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
		res = append(res,max_num)
	}
	return res
}