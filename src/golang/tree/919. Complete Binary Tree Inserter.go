package tree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Input: inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
//Output: [null,3,4,[1,2,3,4,5,6,7,8]]
type CBTInserter struct {
	root *TreeNode
	notfull *list.List
}

func Constructor919(root *TreeNode) CBTInserter {
	var obj CBTInserter
	obj.root = root
	obj.notfull = new(list.List)
	if root == nil{
		return obj
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var node *TreeNode = q.Front().Value.(*TreeNode)
		q.Remove(q.Front())
		if node.Left != nil{
			q.PushBack(node.Left)
		}
		if node.Right != nil{
			q.PushBack(node.Right)
		}
		if node.Left == nil || node.Right == nil{
			obj.notfull.PushBack(node)
		}
	}
	return obj
}

func (this *CBTInserter) Insert(v int) int {
	front := this.notfull.Front()
	var node *TreeNode = front.Value.(*TreeNode)
	var new_node TreeNode
	new_node.Val = v
	if node.Left == nil{
		node.Left = &new_node
	}else if node.Right == nil{
		node.Right = &new_node
		this.notfull.Remove(front)
	}
	this.notfull.PushBack(&new_node)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

