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
	var head *Node = new(Node)
	head.Val = node.Val
	var record map[*Node]*Node = make(map[*Node]*Node)
	record[node] = head
	var q list.List
	q.PushBack(node)
	for q.Len() > 0{
		top := q.Front()
		copy_node := record[top.Value.(*Node)]
		q.Remove(top)
		for _, n := range top.Value.(*Node).Neighbors{
			if _,ok := record[n];ok{
				copy_node.Neighbors = append(copy_node.Neighbors,record[n])
				continue
			}
			var sub_node *Node = new(Node)
			sub_node.Val = n.Val
			record[n] = sub_node
			copy_node.Neighbors = append(copy_node.Neighbors,sub_node)
			q.PushBack(n)
		}
	}
	return head
}

//dfs
func dfs_cloneGraph(node *Node,record map[*Node]*Node)*Node{
	if node == nil{
		return nil
	}
	if _,ok := record[node];ok{
		return record[node]
	}
	var copy *Node = new(Node)
	copy.Val = node.Val
	record[node] = copy
	for _,n := range node.Neighbors{
		copy.Neighbors = append(copy.Neighbors,dfs_cloneGraph(n,record))
	}

	return copy
}

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
	var record map[*Node]*Node = make(map[*Node]*Node)
	return dfs_cloneGraph(node,record)
}