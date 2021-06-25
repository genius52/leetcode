package tree

func preorder_leafSimilar(node *TreeNode,res *[]int){
	if node == nil{
		return
	}
	preorder_leafSimilar(node.Left,res)
	if node.Left == nil && node.Right == nil{
		*res = append(*res,node.Val)
	}
	preorder_leafSimilar(node.Right,res)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil{
		return true
	}
	var res1 []int
	preorder_leafSimilar(root1,&res1)
	var res2 []int
	preorder_leafSimilar(root2,&res2)
	var l1 int = len(res1)
	var l2 int = len(res2)
	if l1 != l2{
		return false
	}
	for i := 0;i < l1;i++{
		if res1[i] != res2[i]{
			return false
		}
	}
	return true
}