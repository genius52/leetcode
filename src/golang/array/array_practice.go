package array

import (
	"bytes"
	"container/heap"
	"container/list"
	"math"
	"sort"
	"strconv"
	"strings"
)

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func min_int(a,b int)int{
	if a < b {
		return a
	}else{
		return b
	}
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

func max_int_number(nums ...int)int{
	var max int = math.MinInt32
	for _,n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}
//Definition for singly-linked list_queue.
type ListNode struct {
	Val int
	Next *ListNode
}

func remove_duplicated_sorted_array(arr []int)(length int){
	len := len(arr)
	if len <= 1{
		return
	}
	cur_pos := 0
	for i := 1 ;i<len;i++{
		if arr[cur_pos] != arr[i]{
			cur_pos++
			arr[cur_pos] = arr[i]
		}
	}
	return cur_pos + 1
}

func maxArea(height []int) int {
	var length int = len(height)
	var low int = 0
	var high int = length - 1
	var max_cap int = 0
	for low < high  {
		var cap int = int(math.Abs(float64(high - low)) * math.Min(float64(height[high]),float64(height[low])))
		if cap > max_cap{
			max_cap = cap
		}
		if (height[high - 1] >= height[low+1]){
			high--
		} else{
			low++
		}
	}
	return max_cap
}

func lemonadeChange(bills []int) bool {
	var change map[int]int = make(map[int]int)
	for _,pay  := range bills{
		if pay == 5 {
			if _,ok := change[pay];ok{
				change[pay]++
			}else{
				change[pay] = 1
			}
		}else if pay == 10{
			if val,ok := change[5]; !ok || val <= 0{
				return false
			}
			change[5]--
			if _,ok := change[pay];ok{
				change[pay]++
			}else{
				change[pay] = 1
			}
		}else if pay == 20{
			if val,ok := change[5]; !ok || val <= 0{
				return false
			}
			if val,ok := change[10]; ok && val > 0{
				change[10]--
				change[5]--
			}else{
				if change[5] < 3{
					return false
				} else{
					change[5] -= 3
				}
			}
		}
	}
	return true
}

func is_binary_string(s string,len int) bool{
	start := 0
	end := len-1
	mid := len/2
	for pre_val := s[start];start < mid;start++{
		if s[start] != pre_val{
			return false
		}
	}
	for post_val := s[end]; end >= mid;end--{
		if s[end] != post_val{
			return false
		}
	}
	return true
}

func is_point_in_rectangle(px int,py int,rec []int) bool{
	return px > rec[0] && px < rec[2] && py > rec[1] && py < rec[3]
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0]<rec2[2] && rec2[0] < rec1[2] && rec1[3] > rec2[1] && rec1[1] < rec2[3];
}

type KthLargest struct {
	k int
	nums []int
	largest []int
}


func ConstructorK(k int, nums []int) KthLargest {
	var ele KthLargest = KthLargest{
		k:k,
		nums:nums,
	}
	sort.Ints(ele.nums)
	ele.largest = make([]int, k, k)
	copy(ele.largest,ele.nums[0:k])
	return ele
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	this.largest = append(this.largest, val)
	sort.Ints(this.largest)
	this.largest = this.largest[0:this.k]
	return this.nums[this.k - 1]
}

//1255
// Input: words = ["dog","cat","dad","good"],
// letters = ["a","a","c","d","d","d","g","o","o"],
// score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
//Output: 23
//Explanation:
//Score  a=1, c=9, d=5, g=3, o=2
//Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
//Words "dad" and "dog" only get a score of 21.
func maxScoreWords(words []string, letters []byte, score []int) int {
	var res int

	return res
}

//973
//func compare_distance(x1 int,y1 int,x2 int,y2 int) bool{
//	return (x1 * x1 + y1 * y1) - (x2 * x2 + y2 * y2) >= 0
//}
func compare_distance(x1 int,y1 int,x2 int,y2 int) bool{
	res :=  (x1 * x1 + y1 * y1) - (x2 * x2 + y2 * y2)
	return res > 0
}

//最小的k个元素用大顶堆
func init_mink(points [][]int,i int,N int){
	temp_x := points[i][0]
	temp_y := points[i][1]
	for i < N {
		left_child := 2*i + 1
		right_child := 2*i + 2
		if left_child >= N {
			break
		}
		if right_child >= N {
			if compare_distance(temp_x,temp_y,points[left_child][0],points[left_child][1]){
				break
			}else{
				points[i][0] = points[left_child][0]
				points[i][1] = points[left_child][1]
				i = left_child
				break
			}
		}else{
			var max_point_index int
			if compare_distance(points[right_child][0],points[right_child][1],points[left_child][0],points[left_child][1]) {
				max_point_index = right_child
			}else{
				max_point_index = left_child
			}
			if compare_distance(temp_x,temp_y,points[max_point_index][0],points[max_point_index][1]){
				break
			}else{
				points[i][0] = points[max_point_index][0]
				points[i][1] = points[max_point_index][1]
				i = max_point_index
			}
		}
	}
	points[i][0] = temp_x
	points[i][1] = temp_y
}

func kClosest(points [][]int, K int) [][]int {
	if len(points) == K{
		return points
	}
	//堆化前K个数
	for i := K/2 - 1;i >= 0;i-- {
		init_mink(points, i,K)
	}
	//从K+1个数开始，和堆顶元素进行比较。如果元素小于堆顶，则替换堆顶元素，并调整堆
	for i := K;i < len(points);i++{
		if compare_distance(points[i][0],points[i][1],points[0][0],points[0][1]){
			continue
		}
		temp_x := points[i][0]
		temp_y := points[i][1]
		points[i][0] = points[0][0]
		points[i][1] = points[0][1]
		points[0][0] = temp_x
		points[0][1] = temp_y
		init_mink(points,0,K)
	}
	return points[0:K]
}

func kClosest2(points [][]int, K int) [][]int{
	if len(points) == K{
		return points
	}
	for i := 0;i < K;i++{
		for j := len(points) - 1;j > i; j--{
			if compare_distance(points[j-1][0],points[j-1][1],points[j][0],points[j][1]){
				points[j],points[j-1] = points[j-1],points[j]
			}
		}
	}
	return points[0:K]
}

//221
func min_byte_number(nums ...byte)byte{
	var min byte = 127
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

//[2,3,1,1,4]
//bfs solution
func jump2(nums []int) int{
	l := len(nums)
	var cur_fastest_pos int = nums[0]
	var end int = 0
	var steps int = 0
	for i := 0;i < l - 1;i++{
		cur_fastest_pos = max_int(cur_fastest_pos,i + nums[i])
		if cur_fastest_pos >= l - 1{
			steps++
			break
		}
		if i == end{//beyond end edge
			end = cur_fastest_pos
			steps++
		}
	}
	return steps
}

//239
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeap) Pop() interface{}{
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

//[]int{1,3,1,2,0,5}
func maxSlidingWindow(nums []int, k int) []int {
	var l int = len(nums)
	var res []int = make([]int,l - k + 1)
	h := make(MaxHeap,0)
	for i := 0;i < k;i++{
		h.Push(nums[i])
	}
	heap.Init(&h)
	res[0] = h[0]
	for i := 1;i <= (l - k);i++{
		for j := 0;j < k;j++{
			if nums[i-1] == h[j]{
				h = append(h[:j],h[j+1:]...)
				break
			}
		}
		heap.Push(&h,nums[i + k - 1])
		heap.Init(&h)
		res[i] = h[0]
	}
	return res
}

//686
func repeatedStringMatch(A string, B string) int {
	for i,_ := range B{
		if !strings.Contains(A,B[i:i+1]){
			return -1
		}
	}
	la := len(A)
	lb := len(B)
	cur_lena := la
	var buffer bytes.Buffer
	i := 1
	var find bool = false
	for ; i <= lb; i++ {
		buffer.WriteString(A)
		cur_lena += la
		if cur_lena < lb{
			continue
		}
		dupA := buffer.String()
		if strings.Contains(dupA,B){
			find = true
			break
		}
	}
	if !find{
		return -1
	}
	return i
}

//322
func coinChange(coins []int, amount int) int {
	var dp []int = make([]int,amount + 1)
	for i := 0;i < len(dp);i++{
		dp[i] = -1
	}
	dp[0] = 0
	var min_coins int = math.MaxInt32
	for i := 1;i <= amount;i++{
		min_coins = math.MaxInt32
		for _,c := range coins{
			if c > i{
				continue
			}
			if dp[i - c] != - 1{
				min_coins = min_int(1 + dp[i-c],min_coins)
			}
		}
		if min_coins != math.MaxInt32{
			dp[i] = min_coins
		}
	}
	return dp[amount]
}

func dfs_coinChange(coins []int, amount int,record map[int]int)int{
	if amount == 0{
		return 0
	}
	if _,ok := record[amount];ok{
		return record[amount]
	}
	steps := math.MaxInt32
	for _,c := range coins{
		if c > amount{
			continue
		}
		res := dfs_coinChange(coins,amount - c,record)
		if res != -1{
			steps = min_int(res + 1,steps)
		}
	}
	if steps == math.MaxInt32{
		return -1
	}else{
		record[amount] = steps
		return steps
	}
}

func coinChange2(coins []int, amount int)int{
	if amount == 0{
		return 0
	}

	var coin_record map[int]int = make(map[int]int)
	return dfs_coinChange(coins,amount,coin_record)
}

//443
//Input:
//["a","b","b","b","b","b","b","b","b","b","b","b","b"]
//Output:
//Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
func compress(chars []byte) int {
	var end int = len(chars) - 1
	var begin int = end - 1
	var input_pos int = end
	for begin >= 0{
		if chars[begin] == chars[end] && begin != 0{
			begin--
		}else{
			gap := end - begin
			if begin == 0{
				gap++
			}
			if gap >= 2{
				var c byte = chars[end]
				var input []byte
				for gap > 0{
					input = append(input, byte(gap % 10))
					gap = gap / 10
				}
				for _,v := range input{
					chars[input_pos] = v + '0'
					input_pos--
				}
				chars[input_pos] = c
				input_pos--
			}else{
				chars[input_pos] = chars[end]
				input_pos--
			}
			end = begin
			begin--
		}
	}
	copy(chars,chars[input_pos + 1:])
	return len(chars)
}

func dfs_decodeString2(s string,pos *int)string{
	var word string
	var num int = 0
	for ;*pos < len(s);*pos++{
		if s[*pos] == '['{
			*pos++
			var tmp string = dfs_decodeString2(s,pos)
			for j := 0;j < num;j++{
				word += tmp
			}
			num = 0
		}else if s[*pos] == ']'{
			return word
		}else if s[*pos] >= '0' && s[*pos] <= '9'{
			c,_ := strconv.Atoi(string(s[*pos]))
			num = num * 10 + c
		}else{
			word += string(s[*pos])
		}
	}
	return word
}

func decodeString2(s string) string {
	var pos int = 0
	return dfs_decodeString2(s,&pos)
}

//678
func checkValidString(s string) bool {
	var left []int
	var star []int
	for i := 0;i < len(s);i++{
		if s[i] == '('{
			left = append(left, i)
		}else if s[i] == '*'{
			star = append(star,i)
		}else{
			if len(left) == 0 && len(star) == 0{
				return false
			}
			if len(left) > 0{
				left = left[:len(left) - 1]
			}else if len(star) > 0{
				star = star[:len(star) - 1]
			}
		}
	}
	for len(left) > 0 && len(star) > 0{
		l := left[len(left) - 1]
		s := star[len(star) - 1]
		left = left[:len(left) - 1]
		star = star[:len(star) - 1]
		if l > s{
			return false
		}
	}
 	if len(left) > 0{
		return false
	}
	return true
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

//72
func max_rectange(arr [][]int)int{
	rows := len(arr)
	columns := len(arr[0])
	var dp [][]int = make([][]int,rows)
	for i := 0; i < rows;i++{
		dp[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		dp[i][0] = arr[i][0]
	}
	for j := 0;j < columns;j++{
		dp[0][j] = arr[j][0]
	}
	var max int = 0
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if arr[i][j] == 0{
				continue
			}
			if dp[i - 1][j - 1] != 0 && dp[i - 1][j] != 0 && dp[i][j - 1] != 0{
				dp[i][j] = min_int_number(dp[i-1][j-1],dp[i - 1][j],dp[i][j - 1]) + 1
			}else{
				dp[i][j] = 1
			}
			if dp[i][j] > max{
				max = dp[i][j]
			}
		}
	}
	return max
}

//324
//Input: nums = [1, 5, 1, 1, 6, 4]
//Output: One possible answer is [1, 4, 1, 5, 1, 6].
func qsort_wiggleSort(nums []int,low int,high int,mid int){
	if low >= high{
		return
	}

	l := low
	h := high
	tag := nums[l]
	for l < h{
		for l < h && nums[h] > tag{
			h--
		}
		if l < h{
			nums[l] = nums[h]
			l++
		}
		for l < h && nums[l] < tag{
			l++
		}
		if l < h{
			nums[h] = nums[l]
			h--
		}
	}
	nums[l] = tag
	if l == mid{
		return
	}
	if l < mid{
		qsort_wiggleSort(nums,l+1,high,mid)
	}else{
		qsort_wiggleSort(nums,low,l-1,mid)
	}
}

func Recursive_qsort(data []int,low int,high int){
	if low > high{
		return
	}
	l := low
	h := high
	tag := data[l]
	for l < h {
		for l < h && data[h] > tag {
			h--
		}
		if l < h {
			data[l] = data[h]
			l++
		}
		for l < h && data[l] < tag{
			l++
		}
		if l < h{
			data[h] = data[l]
			h--
		}
	}
	data[l] = tag
	Recursive_qsort(data,low,l - 1)
	Recursive_qsort(data,l + 1,high)
}

type Boundary struct {
	low int
	high int
}

func Norecursive_qsort(data []int){
	length := len(data)
	if length == 0{
		return
	}
	var s list.List
	var b Boundary
	b.low = 0
	b.high = length - 1
	s.PushBack(b)
	for s.Len() > 0{
		top := s.Back()
		s.Remove(top)
		l := top.Value.(Boundary).low
		h := top.Value.(Boundary).high
		tag := data[l]
		for l < h{
			for l < h && data[h] > tag {
				h--
			}
			if l < h{
				data[l] = data[h]
				l++
			}
			for l < h && data[l] < tag{
				l++
			}
			if l < h{
				data[h] = data[l]
				h--
			}
		}
		data[l] = tag
		if l - 1 > top.Value.(Boundary).low{
			var b Boundary
			b.low = top.Value.(Boundary).low
			b.high = l - 1
			s.PushBack(b)
		}
		if h + 1 < top.Value.(Boundary).high{
			var b Boundary
			b.low = h + 1
			b.high = top.Value.(Boundary).high
			s.PushBack(b)
		}
	}
}