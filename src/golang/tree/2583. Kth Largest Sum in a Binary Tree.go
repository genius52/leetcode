package tree

import (
	"container/list"
	"sort"
)

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var q list.List
	q.PushFront(root)
	var data []int64
	for q.Len() > 0 {
		var cur_len int = q.Len()
		var sum int64 = 0
		for i := 0; i < cur_len; i++ {
			var cur *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			sum += int64(cur.Val)
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
		data = append(data, sum)
	}
	if len(data) < k {
		return -1
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	return data[k-1]
}
