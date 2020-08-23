package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func post_FindFrequentTreeSum(node *TreeNode,record map[int]int) int{
	if node == nil{
		return 0
	}
	var sum int = node.Val
	sum += post_FindFrequentTreeSum(node.Left,record)
	sum += post_FindFrequentTreeSum(node.Right,record)
	record[sum]++
	return sum
}

func FindFrequentTreeSum(root *TreeNode) []int {
	var record map[int]int = make(map[int]int)
	post_FindFrequentTreeSum(root,record)
	var res []int
	var max_occurs int = 0
	for _,cnt := range record{
		if cnt > max_occurs{
			max_occurs = cnt
		}
	}
	for sum,cnt := range record{
		if cnt == max_occurs{
			res = append(res,sum)
		}
	}
	return res
}