package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder_distanceK(parent *TreeNode, node *TreeNode,graph map[*TreeNode]*TreeNode){
	if node == nil{
		return
	}
	graph[node] = parent
	inorder_distanceK(node,node.Left,graph)
	inorder_distanceK(node,node.Right,graph)
}

func dfs_distanceK(node *TreeNode,K int,graph map[*TreeNode]*TreeNode,visited map[int]bool,res *[]int){
	if node == nil{
		return
	}
	if visited[node.Val]{
		return
	}
	visited[node.Val] = true
	if K == 0{
		*res = append(*res,node.Val)
	}
	if node.Left != nil{
		dfs_distanceK(node.Left,K - 1,graph,visited,res)
	}
	if node.Right != nil{
		dfs_distanceK(node.Right,K - 1,graph,visited,res)
	}
	if parent,ok := graph[node];ok{
		dfs_distanceK(parent,K - 1,graph,visited,res)
	}
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}
	var res []int
	var graph map[*TreeNode]*TreeNode = make(map[*TreeNode]*TreeNode)
	inorder_distanceK(root,root.Left,graph)
	inorder_distanceK(root,root.Right,graph)
	var visited map[int]bool = make(map[int]bool)
	dfs_distanceK(target,K,graph,visited,&res)
	return res
}