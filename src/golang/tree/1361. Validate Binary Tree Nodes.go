package tree

import "container/list"

//Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
//Output: true
//Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
//Output: false
func ValidateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool{
	var visited []bool = make([]bool,n)
	var has_parent []bool = make([]bool,n)
	for i := 0;i < n;i++{
		if leftChild[i] != -1{
			has_parent[leftChild[i]] = true
		}
		if rightChild[i] != -1{
			has_parent[rightChild[i]] = true
		}
	}
	var root int = -1
	var root_cnt int = 0
	for i := 0;i < n;i++{
		if !has_parent[i]{
			root_cnt++
			root = i
		}
	}
	if root_cnt == 0 || root_cnt >= 2{
		return false
	}
	var q list.List
	q.PushBack(root)
	visited[root] = true
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur int = q.Front().Value.(int)
			q.Remove(q.Front())
			if leftChild[cur] != -1{
				if visited[leftChild[cur]]{
					return false
				}
				visited[leftChild[cur]] = true
				q.PushBack(leftChild[cur])
			}
			if rightChild[cur] != -1{
				if visited[rightChild[cur]]{
					return false
				}
				visited[rightChild[cur]] = true
				q.PushBack(rightChild[cur])
			}
		}
	}
	for i := 0;i < n;i++{
		if !visited[i]{
			return false
		}
	}
	return true
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var record map[int]int = make(map[int]int)
	for i := 0;i < n;i++{
		left_node := leftChild[i]
		right_node := rightChild[i]
		if (left_node != -1 && left_node < i) || (right_node != -1 && right_node < i){
			return false
		}
		if _,ok := record[left_node];ok{
			return false
		}
		if _,ok := record[right_node];ok{
			return false
		}
		if left_node != -1{
			record[left_node] = 1
		}
		if right_node != -1{
			record[right_node] = 1
		}
	}
	for i := 1;i < n;i++{
		if _,ok := record[i];!ok{
			return false
		}
	}
	return true
}