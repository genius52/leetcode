package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorder_CountPairs(node *TreeNode,distance int,total *int)map[int]int{
	if(node == nil){
		return nil
	}
	if(node.Left == nil && node.Right == nil){
		m := make(map[int]int)
		m[1] = 1
		return m
	}
	left_map := postorder_CountPairs(node.Left,distance,total)
	right_map := postorder_CountPairs(node.Right,distance,total)
	if(left_map == nil){
		var m map[int]int = make(map[int]int)
		for k,v := range right_map{
			m[k + 1] = v
		}
		return m
	}
	if(right_map == nil){
		var m map[int]int = make(map[int]int)
		for k,v := range left_map{
			m[k + 1] = v
		}
		return m
	}
	var m map[int]int
	for l_dis,l_cnt := range left_map{
		if(l_dis > distance){
			continue
		}
		for r_dis,r_cnt := range right_map{
			if(r_dis > distance){
				continue
			}
			if((l_dis + r_dis) <= distance){
				*total = *total + l_cnt * r_cnt
			}
		}
	}
	for l_dis,l_cnt := range left_map {
		if (l_dis > distance) {
			continue
		}
		if (m == nil) {
			m = make(map[int]int)
		}
		if _, ok := m[l_dis+1]; ok {
			m[l_dis+1] += l_cnt
		} else {
			m[l_dis+1] = l_cnt
		}
	}
	for r_dis,r_cnt := range right_map{
		if(r_dis > distance){
			continue
		}
		if (m == nil) {
			m = make(map[int]int)
		}
		if _,ok := m[r_dis + 1];ok{
			m[r_dis + 1] += r_cnt
		}else{
			m[r_dis + 1] = r_cnt
		}
	}
	return m
}

func CountPairs(root *TreeNode, distance int) int {
	if(root == nil){
		return 0
	}
	var total int = 0
	postorder_CountPairs(root,distance,&total)
	return total
}