package tree

import "container/list"

//Input:
//A binary tree as following:
//       4
//     /   \
//    2     6
//   / \   /
//  3   1 5
//
//v = 1
//
//d = 2
//
//Output:
//       4
//      / \
//     1   1
//    /     \
//   2       6
//  / \     /
// 3   1   5
func visit_addOneRow(node *TreeNode,v int,depth int,target_depth int,is_left bool) *TreeNode {
	if node == nil{
		if depth == target_depth{
			var new_node TreeNode
			new_node.Val = v
			return &new_node
		}
		return nil
	}
	if depth == target_depth{
		pre := node
		var new_node TreeNode
		new_node.Val = v
		if is_left{
			new_node.Left = pre
		}else{
			new_node.Right = pre
		}
		return &new_node
	}else{
		node.Left = visit_addOneRow(node.Left,v,depth + 1,target_depth,true)
		node.Right = visit_addOneRow(node.Right,v,depth + 1,target_depth,false)
		return node
	}
}

func AddOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1{
		var new_node TreeNode
		new_node.Val = v
		new_node.Left = root
		return &new_node
	}
	root.Left = visit_addOneRow(root,v,2,d,true)
	root.Right = visit_addOneRow(root,v,2,d,false)
	return root
}

func addOneRow2(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil{
		return nil
	}
	if depth == 1{
		var new_root TreeNode
		new_root.Val = val
		new_root.Left = root
		return &new_root
	}
	var q list.List
	q.PushBack(root)
	for depth > 1{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var node *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if depth == 2{
				pre_left := node.Left
				pre_right := node.Right
				var new_node1 TreeNode
				new_node1.Val = val
				node.Left = &new_node1
				if pre_left != nil{
					new_node1.Left = pre_left
				}
				var new_node2 TreeNode
				new_node2.Val = val
				node.Right = &new_node2
				if pre_right != nil{
					new_node2.Right = pre_right
				}
			}else{
				if node.Left != nil{
					q.PushBack(node.Left)
				}
				if node.Right != nil{
					q.PushBack(node.Right)
				}
			}
		}
		depth--
	}
	return root
}