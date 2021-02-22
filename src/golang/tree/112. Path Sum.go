package tree

//Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//Output: true
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil{
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil{
		return true
	}
	return hasPathSum(root.Left,targetSum - root.Val) || hasPathSum(root.Right,targetSum - root.Val)
}