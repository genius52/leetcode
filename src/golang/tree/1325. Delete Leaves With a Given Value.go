package tree

func removeLeafNodes(root *TreeNode, target int) *TreeNode{
	if root == nil{
		return nil
	}
	root.Left = removeLeafNodes(root.Left,target)
	root.Right = removeLeafNodes(root.Right,target)
	if root.Left == nil && root.Right == nil{
		if root.Val == target{
			return nil
		}
	}
	return root
}

//func postvisit_removeLeafNodes(node *TreeNode,target int)bool{
//	if nil == node{
//		return true
//	}
//	left_res :=  postvisit_removeLeafNodes(node.Left,target)
//	right_res := postvisit_removeLeafNodes(node.Right,target)
//	if left_res{
//		node.Left = nil
//	}
//	if right_res{
//		node.Right = nil
//	}
//	if left_res && right_res{
//		if node.Val == target{
//			return true
//		}
//	}
//	return false
//}
//
//func removeLeafNodes(root *TreeNode, target int) *TreeNode {
//	if postvisit_removeLeafNodes(root,target){
//		return nil
//	}else{
//		return root
//	}
//}