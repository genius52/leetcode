package tree

func pseudoPalindromicPaths2(node *TreeNode,record *[10]int)int{
	if node == nil{
		return 0
	}
	(*record)[node.Val]++
	var res int = 0
	if node.Left == nil && node.Right == nil{
		var odd_cnt int = 0
		for i := 0;i <= 9;i++{
			if (*record)[i] | 1 == (*record)[i]{
				odd_cnt++
				if odd_cnt > 1{
					break
				}
			}
		}
		if odd_cnt <= 1{
			res = 1
		}
	}else{
		res = pseudoPalindromicPaths2(node.Left,record) + pseudoPalindromicPaths2(node.Right,record)
	}
	(*record)[node.Val]--
	return res
}

func pseudoPalindromicPaths (root *TreeNode) int {
	var record [10]int
	return pseudoPalindromicPaths2(root,&record)
}