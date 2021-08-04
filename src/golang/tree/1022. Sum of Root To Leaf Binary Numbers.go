package tree

//Input: root = [1,0,1,0,1,0,1]
//Output: 22
//Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
func recursive_sumRootToLeaf(node *TreeNode,pre int,res *int){
	if node == nil{
		return
	}
	pre = pre * 2 + node.Val
	if node.Left == nil && node.Right == nil{
		*res = *res + pre
		return
	}
	recursive_sumRootToLeaf(node.Left,pre,res)
	recursive_sumRootToLeaf(node.Right,pre,res)
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil{
		return 0
	}
	var res int = 0
	recursive_sumRootToLeaf(root,0,&res)
	return res
}