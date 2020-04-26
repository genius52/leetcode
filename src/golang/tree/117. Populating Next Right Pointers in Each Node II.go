package tree

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func levelvisit_connect(node *Node,level int,record map[int]*Node){
	if node == nil{
		return
	}
	if _,ok := record[level];ok{
		record[level].Next = node
		record[level] = node
	}else{
		record[level] = node
	}
	levelvisit_connect(node.Left,level + 1,record)
	levelvisit_connect(node.Right,level + 1,record)
}

func connect(root *Node) *Node {
	var record map[int]*Node = make(map[int]*Node)
	levelvisit_connect(root,0,record)
	return root
}
