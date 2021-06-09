package tree

//remove zero subtree
func pruneTree(root *TreeNode) *TreeNode{
	if root == nil{
		return root
	}
	left := pruneTree(root.Left)
	right := pruneTree(root.Right)
	//remove zero node
	if root.Val == 0 && left == nil && right == nil{
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}

//
//func pruneTree(root *TreeNode) *TreeNode {
//	if(nil == root){
//		return root;
//	}
//	if(nil != root.Left){
//		root.Left = pruneTree(root.Left)
//	}
//	if(nil != root.Right){
//		root.Right = pruneTree(root.Right)
//	}
//	if(0 == root.Val){
//		if(nil == root.Left && nil == root.Right){
//			return nil
//		}
//	}
//	return root
//}
