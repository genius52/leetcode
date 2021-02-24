package tree

import "sort"

func inorder_getvalue(node *TreeNode,data* []int){
	if node == nil{
		return
	}
	inorder_getvalue(node.Left,data)
	*data = append(*data,node.Val)
	inorder_getvalue(node.Right,data)
}

func inorder_recoverTree(node *TreeNode,data []int,index *int,total *int){
	if node == nil{
		return
	}
	inorder_recoverTree(node.Left,data,index,total)
	if *total == 0{
		return
	}
	if node.Val != data[*index]{
		*total--
		node.Val = data[*index]
		if *total == 0{
			return
		}
	}
	*index++
	inorder_recoverTree(node.Right,data,index,total)
}

func RecoverTree(root *TreeNode)  {
	if root == nil{
		return
	}
	var data []int
	inorder_getvalue(root,&data)
	sort.Ints(data)
	var index int = 0
	var total int = 2
	inorder_recoverTree(root,data,&index,&total)
}
