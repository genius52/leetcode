package tree

//Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given the below binary tree and sum = 22,
//
//      5
//     / \
//    4   8
//   /   / \
//  11  13  4
// /  \    / \
//7    2  5   1
//Return:
//
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]
func dfs_PathSum2(node *TreeNode,target int,cur_sum int,path []int,record *[][]int){
	if node == nil{
		return
	}
	cur_sum += node.Val
	path = append(path,node.Val)
	if node.Left == nil && node.Right == nil{
		if cur_sum == target{
			var data []int = make([]int,len(path))
			copy(data,path)
			*record = append(*record,data)
		}
		path = (path)[:len(path) - 1]
		return
	}
	dfs_PathSum2(node.Left,target,cur_sum,path,record)
	dfs_PathSum2(node.Right,target,cur_sum,path,record)
}

func PathSum2(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	dfs_PathSum2(root,sum,0,path,&res)
	return res
}