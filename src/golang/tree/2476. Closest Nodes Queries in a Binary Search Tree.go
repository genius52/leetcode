package tree

import "sort"

func pre_closestNodes(node *TreeNode,data *[]int){
	if node == nil{
		return
	}
	pre_closestNodes(node.Left,data)
	*data = append(*data,node.Val)
	pre_closestNodes(node.Right,data)
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var data []int
	pre_closestNodes(root,&data)
	var l int = len(queries)
	var l2 int = len(data)
	var res [][]int = make([][]int,l)
	for j := 0;j < l;j++{
		res[j] = make([]int,2)
		find_idx := sort.Search(l2, func(i int) bool { //最小的使函数f(index)为True的值
			return data[i] >= queries[j]
		})
		if find_idx == l2{
			res[j][1] = -1
			res[j][0] = data[l2 - 1]
		}else{
			res[j][1] = data[find_idx]
			if queries[j] == data[find_idx]{
				res[j][0] = data[find_idx]
			}else{
				if find_idx > 0{
					res[j][0] = data[find_idx - 1]
				}else{
					res[j][0] = -1
				}

			}
		}
	}
	return res
}