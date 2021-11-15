package tree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//recrusive solution
func previsit_isEvenOddTree(node *TreeNode,level int,even_queue map[int][]int,odd_queue map[int][]int)bool{
	if node == nil{
		return true
	}
	if level % 2 == 1{ // index 1,3,5..
		if node.Val % 2 == 1{//should be 2,4,6
			return false
		}
		if _,ok := even_queue[level];ok{
			var even_len int = len(even_queue[level])
			if even_len > 0{
				if even_queue[level][even_len - 1] <= node.Val{
					return false
				}
			}
		}
		even_queue[level] = append(even_queue[level],node.Val)
	}else{
		if node.Val % 2 == 0{
			return false
		}
		if _,ok := odd_queue[level];ok{
			var odd_len = len(odd_queue[level])
			if odd_len > 0{
				if odd_queue[level][odd_len - 1] >= node.Val{
					return false
				}
			}
		}
		odd_queue[level] = append(odd_queue[level],node.Val)
	}
	return previsit_isEvenOddTree(node.Left,level + 1,even_queue,odd_queue) &&
		previsit_isEvenOddTree(node.Right,level + 1,even_queue,odd_queue)
}

func IsEvenOddTree(root *TreeNode) bool {
	if root == nil{
		return true
	}
	var even_queue map[int][]int = make(map[int][]int)
	var odd_queue map[int][]int = make(map[int][]int)
	return previsit_isEvenOddTree(root,0,even_queue,odd_queue)
}

//偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
//奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
//bfs solution
func isEvenOddTree(root *TreeNode) bool{
	var q list.List
	even_level := true
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var pre *TreeNode = nil
		for i := 0;i < l;i++{
			cur := q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if cur.Left != nil{
				q.PushBack(cur.Left)
			}
			if cur.Right != nil{
				q.PushBack(cur.Right)
			}
			if even_level{
				if (cur.Val | 1) != cur.Val{
					return false
				}
			}else{
				if (cur.Val | 1) == cur.Val{
					return false
				}
			}
			if pre == nil{
				pre = cur
			}else{
				if even_level{
					if pre.Val >= cur.Val{
						return false
					}
				}else{
					if pre.Val <= cur.Val{
						return false
					}
				}
				pre = cur
			}
		}
		even_level = !even_level
	}
	return true
}