package tree

import "strconv"

//Output: ["1->2->5", "1->3"]
func dfs_binaryTreePaths(node *TreeNode,pre string,res *[]string){
	if node == nil{
		return
	}
	var cur string
	if len(pre) == 0{
		cur = strconv.Itoa(node.Val)
	}else{
		cur += pre + "->" + strconv.Itoa(node.Val)
	}
	if node.Left == nil && node.Right == nil{
		*res = append(*res,cur)
		return
	}
	dfs_binaryTreePaths(node.Left,cur,res)
	dfs_binaryTreePaths(node.Right,cur,res)
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	dfs_binaryTreePaths(root,"",&res)
	return res
}