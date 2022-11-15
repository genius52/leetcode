package tree

import (
	"container/list"
	"sort"
)

func minimumOperations(root *TreeNode) int {
	var q list.List
	q.PushBack(root)
	var res int = 0
	for q.Len() > 0{
		var cur_len int = q.Len()
		var data []int
		for i := 0;i < cur_len;i++{
			var cur *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			data = append(data,cur.Val)
			if cur.Left != nil{
				q.PushBack(cur.Left)
			}
			if cur.Right != nil{
				q.PushBack(cur.Right)
			}
		}
		var l int = len(data)
		var data2 []int = make([]int,l)
		copy(data2,data)
		sort.Ints(data2)
		var data_idx map[int]int = make(map[int]int)
		for i := 0;i < l;i++{
			data_idx[data[i]] = i
		}
		for i := 0;i < l;i++{
			if data[i] == data2[i]{
				continue
			}
			swap_idx := data_idx[data2[i]]
			pre_val := data[i]
			data[i],data[data_idx[data2[i]]] = data2[i],data[i]
			data_idx[pre_val] = swap_idx
			res++
		}
	}
	return res
}