package tree

import "math"

func visit_WidthOfBinaryTree(node *TreeNode,level int,pos int,record map[int][]int){
	if node == nil{
		return
	}
	if _,ok := record[level];ok{
		record[level] = append(record[level],pos)
	}else{
		record[level] = []int{pos}
	}
	if pos < 0{
		visit_WidthOfBinaryTree(node.Left,level + 1,pos * 2,record)
		visit_WidthOfBinaryTree(node.Right,level + 1,pos * 2 + 1,record)
	}else{
		visit_WidthOfBinaryTree(node.Left,level + 1,pos * 2 - 1,record)
		visit_WidthOfBinaryTree(node.Right,level + 1,pos * 2,record)
	}
}

func WidthOfBinaryTree(root *TreeNode) int {
	if root == nil{
		return 0
	}
	var record map[int][]int = make(map[int][]int)
	visit_WidthOfBinaryTree(root.Left,1,-1,record)
	visit_WidthOfBinaryTree(root.Right,1,1,record)
	var diff int = 1
	for _,positions := range record{
		var max int = math.MinInt32
		var min int = math.MaxInt32
		if len(positions) == 1{
			continue
		}
		for _,p := range positions{
			if p > max{
				max = p
			}
			if p < min{
				min = p
			}
		}
		if max > 0 && min < 0{
			if (max - min) > diff{
				diff = max - min
			}
		}else{
			if (max - min + 1) > diff{
				diff = max - min + 1
			}
		}
	}
	return diff
}
