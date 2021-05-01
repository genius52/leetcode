package tree

func recursive_constructMaximumBinaryTree(nums []int,left int,right int) *TreeNode{
	if left > right{
		return nil
	}
	var root TreeNode
	var max_num int = -2147483648
	var max_num_index int = 0
	for i := left;i <= right;i++{
		if nums[i] > max_num{
			max_num = nums[i]
			max_num_index = i
		}
	}
	root.Val = max_num
	root.Left = recursive_constructMaximumBinaryTree(nums,left,max_num_index - 1)
	root.Right = recursive_constructMaximumBinaryTree(nums,max_num_index + 1,right)
	return &root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var l int = len(nums)
	return recursive_constructMaximumBinaryTree(nums,0,l - 1)
}