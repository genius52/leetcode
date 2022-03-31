package tree

import "container/list"

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

//bfs
func Connect(root *Node) *Node{
	if root == nil{
		return root
	}
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var pre *Node = nil
		for i := 0;i < l;i++{
			cur := q.Front().Value.(*Node)
			q.Remove(q.Front())
			if pre == nil{
				pre = cur
			}else{
				pre.Next = cur
				pre = cur
			}
			if cur.Left != nil{
				q.PushBack(cur.Left)
			}
			if cur.Right != nil{
				q.PushBack(cur.Right)
			}
		}
	}
	return root
}

//func levelvisit_connect(node *Node,level int,record map[int]*Node){
//	if node == nil{
//		return
//	}
//	if _,ok := record[level];ok{
//		record[level].Next = node
//		record[level] = node
//	}else{
//		record[level] = node
//	}
//	levelvisit_connect(node.Left,level + 1,record)
//	levelvisit_connect(node.Right,level + 1,record)
//}
//
//func connect(root *Node) *Node {
//	var record map[int]*Node = make(map[int]*Node)
//	levelvisit_connect(root,0,record)
//	return root
//}
