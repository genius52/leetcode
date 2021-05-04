package tree

import "container/list"

//[["", "", "", "1", "", "", ""],
// ["", "2", "", "", "", "3", ""],
// ["", "", "4", "", "", "", ""]]
func get_depth(node *TreeNode)int{
	if node == nil{
		return 0
	}
	return 1 + max_int(get_depth(node.Left),get_depth(node.Right))
}

func preorder_printTree(node *TreeNode,cur_depth int,left int,right int,res [][]string){
	if node == nil{
		return
	}

}

//depth i + 1 = (depth i * 2) + 1
func printTree(root *TreeNode) [][]string {
	var res [][]string
	if root == nil{
		return res
	}
	depth := get_depth(root)

	res = make([][]string,depth)
	for i := 0;i < depth;i++{
		res[i] = make([]string,depth * 2 + 1)
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{

	}
	return res
}