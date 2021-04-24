package tree

func dfs_preorder(node *Node2,res *[]int){
	if node == nil{
		return
	}
	*res = append(*res,node.Val)
	for _,child := range node.Children{
		dfs_preorder(child,res)
	}
}

func preorder(root *Node2) []int {
	var res []int
	if root == nil{
		return res
	}
	dfs_preorder(root,&res)
	return res
}