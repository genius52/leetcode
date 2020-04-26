package tree

//Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
//
//Calling next() will return the next smallest number in the BST.
type BSTIterator struct {
	data []int
	index int
}

func inorder_tree173(node *TreeNode,data *[]int){
	if node == nil{
		return
	}
	inorder_tree173(node.Left,data)
	*data = append(*data,node.Val)
	inorder_tree173(node.Right,data)
}

func Constructor(root *TreeNode) BSTIterator {
	var obj BSTIterator
	inorder_tree173(root,&obj.data)
	obj.index = 0
	return obj
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.HasNext(){
		pos := this.index
		this.index++
		return this.data[pos]
	}
	return -1
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.index < len(this.data){
		return true
	}
	return false
}
