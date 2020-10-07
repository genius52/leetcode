package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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