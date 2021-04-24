package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func dfs_postorder(node *Node2,res *[]int){
	if node == nil{
		return
	}
	for _,child := range node.Children{
		dfs_postorder(child,res)
	}
	*res = append(*res,node.Val)
}

func postorder(root *Node2) []int {
	var res []int
	dfs_postorder(root,&res)
	return res
}