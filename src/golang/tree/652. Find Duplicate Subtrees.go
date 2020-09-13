package tree

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func post_findDuplicateSubtrees(node *TreeNode,record map[string][]*TreeNode,res []*TreeNode)string{
	var s string
	if node == nil{
		return s
	}
	s += strconv.Itoa(node.Val)
	left_str := post_findDuplicateSubtrees(node.Left,record,res)
	right_str := post_findDuplicateSubtrees(node.Right,record,res)
	if len(left_str) > 0 {
		s += "l(" + left_str + ")"
	}
	if len(right_str) > 0{
		s += "r(" + right_str + ")"
	}
	record[s] = append(record[s],node)
	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var record map[string][]*TreeNode = make(map[string][]*TreeNode)
	var res []*TreeNode
	post_findDuplicateSubtrees(root,record,res)
	for _,nodes := range record{
		if len(nodes) < 2{
			continue
		}
		res = append(res,nodes[0])
	}
	return res
}