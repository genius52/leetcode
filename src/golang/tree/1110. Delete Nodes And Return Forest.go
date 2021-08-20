package tree

func recursive_delNodes(node *TreeNode,del map[int]bool,res *[]*TreeNode)*TreeNode{
	if node == nil{
		return nil
	}
	if _,ok := del[node.Val];ok{
		left := recursive_delNodes(node.Left,del,res)
		right := recursive_delNodes(node.Right,del,res)
		if left != nil{
			*res = append(*res,left)
		}
		if right != nil{
			*res = append(*res,right)
		}
		return nil
	}
	node.Left = recursive_delNodes(node.Left,del,res)
	node.Right = recursive_delNodes(node.Right,del,res)
	return node
}

func delNodes2(root *TreeNode, to_delete []int) []*TreeNode {
	var res []*TreeNode
	var del map[int]bool = make(map[int]bool)
	for _,n := range to_delete{
		del[n] = true
	}
	node := recursive_delNodes(root,del,&res)
	if node != nil{
		res = append(res,node)
	}
	return res
}


//previous solution
func inorder_visit(node **TreeNode,init bool,to_delete []int,res *[]*TreeNode){
	var should_delete bool = false
	if nil == *node{
		return
	}
	for _,val := range to_delete{
		if val == (*node).Val{
			should_delete = true
			break
		}
	}
	temp := *node
	if should_delete{
		*node = nil
		init = true
	}else{
		if init{
			*res = append(*res, *node)
			init = false
		}
	}
	inorder_visit(&(temp.Left),init,to_delete,res)
	inorder_visit(&(temp.Right),init,to_delete,res)
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var res []*TreeNode
	inorder_visit(&root,true,to_delete,&res)
	return res
}