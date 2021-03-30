package list_queue


//Definition for a Node.
type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}

func dfs_flatten(node *Node)(*Node,*Node){
	if node == nil{
		return nil,nil
	}
	var visit *Node = node
	var tail *Node = node
	for visit != nil{
		if visit.Child != nil{
			next := visit.Next
			sub_head,sub_tail := dfs_flatten(visit.Child)
			sub_head.Prev = visit
			sub_tail.Next = next
			if next != nil{
				next.Prev = sub_tail
				tail = next
			}else{
				tail = sub_tail
			}
			visit.Next = sub_head
			visit.Child = nil
			visit = next
		}else{
			tail = visit
			visit = visit.Next
		}
	}
	return node,tail
}

func Flatten(root *Node) *Node {
	head,_ := dfs_flatten(root)
	return head
}