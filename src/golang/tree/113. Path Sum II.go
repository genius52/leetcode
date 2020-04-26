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
func dfs_PathSum2(node *TreeNode,target int,cur_sum int,cur_list *[]int,record *[][]int){
	if node == nil{
		return
	}
	if node.Left == nil && node.Right == nil{
		if (cur_sum + node.Val) == target{
			*cur_list = append(*cur_list,node.Val)
			*record = append(*record,*cur_list)
		}
		return
	}
	*cur_list = append(*cur_list,node.Val)
	dfs_PathSum2(node.Left,target,cur_sum,cur_list,record)
	dfs_PathSum2(node.Right,target,cur_sum,cur_list,record)
	*cur_list = (*cur_list)[:len(*cur_list) - 1]
}

func PathSum2(root *TreeNode, sum int) [][]int {
	var res [][]int
	var cur_list []int
	dfs_PathSum2(root,sum,0,&cur_list,&res)
	return res
}
