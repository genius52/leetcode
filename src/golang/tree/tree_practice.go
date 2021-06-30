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

func inorder_visit2(root *TreeNode){
	if root == nil{
		return
	}
	node := root
	for node != nil{

	}
}

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

//914. X of a Kind in a Deck of Cards
func hasGroupsSizeX(deck []int) bool {
	var record map[int]int = make(map[int]int)
	for _,num := range deck{
		if _,ok := record[num];ok{
			record[num]++
		}else{
			record[num] = 1
		}
	}

	var min_cnt int = math.MaxInt32
	var max_cnt int = 0
	for _,v := range record{
		if v > max_cnt{
			max_cnt = v
		}
		if v < min_cnt{
			min_cnt = v
		}
	}
	if min_cnt < 2{
		return false
	}
	min_cnt = 2
	for ;min_cnt <= max_cnt;min_cnt++{
		var match bool = true
		for _,v := range record{
			if v % min_cnt != 0{
				match = false
				break
			}
		}
		if match{
			return true
		}
	}
	return false
}

//979
func post_visit_tree(node *TreeNode,steps *int) int{
	if nil == node{
		return 0
	}

	left_val := post_visit_tree(node.Left,steps)
	right_val := post_visit_tree(node.Right,steps)
	*steps += int(math.Abs(float64(left_val + right_val + node.Val - 1)))
	return left_val + right_val + node.Val - 1
}

func distributeCoins(root *TreeNode) int {
	var steps int = 0
	post_visit_tree(root,&steps)
	return steps
}

//1110
func inorder_visit(node **TreeNode,init bool,to_delete []int,res *[]*TreeNode){
	var should_delete bool = false
	if nil == *node{
		return
	}
	for _,val := range to_delete{
		if val == (*node).Val{
			should_delete = true
			break
		}
	}
	temp := *node
	if should_delete{
		*node = nil
		init = true
	}else{
		if init{
			*res = append(*res, *node)
			init = false
		}
	}
	inorder_visit(&(temp.Left),init,to_delete,res)
	inorder_visit(&(temp.Right),init,to_delete,res)
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var res []*TreeNode
	inorder_visit(&root,true,to_delete,&res)
	return res
}


//1130

func dp_arr(arr []int,max_records [][]int,memo [][]int,left,right int)int{
	if left >= right{
		return 0
	}
	if memo[left][right] != 0{
		return memo[left][right]
	}
	var res int = 1<<30
	for i := left;i < right;i++{
		res = int(math.Min(float64(res),float64(max_records[left][i] * max_records[i+1][right] + dp_arr(arr,max_records,memo,left,i) + dp_arr(arr,max_records,memo,i+1,right))))
	}
	memo[left][right] = res
	return res
}

func mctFromLeafValues(arr []int) int {
	var memo [][]int = make([][]int,len(arr))
	var max_records [][]int = make([][]int,len(arr))
	for i := 0;i < len(arr);i++{
		max_records[i] = make([]int,len(arr))
		memo[i] = make([]int,len(arr))
		max_records[i][i] = arr[i]
		for j := i+1;j < len(arr);j++{
			max_records[i][j] = int(math.Max(float64(arr[j]),float64(max_records[i][j-1])))
		}
	}
	return dp_arr(arr,max_records,memo,0,len(arr)-1)
}


//1261
type FindElements struct {
	root *TreeNode
	set map[int]bool
}

func dfs_tree(node *TreeNode,val int,ele FindElements){
	if nil == node{
		return
	}
	node.Val = val
	ele.set[node.Val] = true
	dfs_tree(node.Left,node.Val*2+1,ele)
	dfs_tree(node.Right,node.Val*2+2,ele)
}

func Constructor2(root *TreeNode) FindElements {
	root.Val = 0
	ele := FindElements{root,make(map[int]bool)}
	ele.set[root.Val] = true
	dfs_tree(root.Left,root.Val*2+1,ele)
	dfs_tree(root.Right,root.Val*2+2,ele)
	return ele
}

func (this *FindElements) Find(target int) bool {
	if _,ok := this.set[target];ok{
		return true
	}
	return false
}

//1315
func dfs_sumEvenGrandparent(node *TreeNode,parent_val int,grand_val int)int{
	if nil == node{
		return 0
	}
	var cur_sum int = 0
	if grand_val % 2 == 0 {
		cur_sum = node.Val
	}
	return cur_sum + dfs_sumEvenGrandparent(node.Left,node.Val,parent_val) + dfs_sumEvenGrandparent(node.Right,node.Val,parent_val)
}

func sumEvenGrandparent(root *TreeNode) int {
	if nil == root{
		return 0
	}
	return dfs_sumEvenGrandparent(root.Left,root.Val,1) + dfs_sumEvenGrandparent(root.Right,root.Val,1)
}

//1325
func postvisit_removeLeafNodes(node *TreeNode,target int)bool{
	if nil == node{
		return true
	}
	left_res :=  postvisit_removeLeafNodes(node.Left,target)
	right_res := postvisit_removeLeafNodes(node.Right,target)
	if left_res{
		node.Left = nil
	}
	if right_res{
		node.Right = nil
	}
	if left_res && right_res{
		if node.Val == target{
			return true
		}
	}
	return false
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if postvisit_removeLeafNodes(root,target){
		return nil
	}else{
		return root
	}
}

//1361
//Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
//Output: true
//Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
//Output: false
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

//1305
//Input: root1 = [2,1,4], root2 = [1,0,3]
//Output: [0,1,1,2,3,4]
func pre_visit1305(node *TreeNode,record *[]int){
	if node == nil {
		return
	}
	pre_visit1305(node.Left,record)
	*record = append(*record, node.Val)
	pre_visit1305(node.Right,record)
}

func merge_array(record1 []int,record2 []int)[]int{
	l1 := len(record1)
	l2 := len(record2)
	if l1 == 0{
		return record2
	}
	if l2 == 0{
		return record1
	}
	var res []int = make([]int,l1 + l2)
	i := 0
	h1 := 0
	h2 := 0
	for h1 < l1 && h2 < l2{
		if record1[h1] < record2[h2]{
			res[i] = record1[h1]
			h1++
		}else{
			res[i] = record2[h2]
			h2++
		}
		i++
	}
	if h1 == l1{
		for h2 < l2{
			res[i] = record2[h2]
			i++
			h2++
		}
	}else if h2 == l2{
		for h1 < l1{
			res[i] = record1[h1]
			i++
			h1++
		}
	}
	return res
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var record1 []int
	pre_visit1305(root1,&record1)
	var record2 []int
	pre_visit1305(root2,&record2)
	return merge_array(record1,record2)
}

//894
func allPossibleFBT(N int)[]*TreeNode{
	var res []*TreeNode
	if N % 2 == 0{
		return res
	}
	if N == 1{
		var node TreeNode
		node.Val = 0
		res = append(res,&node)
		return res
	}
	for left := 1; left < N - 1;left += 2{
		left_nodes := allPossibleFBT(left)
		right_nodes := allPossibleFBT(N - left - 1)
		for _,ln := range left_nodes{
			for _,rn := range right_nodes{
				var node TreeNode
				node.Val = 0
				node.Left = ln
				node.Right = rn
				res = append(res, &node)
			}
		}
	}
	return res
}

//func allPossibleFBT(N int) []*TreeNode {
//	var record []*TreeNode
//	if N % 2 == 0{
//		return record
//	}
//	return divide(N)
//}

//
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

//1123
func dp_lcaDeepestLeaves(node *TreeNode,level int)(int,*TreeNode){
	if node == nil{
		return level - 1,nil
	}
	if node.Left == nil && node.Right == nil{
		return level,node
	}
	left_depth,left := dp_lcaDeepestLeaves(node.Left,level + 1)
	right_depth,right := dp_lcaDeepestLeaves(node.Right,level + 1)
	if left_depth == right_depth{
		return left_depth,node
	}else if left_depth > right_depth{
		return left_depth,left
	}else{
		return right_depth,right
	}
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_,ret := dp_lcaDeepestLeaves(root,0)
	return ret
}

//1026
func pre_visit(node *TreeNode,min int,max int,max_diff int)int{
	if nil == node{
		return 0
	}
	diff := int(math.Max(math.Abs(float64(node.Val - min)),math.Abs(float64(node.Val - max))))
	max_diff = int(math.Max(float64(max_diff),float64(diff)))
	min = int(math.Min(float64(node.Val),float64(min)))
	max = int(math.Max(float64(node.Val),float64(max)))
	res := int(math.Max(float64(pre_visit(node.Left,min,max,max_diff)),float64(pre_visit(node.Right,min,max,max_diff))))
	res = int(math.Max(float64(res),float64(max_diff)))
	return res
}

func maxAncestorDiff(root *TreeNode) int {
	if nil == root{
		return 0
	}
	return pre_visit(root,root.Val,root.Val,0)
}

//951
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if nil == root1 && nil == root2{
		return true
	}
	if (nil == root1 || nil == root2) || (root1.Val != root2.Val){
		return false
	}
	return (flipEquiv(root1.Left,root2.Left) && flipEquiv(root1.Right,root2.Right)) || (flipEquiv(root1.Left,root2.Right) && flipEquiv(root1.Right,root2.Left))
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