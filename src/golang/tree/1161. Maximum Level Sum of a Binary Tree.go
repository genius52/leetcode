package tree

import "container/list"

func maxLevelSum(root *TreeNode) int {
	var q list.List
	var level int = 1
	var record map[int]int = make(map[int]int)
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var sum int = 0
		for i := 0;i < l;i++{
			cur := q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			sum += cur.Val
			if cur.Left != nil{
				q.PushBack(cur.Left)
			}
			if cur.Right != nil{
				q.PushBack(cur.Right)
			}
		}
		record[level] = sum
		level++
	}
	var max_sum int = -2147483648
	var max_sum_level int = -1
	for k,v := range record{
		if v > max_sum{
			max_sum = v
			max_sum_level = k
		}else if v == max_sum{
			if k < max_sum_level{
				max_sum_level = k
			}
		}
	}
	return max_sum_level
}

//func level_visit(node *TreeNode,level int,record map[int]int){
//	if nil == node{
//		return
//	}
//	if val,ok := record[level];ok{
//		record[level] += val
//	}else{
//		record[level] = node.Val
//	}
//	level_visit(node.Left,level+1,record)
//	level_visit(node.Right,level+1,record)
//}
//
//func maxLevelSum(root *TreeNode) int {
//	var record map[int]int = make(map[int]int)
//	level_visit(root,1,record)
//	var max int = math.MinInt32
//	var res int
//	for index,val := range record{
//		if val > max{
//            max = val
//			res = index
//		}
//	}
//	return res
//}