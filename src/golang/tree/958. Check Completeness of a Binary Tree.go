package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func visit_isCompleteTree(node *TreeNode,parent int,cur int,record map[int][]int){
	if node == nil{
		return
	}
	record[parent] = append(record[parent],cur)
	visit_isCompleteTree(node.Left,cur,cur * 2,record)
	visit_isCompleteTree(node.Right,cur,cur * 2 + 1,record)
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil{
		return true
	}
	var record map[int][]int = make(map[int][]int)
	visit_isCompleteTree(root,0,1,record)
	var total int = 0
	var max_num int = 0
	for _,nums := range record{
		for _,n := range nums{
			total++
			if n > max_num{
				max_num = n
			}
		}
	}
	return total == max_num
}