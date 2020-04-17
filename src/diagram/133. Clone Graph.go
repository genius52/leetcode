package diagram

import "container/list"

type Node struct {
	Val int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
	var head *Node = nil
	var record map[int]*Node = make(map[int]*Node)
	var q list.List
	q.PushBack(node)
	for q.Len() > 0{
		top := q.Front()
		var copy_node *Node
		if _,ok := record[top.Value.(*Node).Val];ok {
			copy_node = record[top.Value.(*Node).Val]
		}else{
			copy_node = new(Node)
			copy_node.Val = top.Value.(*Node).Val
			q.Remove(top)
			if head == nil{
				head = copy_node
				record[copy_node.Val] = copy_node
			}
		}
		q.Remove(top)
		for _, n := range top.Value.(*Node).Neighbors{
			if _,ok := record[n.Val];ok{
				copy_node.Neighbors = append(copy_node.Neighbors,record[n.Val])
				continue
			}
			var sub_node *Node = new(Node)
			sub_node.Val = n.Val
			record[sub_node.Val] = sub_node
			copy_node.Neighbors = append(copy_node.Neighbors,sub_node)
			q.PushBack(n)
		}
	}
	return head
}