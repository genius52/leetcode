package tree


//105
//preorder = [3,9,20,15,7]
//inorder = [9,3,15,20,7]
//Return the following binary tree:
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
func recursive_buildTree(preorder []int,pre_start int,pre_end int,inorder []int,in_start int,in_end int) *TreeNode{
	if pre_start > pre_end || in_start> in_end {
		return nil
	}
	var node *TreeNode = new(TreeNode)
	node.Val = preorder[pre_start]
	find_inorder_index := in_start
	for ;find_inorder_index <= in_end;find_inorder_index++{
		if preorder[pre_start] == inorder[find_inorder_index]{
			break
		}
	}
	node.Left = recursive_buildTree(preorder,pre_start + 1,pre_end,inorder,in_start,find_inorder_index-1)
	node.Right = recursive_buildTree(preorder,pre_start + 1 - in_start + find_inorder_index ,pre_end,inorder,find_inorder_index+1,in_end)
	return node
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}
	return recursive_buildTree(preorder,0,len(preorder) - 1,inorder,0,len(inorder) - 1)
}