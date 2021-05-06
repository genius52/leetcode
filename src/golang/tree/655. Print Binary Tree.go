package tree

import "strconv"

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
	cur_pos := (left + right)/2
	res[cur_depth][cur_pos] = strconv.Itoa(node.Val)
	preorder_printTree(node.Left,cur_depth + 1,left,cur_pos - 1,res)
	preorder_printTree(node.Right,cur_depth + 1,cur_pos + 1,right,res)
}

func PrintTree(root *TreeNode) [][]string {
	var res [][]string
	if root == nil{
		return res
	}
	depth := get_depth(root)
	var columns int = 1
	for i := 1;i < depth;i++{
		columns = columns * 2 + 1
	}
	res = make([][]string,depth)
	for i := 0;i < depth;i++{
		res[i] = make([]string,columns)
	}
	preorder_printTree(root,0,0,columns - 1,res)
	return res
}