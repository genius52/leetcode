package tree

func AllPossibleFBT(N int)[]*TreeNode{
	var res []*TreeNode
	if N % 2 == 0{
		return res
	}
	if N == 1{
		var node TreeNode
		node.Val = 0
		res = append(res,&node)
		return res
	}
	for left := 1; left < N - 1;left += 2{
		left_nodes := AllPossibleFBT(left)
		right_nodes := AllPossibleFBT(N - left - 1)
		for _,ln := range left_nodes{
			for _,rn := range right_nodes{
				var node TreeNode
				node.Val = 0
				node.Left = ln
				node.Right = rn
				res = append(res, &node)
			}
		}
	}
	return res
}