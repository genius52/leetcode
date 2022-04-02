package tree

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type tree_node struct{
	value int
	left *tree_node
	right *tree_node
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func max_int_number(nums ...int)int{
	var max int = math.MinInt32
	for _,n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}

func min_int_number(nums ...int)int{
	var min int = math.MaxInt32
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

var prev *tree_node = nil

func inorder_tree(node *tree_node){
	if node == nil {
		return
	}
	inorder_tree(node.left)

	if prev != nil{
		node.left = prev
		prev.right = node
	}else{
		node.left = nil
	}

	prev = node
	inorder_tree(node.right)
}

func depth_visit(node *TreeNode,cnt_map map[int]int){
	if nil == node{
		return
	}
	if _,ok := cnt_map[node.Val];ok{
		cnt_map[node.Val]++
	}else{
		cnt_map[node.Val] = 1
	}
	depth_visit(node.Left,cnt_map)
	depth_visit(node.Right,cnt_map)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var record map[int]int = make(map[int]int)
	for index,v := range nums{
		if pos,ok := record[v];ok{
			if index - pos <= k{
				return true
			}else{
				record[v] = index
			}
		}else{
			record[v] = index
		}
	}
	return false
}

//754. Reach a Number
func reachNumber(target int) int {
	var res int = 0
	if target < 0 {
		return reachNumber(-target)
	}

	return res
}

func previsit_balanceBST(node *TreeNode,record *[]*TreeNode){
	if node == nil{
		return
	}
	previsit_balanceBST(node.Left,record)
	*record = append(*record, node)
	previsit_balanceBST(node.Right,record)
}

func build_bst(record []*TreeNode,low int,high int)*TreeNode{
	if low > high{
		return nil
	}
	mid := (low + high)/2
	node := record[mid]
	if low == high{
		return node
	}
	node.Left = build_bst(record,low ,mid - 1)
	node.Right = build_bst(record,mid + 1,high)
	return node
}

func balanceBST(root *TreeNode) *TreeNode {
	var record []*TreeNode
	previsit_balanceBST(root,&record)
	for i := 0;i < len(record);i++{
		record[i].Right = nil
		record[i].Left = nil
	}
	return build_bst(record,0,len(record) - 1)
}

//1028
//Input: "1-2--3--4-5--6--7"
//Output: [1,2,5,3,4,6,7]
//Input: "1-2--3---4-5--6---7"
//Output: [1,2,5,3,null,6,null,4,null,7]
func recoverFromPreorder(S string) *TreeNode {
	l := len(S)
	if l == 0{
		return nil
	}
	var node TreeNode
	begin := 0
	for begin < l{//skip current level tag
		if S[begin] != '-'{
			break
		}
		begin++
	}

	begin2 := begin
	for begin2 < l{//find root value and
		if S[begin2] == '-'{
			val,err := strconv.ParseInt(S[begin:begin2],10,64)
			if err == nil{
				node.Val = int(val)
			}
			break
		}
		begin2++
	}
	if begin2 == l{
		val,err := strconv.ParseInt(S[begin:begin2],10,64)
		if err == nil{
			node.Val = int(val)
		}
	}
	var prefix_child string
	begin3 := begin2
	for begin3 < l{//find left child prefix
		if S[begin3] != '-'{
			prefix_child += S[begin2:begin3]
			break
		}
		begin3++
	}
	begin4 := begin3
	prefix_len := len(prefix_child)
	var find_right bool = false
	for begin4 < l - prefix_len{
		if S[begin4:begin4 + prefix_len] == prefix_child && S[begin4 - 1] != '-' && S[begin4 + prefix_len] != '-'{
			find_right = true
			break
		}
		begin4++
	}
	if begin4 > begin2{
		var left_s string
		if find_right{
			left_s = S[begin2:begin4]
		}else{
			left_s = S[begin2:]
		}
		node.Left = recoverFromPreorder(left_s)
	}
	if find_right{
		right_s := S[begin4:]
		node.Right = recoverFromPreorder(right_s)
	}
	return &node
}


//no recursive pre_visit
func Preorder_visit_norecursive(root *TreeNode){
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		top := q.Back()
		node := top.Value.(*TreeNode)
		fmt.Println(node.Val)
		q.Remove(top)
		if node.Right != nil{
			q.PushBack(node.Right)
		}
		if node.Left != nil{
			q.PushBack(node.Left)
		}
	}
}

//no recursive inorder visit
func Inorder_visit_norecursive(root *TreeNode){
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		for root.Left != nil{
			q.PushBack(root.Left)
			root = root.Left
		}
		if q.Len() > 0{
			top := q.Back()
			node := top.Value.(*TreeNode)
			fmt.Println(node.Val)//visit here
			q.Remove(top)
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
	}
}