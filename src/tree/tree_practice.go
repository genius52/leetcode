package tree

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
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


func pruneTree(root *TreeNode) *TreeNode {
	if(nil == root){
		return root;
	}
	if(nil != root.Left){
		root.Left = pruneTree(root.Left)
	}
	if(nil != root.Right){
		root.Right = pruneTree(root.Right)
	}
	if(0 == root.Val){
		if(nil == root.Left && nil == root.Right){
			return nil
		}
	}
	return root
}

func isSametree(t1 *TreeNode, t2 *TreeNode) bool{
	if nil == t1 && nil == t2{
		return true
	}
	if nil == t1 || nil == t2{
		return false
	}
	if t1.Val != t2.Val{
		return false
	}
	return isSametree(t1.Left,t2.Left) && isSametree(t1.Right,t2.Right);
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if nil == t{
		return true
	}
	if nil == s{
		return false
	}
	if  isSametree(s,t){
		return true
	}
	return isSametree(s.Left,t) || isSametree(s.Right,t);
}

func find_path(node *TreeNode,sum int) int{
	if nil == node {
		return 0
	}
	var res int = 0
	if node.Val == sum{
		res += 1
	}
	return res + find_path(node.Left,sum - node.Val) + find_path(node.Right,sum - node.Val)
}

func pathSum(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}
	return find_path(root,sum) + pathSum(root.Left,sum) + pathSum(root.Right,sum)
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


func findMode(root *TreeNode) []int {
	var res []int
	if nil == root{
		return res
	}
	var cnt_map map[int]int = make(map[int]int)
	depth_visit(root,cnt_map)
	max_occurs := 0
	for _,v := range cnt_map{
		if v > max_occurs{
			max_occurs = v
		}
	}
	for k,v := range cnt_map{
		if v == max_occurs{
			res = append(res, k)
		}
	}
	return res
}

var cur_val int = 0
var tmp_cnt int = 0
var max_cnt int = 0

func inorder(node *TreeNode,res *[]int){
	if nil == node{
		return
	}
	inorder(node.Left,res)
	if node.Val != cur_val{
		cur_val = node.Val
		tmp_cnt = 1
	}else{
		tmp_cnt++
	}
	if tmp_cnt > max_cnt{
		*res = (*res)[0:0]
		*res = append(*res,cur_val)
		max_cnt = tmp_cnt
	} else if tmp_cnt == max_cnt{
		*res = append(*res,cur_val)
	}
	inorder(node.Right,res)
}

func findMode2(root *TreeNode) []int{
	var res []int
	if nil == root{
		return res
	}
	inorder(root,&res)
	return res
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

//513
func level_visit(node *TreeNode,level_num int,record *map[int][]int){
	if nil == node{
		return
	}
	if _,ok := (*record)[level_num];ok{
		(*record)[level_num] = append((*record)[level_num],node.Val)
		fmt.Println((*record)[level_num])
	}else{
		(*record)[level_num] = []int{node.Val}
		fmt.Println("clean map")
		fmt.Println(len((*record)))
	}
	level_visit(node.Left,level_num + 1,record)
	level_visit(node.Right,level_num + 1,record)
}

func findBottomLeftValue(root *TreeNode) int {
	if nil == root{
		return 0
	}
	var record map[int][]int = make(map[int][]int)
	level_visit(root,0,&record)

	var max_level = -1
	for level_num,_ := range record{
		if level_num > max_level{
			max_level = level_num
		}
	}
	fmt.Println(len(record))
	return record[max_level][0]
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

//103
//[1,2,3,4,null,null,5]
//[3,9,20,null,null,15,7]
func level_visit103(node *TreeNode,level int,res *[][]int){
	if nil == node{
		return
	}
	l := len(*res)
	if l <= level{
		*res = append(*res,[]int{node.Val})
	}else{
		(*res)[level] = append((*res)[level], node.Val)
	}
	//if level % 2 == 1{
		level_visit103(node.Left,level + 1,res)
		level_visit103(node.Right,level + 1,res)
	//}else{
	//	level_visit103(node.Right,level + 1,res)
	//	level_visit103(node.Left,level + 1,res)
	//}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	level_visit103(root,0,&res)
	for i := 0; i < len(res);i++{
		if i % 2 == 1{
			for begin, end := 0, len(res[i])-1; begin < end; begin, end = begin+1, end-1 {
				res[i][begin], res[i][end] = res[i][end], res[i][begin]
			}
		}
	}
	return res
}

//337
func dp_rob(node *TreeNode,is_rob bool,record map[*TreeNode]int)(ret int){
	if nil == node{
		return 0
	}
	var max int = 0
	if is_rob{
		max = dp_rob(node.Left,false,record) + dp_rob(node.Right,false,record)
	}else{
		max = max_int(node.Val + dp_rob(node.Left,true,record) + dp_rob(node.Right,true,record),dp_rob(node.Left,false,record) + dp_rob(node.Right,false,record))
	}
	return max
}

func rob(root *TreeNode) int {
	if nil == root{
		return 0
	}
	var record map[*TreeNode]int = make(map[*TreeNode]int)
	return dp_rob(root,false,record)
}

//96
//Input: 3
//Output: 5
//Explanation:
//Given n = 3, there are a total of 5 unique BST's:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
func numTrees(n int) int {
	var dp []int = make([]int,n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2;i <= n;i++{
		for j := 0;j < i;j++{
			dp[i] += dp[i - j - 1] * dp[j]
		}
	}
	return dp[n]
}

//105
//preorder = [3,9,20,15,7]
//inorder = [9,3,15,20,7]
//Return the following binary tree:
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
func recursive_buildTree(preorder []int,pre_start int,pre_end int,inorder []int,in_start int,in_end int) *TreeNode{
	if pre_start > pre_end || in_start> in_end {
		return nil
	}

	var node *TreeNode = new(TreeNode)
	node.Val = preorder[pre_start]
	find_inorder_index:= in_start
	var find bool = false
	for ;find_inorder_index <= in_end;find_inorder_index++{
		if preorder[pre_start] == inorder[find_inorder_index]{
			find = true
			break
		}
	}
	if find{
		node.Left = recursive_buildTree(preorder,pre_start + 1,pre_end,inorder,in_start,find_inorder_index-1)
		node.Right = recursive_buildTree(preorder,pre_start + 1 - in_start + find_inorder_index ,pre_end,inorder,find_inorder_index+1,in_end)
	}
	return node
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}
	return recursive_buildTree(preorder,0,len(preorder) - 1,inorder,0,len(inorder) - 1)
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

//109
func divide_sortedListToBST(lnode *ListNode,begin int,end int)*TreeNode{
	if begin >= end{
		return nil
	}
	var mid int = begin + (end - begin)/2
	var i int = 0
	var visit *ListNode = lnode
	for i != mid{
		visit = visit.Next
		i++
	}
	var tnode TreeNode
	tnode.Val = visit.Val
	tnode.Left = divide_sortedListToBST(lnode,begin,mid)
	tnode.Right = divide_sortedListToBST(lnode,mid+1,end)
	return &tnode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var l int = 0
	var visit *ListNode = head
	for visit != nil{
		visit = visit.Next
		l++
	}
	return divide_sortedListToBST(head,0,l)
}

//106
//inorder = [9,3,15,20,7]
//postorder = [9,15,7,20,3]
//Return the following binary tree:
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
func inpost_buildTree(inorder []int,in_begin int,in_end int,postorder []int,root_pos int)*TreeNode{
	if root_pos < 0 || in_begin > in_end{
		return nil
	}
	//find root's position in inorder
	var i int = in_begin
	for i < in_end && inorder[i] != postorder[root_pos]{
		i++
	}
	//length of right tree nodes
	var right_tree_length int = in_end - i + 1
	var node TreeNode
	node.Val = postorder[root_pos]
	node.Left = inpost_buildTree(inorder,in_begin,i - 1,postorder,root_pos - right_tree_length)
	node.Right = inpost_buildTree(inorder,i + 1, in_end,postorder,root_pos - 1)
	return &node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	lin := len(inorder)
	lpost := len(postorder)
	if lin == 0 || lpost == 0 || lin != lpost{
		return nil
	}
	return inpost_buildTree(inorder,0,lin - 1,postorder,lpost - 1)
}

//114
//For example, given the following tree:
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//The flattened tree should look like:
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
func pre_faltten(node *TreeNode)*TreeNode{
	if nil == node{
		return nil
	}
	pre_right := node.Right
	left_tail := pre_faltten(node.Left)
	if nil != left_tail{
		left_tail.Right = pre_right
	}
	if nil != node.Left{
		node.Right = node.Left
	}
	node.Left = nil
	right_tail := pre_faltten(pre_right)
	if nil != right_tail{
		return right_tail
	}
	if nil != left_tail{
		return left_tail
	}
	return node
}

func flatten(root *TreeNode)  {
	pre_faltten(root)
}

//208


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
type Trie struct {
	data map[string]bool
}


/** Initialize your data structure here. */
func Constructor4() Trie {
	var obj Trie
	obj.data = make(map[string]bool)
	return obj
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	this.data[word] = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if _,ok := this.data[word];ok{
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for word,_ := range this.data{
		if strings.HasPrefix(word,prefix){
			return true
		}
	}
	return false
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

//124
//return1: open way max
//return2: closed way max
func post_visit(node *TreeNode)(int,int){
	if node == nil{
		return math.MinInt32,math.MinInt32
	}
	if node.Left == nil && node.Right == nil{
		return node.Val,node.Val
	}
	left_open_sum,left_close_sum := post_visit(node.Left)
	right_open_sum,right_close_sum := post_visit(node.Right)
	return max_int_number(left_open_sum + node.Val,right_open_sum + node.Val,node.Val),
	max_int_number(left_open_sum,left_close_sum,right_open_sum,right_close_sum,left_open_sum + right_open_sum + node.Val)
}

func maxPathSum(root *TreeNode) int {
	r1,r2 := post_visit(root)
	return max_int_number(r1,r2)
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

//889
//Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//Output: [1,2,3,4,5,6,7]
func build_constructFromPrePost(pre []int,pre_start int,pre_end int,post []int,post_start int,post_end int)*TreeNode{
	if pre_end >= len(pre) || post_end >= len(post) || pre_start > pre_end || post_start > post_end{
		return nil
	}
	var node TreeNode
	node.Val = pre[pre_start]
	if pre_start == pre_end && post_start == post_end{
		return &node
	}
	var left_end_in_post int = post_end
	for post[left_end_in_post] != pre[pre_start + 1] && left_end_in_post > post_start{
		left_end_in_post--
	}
	left_end_in_pre := pre_start + left_end_in_post - post_start + 1
	node.Left = build_constructFromPrePost(pre,pre_start + 1, left_end_in_pre,post,post_start,left_end_in_post)
	node.Right = build_constructFromPrePost(pre,left_end_in_pre + 1,pre_end,post,left_end_in_post + 1,post_end - 1)
	return &node
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	l := len(pre)
	if l == 0{
		return nil
	}
	return build_constructFromPrePost(pre,0,l - 1,post,0,l - 1)
}


//889
//var pre_cnt int = 0
//var post_cnt int = 0
//func constructFromPrePost(pre []int, post []int) *TreeNode {
//	var node *TreeNode = new(TreeNode)
//	node.Val = pre[pre_cnt]
//	pre_cnt++
//	if node.Val != post[post_cnt] {
//		node.Left = constructFromPrePost(pre,post)
//	}
//	if node.Val != post[post_cnt] {
//		node.Right = constructFromPrePost(pre,post)
//	}
//	post_cnt++
//	return node
//}


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


//865
func level_visit_865(node *TreeNode,level int)(max_depth int,res *TreeNode){
	if nil == node{
		return level,node
	}
	level++
	left_max_depth,leftnode := level_visit_865(node.Left,level)
	right_max_depth,rightnode := level_visit_865(node.Right,level)
	if left_max_depth > right_max_depth {
		if left_max_depth >level {
			return left_max_depth,leftnode
		}
		return level,node
	}else if left_max_depth < right_max_depth{
		if right_max_depth > level{
			return right_max_depth,rightnode
		}
		return level,node
	}else{
		return left_max_depth,node
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if nil == root{
		return root
	}
	_,res :=  level_visit_865(root,0)
	return res
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


//515
func preorder_search_max(node *TreeNode,record *[]int,level int){
	if nil == node{
		return
	}
	if level == len(*record){
		*record = append(*record, node.Val)
	}else{
		if node.Val > (*record)[level]{
			(*record)[level] = node.Val
		}
	}
	preorder_search_max(node.Left,record,level + 1)
	preorder_search_max(node.Right,record,level + 1)
}

func largestValues(root *TreeNode) []int {
	var res []int
	if nil == root{
		return res
	}
	preorder_search_max(root,&res,0)
	return res
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
			fmt.Println(node.Val)
			q.Remove(top)
			if node.Right != nil{
				q.PushBack(node.Right)
			}
		}
	}
}