package tree


func findMode(root *TreeNode) []int {
	var res []int
	if nil == root{
		return res
	}
	var cnt_map map[int]int = make(map[int]int)
	depth_visit(root,cnt_map)
	max_occurs := 0
	for _,v := range cnt_map{
		if v > max_occurs{
			max_occurs = v
		}
	}
	for k,v := range cnt_map{
		if v == max_occurs{
			res = append(res, k)
		}
	}
	return res
}

var cur_val int = 0
var tmp_cnt int = 0
var max_cnt int = 0

func inorder(node *TreeNode,res *[]int){
	if nil == node{
		return
	}
	inorder(node.Left,res)
	if node.Val != cur_val{
		cur_val = node.Val
		tmp_cnt = 1
	}else{
		tmp_cnt++
	}
	if tmp_cnt > max_cnt{
		*res = (*res)[0:0]
		*res = append(*res,cur_val)
		max_cnt = tmp_cnt
	} else if tmp_cnt == max_cnt{
		*res = append(*res,cur_val)
	}
	inorder(node.Right,res)
}

func findMode2(root *TreeNode) []int{
	var res []int
	if nil == root{
		return res
	}
	inorder(root,&res)
	return res
}