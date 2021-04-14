package tree

import (
	"container/list"
	"fmt"
)

//513
func level_visit(node *TreeNode,level_num int,record *map[int][]int){
	if nil == node{
		return
	}
	if _,ok := (*record)[level_num];ok{
		(*record)[level_num] = append((*record)[level_num],node.Val)
		fmt.Println((*record)[level_num])
	}else{
		(*record)[level_num] = []int{node.Val}
		fmt.Println("clean map")
		fmt.Println(len((*record)))
	}
	level_visit(node.Left,level_num + 1,record)
	level_visit(node.Right,level_num + 1,record)
}

func findBottomLeftValue(root *TreeNode) int {
	if nil == root{
		return 0
	}
	var record map[int][]int = make(map[int][]int)
	level_visit(root,0,&record)

	var max_level = -1
	for level_num,_ := range record{
		if level_num > max_level{
			max_level = level_num
		}
	}
	return record[max_level][0]
}

func findBottomLeftValue2(root *TreeNode) int {
	if nil == root{
		return 0
	}
	var q list.List
	q.PushBack(root)
	var last_left int = root.Val
	for q.Len() > 0{
		var l int = q.Len()
		var first_node bool = true
		for i := 0;i < l;i++{
			node := q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if first_node{
				last_left = node.Val
				first_node = false
			}
			if node.Left != nil{
				q.PushBack(node.Left)
			}
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
	}
	return last_left
}