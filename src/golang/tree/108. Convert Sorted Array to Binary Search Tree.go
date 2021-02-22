package tree

func recursive_sortedArrayToBST(nums []int,low int,high int)*TreeNode{
	if low > high{
		return nil
	}
	mid := low + (high - low)/ 2
	var node *TreeNode = new(TreeNode)
	node.Val = nums[mid]
	node.Left = recursive_sortedArrayToBST(nums,low,mid - 1)
	node.Right = recursive_sortedArrayToBST(nums,mid + 1,high)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	var l int = len(nums)
	return recursive_sortedArrayToBST(nums,0,l - 1)
}