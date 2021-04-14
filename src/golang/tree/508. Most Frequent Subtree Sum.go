package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func post_FindFrequentTreeSum(node *TreeNode,record map[int]int,max_count *int) int{
	if node == nil{
		return 0
	}
	var sum int = node.Val
	sum += post_FindFrequentTreeSum(node.Left,record,max_count)
	sum += post_FindFrequentTreeSum(node.Right,record,max_count)
	record[sum]++
	if record[sum] > *max_count{
		*max_count = record[sum]
	}
	return sum
}

func FindFrequentTreeSum(root *TreeNode) []int {
	var record map[int]int = make(map[int]int)
	var max_count int = 0
	post_FindFrequentTreeSum(root,record,&max_count)
	var res []int
	for sum,cnt := range record{
		if cnt == max_count{
			res = append(res,sum)
		}
	}
	return res
}