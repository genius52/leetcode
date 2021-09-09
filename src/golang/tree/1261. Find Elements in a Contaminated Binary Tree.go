package tree

type FindElements struct {
	root *TreeNode
}

func dfs_tree(node *TreeNode,val int){
	if nil == node{
		return
	}
	node.Val = val
	dfs_tree(node.Left,node.Val*2+1)
	dfs_tree(node.Right,node.Val*2+2)
}

func Constructor2(root *TreeNode) FindElements {
	var obj FindElements
	obj.root = root
	dfs_tree(obj.root,0)
	return obj
}

func find(node *TreeNode,target int)bool{
	if node == nil{
		return false
	}
	if node.Val == target{
		return true
	}
	if target < node.Val{
		return false
	}
	return find(node.Left,target) || find(node.Right,target)
}

func (this *FindElements) Find(target int) bool {
	return find(this.root,target)
}



//type FindElements struct {
//	root *TreeNode
//	set map[int]bool
//}
//
//func dfs_tree(node *TreeNode,val int,ele FindElements){
//	if nil == node{
//		return
//	}
//	node.Val = val
//	ele.set[node.Val] = true
//	dfs_tree(node.Left,node.Val*2+1,ele)
//	dfs_tree(node.Right,node.Val*2+2,ele)
//}
//
//func Constructor2(root *TreeNode) FindElements {
//	root.Val = 0
//	ele := FindElements{root,make(map[int]bool)}
//	ele.set[root.Val] = true
//	dfs_tree(root.Left,root.Val*2+1,ele)
//	dfs_tree(root.Right,root.Val*2+2,ele)
//	return ele
//}
//
//func (this *FindElements) Find(target int) bool {
//	if _,ok := this.set[target];ok{
//		return true
//	}
//	return false
//}