package tree


//106
//inorder = [9,3,15,20,7]
//postorder = [9,15,7,20,3]
//Return the following binary tree:
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
func inpost_buildTree(inorder []int,in_begin int,in_end int,postorder []int,root_pos int)*TreeNode{
	if root_pos < 0 || in_begin > in_end{
		return nil
	}
	//find root's position in inorder
	var i int = in_begin
	for i < in_end && inorder[i] != postorder[root_pos]{
		i++
	}
	//length of right tree nodes
	var right_tree_length int = in_end - i + 1
	var node TreeNode
	node.Val = postorder[root_pos]
	node.Left = inpost_buildTree(inorder,in_begin,i - 1,postorder,root_pos - right_tree_length)
	node.Right = inpost_buildTree(inorder,i + 1, in_end,postorder,root_pos - 1)
	return &node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	lin := len(inorder)
	lpost := len(postorder)
	if lin == 0 || lpost == 0 || lin != lpost{
		return nil
	}
	return inpost_buildTree(inorder,0,lin - 1,postorder,lpost - 1)
}