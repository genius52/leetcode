package tree

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
func dfs_createBinaryTree(descriptions [][]int,cur int,record map[int]*TreeNode){
	var n *TreeNode
	if node,ok := record[descriptions[cur][1]];ok{
		n = node
	}else{
		n = new(TreeNode)
		n.Val = descriptions[cur][1]
		record[n.Val] = n
	}

	if parent,ok := record[descriptions[cur][0]];ok{
		if descriptions[cur][2] == 1{
			parent.Left = n
		}else{
			parent.Right = n
		}
	}else{
		var n2 *TreeNode = new(TreeNode)
		n2.Val = descriptions[cur][0]
		if descriptions[cur][2] == 1{
			n2.Left = n
		}else{
			n2.Right = n
		}
		record[n2.Val] = n2
	}
}

func CreateBinaryTree(descriptions [][]int) *TreeNode {
	var record map[int]*TreeNode = make(map[int]*TreeNode)//node's value : node pointer
	var graph map[int]int = make(map[int]int)
	var l int = len(descriptions)
	for i := 0;i < l;i++{
		graph[descriptions[i][1]] = descriptions[i][0]
		dfs_createBinaryTree(descriptions,i,record)
	}
	var head *TreeNode
	for _,v := range graph{
		var next int = v
		for true{
			if node,ok := record[next];ok{
				head = node
				next = graph[next]
			}else{
				break
			}
		}
		break
	}
	return head
}