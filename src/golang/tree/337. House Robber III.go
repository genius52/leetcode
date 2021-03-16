package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dp_rob(node *TreeNode,is_rob bool,record map[*TreeNode][]int)(ret int){
	if nil == node{
		return 0
	}

	if val,ok := record[node];ok{
		if is_rob {
			if val[1] >= 0{
				return val[1]
			}
		}else {
			if val[0] >= 0{
				return val[0]
			}
		}
	}
	var max int = 0
	if is_rob{
		max = dp_rob(node.Left,false,record) + dp_rob(node.Right,false,record)
	}else{
		max = max_int(node.Val + dp_rob(node.Left,true,record) + dp_rob(node.Right,true,record),dp_rob(node.Left,false,record) + dp_rob(node.Right,false,record))
	}
	record[node] = make([]int,2)
	if is_rob{
		record[node][1] = max
		record[node][0] = -1
	}else{
		record[node][0] = max
		record[node][1] = -1
	}
	return max
}

func rob(root *TreeNode) int {
	if nil == root{
		return 0
	}
	var record map[*TreeNode][]int = make(map[*TreeNode][]int)
	return dp_rob(root,false,record)
}