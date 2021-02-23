package array

import (
	"bytes"
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//Definition for singly-linked list_queue.
type ListNode struct {
	Val int
	Next *ListNode
}

//func remove_duplicated_sorted_array(arr []int)  []int{
//	result := make([]int,0)
//	len := len(arr)
//	cur := arr[0]
//	result = append(result, cur)
//	for i := 1;i<len;i++{
//		if arr[i] == cur{
//			i++
//		}else{
//			result = append(result, arr[i])
//			cur = arr[i]
//		}
//	}
//	return result
//}

func remove_duplicated_sorted_array(arr []int)(length int){
	len := len(arr)
	if len <= 1{
		return
	}
	//cur_val := arr[0]
	cur_pos := 0
	for i := 1 ;i<len;i++{
		if arr[cur_pos] != arr[i]{
			cur_pos++
			arr[cur_pos] = arr[i]
		}
	}
	return cur_pos + 1
}

func sum_exist(arr []int){

}

func findDuplicates(nums []int) []int {
	var res []int
	for i := 0; i < len(nums);{
		if(nums[nums[i] - 1] != nums[i]){
			nums[nums[i] - 1],nums[i] = nums[i],nums[nums[i] - 1]
		} else{
			i++
		}
	}
	for i := 0;i < len(nums);i++{
		if(nums[i]-1 != i){
			res = append(res, nums[i])
		}
	}
	return res
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

		//var left_cap int = int(math.Abs(float64(high - low - 1)) * math.Min(float64(height[high]),float64(height[low+1])))
		//var right_cap int = int(math.Abs(float64(high - low - 1)) * math.Min(float64(height[high - 1]),float64(height[low])))
		if (height[high - 1] >= height[low+1]){
			high--
		} else{
			low++
		}
	}
	return max_cap
}


func distributeCandies(candies int, num_people int) []int {
	var res []int = make([]int,num_people,num_people)
	var cnt int = 1
	for candies > cnt{
		res[(cnt - 1) % num_people] += cnt
		candies -= cnt
		cnt++
	}
	res[(cnt - 1) % num_people] += candies
	return res
}

func findWords(words []string) []string {
	var char_pos map[string]int = make(map[string]int)
	char_pos["q"] = 1
	char_pos["w"] = 1
	char_pos["e"] = 1
	char_pos["r"] = 1
	char_pos["t"] = 1
	char_pos["y"] = 1
	char_pos["u"] = 1
	char_pos["i"] = 1
	char_pos["o"] = 1
	char_pos["p"] = 1

	char_pos["a"] = 2
	char_pos["s"] = 2
	char_pos["d"] = 2
	char_pos["f"] = 2
	char_pos["g"] = 2
	char_pos["h"] = 2
	char_pos["j"] = 2
	char_pos["k"] = 2
	char_pos["l"] = 2

	char_pos["z"] = 3
	char_pos["x"] = 3
	char_pos["c"] = 3
	char_pos["v"] = 3
	char_pos["b"] = 3
	char_pos["n"] = 3
	char_pos["m"] = 3

	var res []string

	for _,str := range words{
		var last_pos int = -1
		var equal bool = true
		s := strings.ToLower(str)
		for i := 0;i < len(s);i++{
			ele := string(s[i])
			if(last_pos == -1){
				last_pos = char_pos[ele]
			}else{
				if(last_pos != char_pos[ele]){
					equal = false
					break
				}
			}
		}
		if(equal){
			res = append(res, str)
		}
	}
	return res
}

//[[518,518],[71,971],[121,862],[967,607],[138,754],[513,337],[499,873],[337,387],[647,917],[76,417]]
//
func twoCitySchedCost(costs [][]int) int {
	var res int = 0
	var num int = len(costs)
	if num <= 1{
		return res
	}
	var a_cnt int = 0
	var b_cnt int = 0
	for i := 0;i < num;i++{
		for j := 1;j < num - i;j++{
			if math.Abs(float64(costs[j][0] - costs[j][1])) > math.Abs(float64(costs[j-1][0] - costs[j-1][1])){
				costs[j],costs[j-1] = costs[j-1],costs[j]
			}
		}
	}
	for _,ele := range costs{
		if ele[0] > ele[1] && b_cnt < num/2 {
			res += ele[1]
			b_cnt++
		} else{
			res += ele[0]
			a_cnt++
		}
	}
	return res
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

//Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
//Output: 3
//3
//[[[1,3],[2,3],[3,1]]
func findJudge(N int, trust [][]int) int {
	var len int = len(trust)
	var record []int = make([]int,N+1)
	var i int = 0
	for i < len{
		record[trust[i][0]] = -1
		if record[trust[i][1]] != -1{
			record[trust[i][1]]++
		}
		i++
	}
	for index,e := range record{
		if e == (N-1){
			return index
		}
	}
	return -1
}

type IntSlice []int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func findContentChildren(g []int, s []int) int {
	sort.Sort(IntSlice(g))
	sort.Sort(IntSlice(s))
	var len_g int = len(g)
	var len_s int = len(s)
	i := 0
	j := 0
	for j < len_s && i < len_g{
		if s[j] < g[i]{
			j++
		}else{
			i++
			j++
		}
	}
	return i
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

//func countBinarySubstrings(s string) int {
//	var res int = 0
//	var len int = len(s)
//	for i := 2;i <= len;i = i + 2{
//		start := 0
//		end := i
//		for ;end <= len;{
//			if s[start] != s[end-1] {
//				if is_binary_string(s[start:end],end-start) {
//					res++
//				}
//			}
//			start++
//			end++
//		}
//	}
//	return res
//}

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
	sort.Sort(IntSlice(ele.nums))
	ele.largest = make([]int, k, k)
	copy(ele.largest,ele.nums[0:k])
	return ele
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	this.largest = append(this.largest, val)
	sort.Sort(IntSlice(this.largest))
	this.largest = this.largest[0:this.k]
	return this.nums[this.k - 1]
}


func bfs2(grid [][]int,i int,j int,depth int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	if i < 0 || i >= rows || j < 0 || j >= columns || grid[i][j] != 1{
		return 0
	}
	grid[i][j] = 2
	depth++
	max_depth := depth
	res :=  bfs2(grid,i-1,j,depth)
	if res > max_depth{
		max_depth = res
	}
	res = bfs2(grid,i+1,j,depth)
	if res > max_depth{
		max_depth = res
	}
	res = bfs2(grid,i,j-1,depth)
	if res > max_depth{
		max_depth = res
	}
	res = bfs2(grid,i,j+1,depth)
	if res > max_depth{
		max_depth = res
	}
	return max_depth
}

//[[2],[1],[1],[1],[2],[1[1]]
func orangesRotting2(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	max_res := 0
	for i := 0;i < rows;i++{
		for j :=0;j < columns;j++{
			if grid[i][j] == 2{
				res := bfs2(grid,i-1,j,0)
				if res > max_res{
					max_res = res
				}
				res = bfs2(grid,i+1,j,0)
				if res > max_res{
					max_res = res
				}
				res = bfs2(grid,i,j-1,0)
				if res > max_res{
					max_res = res
				}
				res = bfs2(grid,i,j+1,0)
				if res > max_res{
					max_res = res
				}
			}
		}
	}
	for i := 0;i < rows;i++{
		for j :=0;j < columns;j++{
			if grid[i][j] == 1{
				return -1
			}
		}
	}
	return max_res
}
//[30,20,150,100,40]
func numPairsDivisibleBy60(time []int) int {
	res := 0
	dict := make(map[int]int)
	l := len(time)
	for i := 0;i < l;i++{
		abs_val := time[i]%60
		if cnt,ok := dict[abs_val];ok && cnt > 0{
			res += dict[abs_val]
		}
		needed := (600 - time[i])%60
		if cnt,ok := dict[needed];ok && cnt > 0{
			dict[needed]++
		}else{
			dict[needed] = 1
		}
	}
	return res
}

//744
func nextGreatestLetter(letters []byte, target byte) byte {
	var low int = 0
	var high int = len(letters) - 1
	if target < letters[low] || target > letters[high]{
		return letters[low]
	}
	var res byte = 0
	for low < high{
		var mid int = (low + high)/2
		if letters[mid] == target{
			if mid + 1 <= high{
				res = letters[mid+1]
			}
			break
		}else if letters[mid] < target{
			low = mid
			if letters[low] > target{
				res = letters[low]
				break
			}
		}else{
			high = mid
			if letters[high] < target{
				res = letters[mid]
				break
			}
		}
	}
	return res
}

func get_value(n int) int {
	if n == 0{
		return 1
	}else if n == 1 {
		return n
	}else{
		return n * get_value(n - 1)
	}
}

func gen_last_value(n, m int) int{
	first := get_value(n)
	second := get_value(m)
	third := get_value((n - m))
	return first / (second * third)
}

func numEquivDominoPairs(dominoes [][]int) int {
	var res int = 0
	var dict map[[2]int]int = make(map[[2]int]int)
	for _,ele := range dominoes{
		var max,min int
		if ele[0] > ele[1]{
			max = ele[0]
			min = ele[1]
		}else{
			max = ele[1]
			min = ele[0]
		}

		var key [2]int = [2]int{min,max}
		if _,ok := dict[key];ok{
			dict[key] += 1
		}else{
			dict[key] = 1
		}
	}
	for _,num := range dict{
		if num >= 2{
			res += gen_last_value(num,2)
		}
	}
	return res
}

//763 eager way
func partitionLabels(S string) []int {
	var res []int
	var dict [26]int
	for pos,e := range S{
		dict[e - 'a'] = pos
	}
	l := len(S)
	i := 0
	start := 0
	last := 0
	for i < l {
		if dict[S[i] - 'a'] > last{
			last = dict[S[i] - 'a']
		}
		if last == i{
			if last == i{
				res = append(res, last - start + 1)
				start = i + 1
			}
		}
		i++
	}
	return res
}

//1051
func heightChecker(heights []int) int {
	var cnt int = 0

	return cnt
}

func check_balance(s string) bool{
	var L_cnt int = 0
	var R_cnt int = 0
	for i := 0 ;i < len(s);i++{
		if s[i] == 'L'{
			L_cnt++
		}else{
			R_cnt++
		}
	}
	return L_cnt == R_cnt
}

//1221
//"RLRRLLRLRL"
func balancedStringSplit(s string) int {
	var cnt int = 0
	begin := 0
	cur := 2
	for ;begin < len(s) && cur <= len(s);{
		sub := s[begin:cur]
		if check_balance(sub) {
			begin = cur
			cur = begin + 1
			cnt++
		}else{
				cur++
		}
	}
	return cnt
}



//861
func matrixScore(A [][]int) int {
	//reverse line
	for i := 0;i < len(A); i++{
		if A[i][0] != 1{
			for j := 0;j < len(A[i]) ; j++{
				if A[i][j] == 0{
					A[i][j] = 1
				}else{
					A[i][j] = 0
				}
			}
		}
	}
	//reverse column
	for i := 0; i < len(A[0]); i++{
		var zero_cnt int = 0
		for j := 0 ; j < len(A);j++{
			if A[j][i] == 0{
				zero_cnt++
			}
		}
		if zero_cnt > len(A)/2{
			for j := 0 ; j < len(A);j++{
				if A[j][i] == 0{
					A[j][i] = 1
				}else{
					A[j][i] = 0
				}
			}
		}
	}
	//calculate
	var res int = 0
	for i := 0;i < len(A); i++{
		var val int = 0
		for j := len(A[i]) - 1;j >= 0 ; j-- {
			val += A[i][j] << uint32(len(A[i]) - 1 - j)
		}
		res += val
	}
	return res
}

//1222
func check_point_inline(start []int,end []int,third []int) bool{

	if (third[0] - start[0])*(end[1] - start[1]) == (end[0] - start[0]) * (third[1] - start[1]) &&
	   math.Min(float64(start[0]),float64(end[0])) < float64(third[0]) && float64(third[0]) < math.Max(float64(start[0]),float64(end[0])) &&
		math.Min(float64(start[1]),float64(end[1])) < float64(third[1]) && float64(third[1]) < math.Max(float64(start[1]),float64(end[1])){
		return true
	}else{
		return false;
	}
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var res [][] int
	var samelines [][]int
	var samecolumns [][]int
	var sametilts [][]int
	for i := 0; i < len(queens);i++{
		if queens[i][0] == king[0]{
			samelines = append(samelines, queens[i])
		}else if queens[i][1] == king[1]{
			samecolumns = append(samecolumns,queens[i])
		}else if math.Abs(float64(queens[i][0] - king[0])) == math.Abs(float64(queens[i][1] - king[1])){
			sametilts = append(sametilts,queens[i])
		}
	}
	min_left_distance := -1
	min_right_distance := 8
	var left_pos []int
	var right_pos []int
	for i := 0;i < len(samelines);i++{
		if samelines[i][1] < king[1]{
			if samelines[i][1] > min_left_distance{
				min_left_distance = samelines[i][1]
				left_pos = samelines[i]
			}
		}else{
			if samelines[i][1] < min_right_distance{
				min_right_distance = samelines[i][1]
				right_pos = samelines[i]
			}
		}
	}
	if len(left_pos) > 0{
		res = append(res, left_pos)
	}
	if len(right_pos) > 0{
		res = append(res,right_pos)
	}
	min_upper_pos := -1
	min_bottom_pos := 8
	var upper_pos []int
	var bottom_pos []int
	for i := 0;i < len(samecolumns) ;i++ {
		if samecolumns[i][0] < king[0]{
			if samecolumns[i][0] > min_upper_pos{
				min_upper_pos = samecolumns[i][0]
				upper_pos = samecolumns[i]
			}
		}else{
			if samecolumns[i][0] < min_bottom_pos{
				min_bottom_pos = samecolumns[i][0]
				bottom_pos = samecolumns[i]
			}
		}
	}
	if len(upper_pos) > 0{
		res = append(res, upper_pos)
	}
	if len(bottom_pos) > 0{
		res = append(res,bottom_pos)
	}
	var min_tilt_collection [][]int //save nearest queens
	var first bool = true
	for i := 0;i < len(sametilts);i++{
		if first{
			first = false
			min_tilt_collection = append(min_tilt_collection, sametilts[i])
			continue
		}
		var add_to_result bool = true
		for j := 0;j < len(min_tilt_collection);{
			if check_point_inline(king,sametilts[i],min_tilt_collection[j]) {
				add_to_result = false
				break
			}

			if check_point_inline(king,min_tilt_collection[j],sametilts[i]) {
				fmt.Println("tilt ", sametilts[i], " closer than ", min_tilt_collection[j])
				min_tilt_collection = append(min_tilt_collection[:j], min_tilt_collection[j+1:]...)
			}else{
				j++
			}
		}
		if add_to_result{
			fmt.Println( sametilts[i], "  is valid")
			min_tilt_collection = append(min_tilt_collection,sametilts[i])
		}
	}
	res = append(res,min_tilt_collection...)
	return res
}

//811
//Input:
//["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
//Output:
//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
func subdomainVisits(cpdomains []string) []string {
	var res []string
	var records map[string]int = make(map[string]int);
	for _,ss := range cpdomains{
		item := strings.Split(ss," ")
		cnt,ok := strconv.Atoi(item[0])
		if ok != nil{
			return res
		}
		domain := item[1]
		for len(domain) > 0{
			if _,ok := records[domain];ok{
				records[domain] += cnt
			}else{
				records[domain] = cnt
			}
			pos := strings.Index(domain,".")
			if pos == -1{
				break
			}
			domain = domain[pos+1:]
		}
	}
	for domain,cnt := range records{
		s := strconv.Itoa(cnt) + " " + domain
		res = append(res,s)
	}
	return res
}

//419
//var total int = 0
//func dfs_419(board [][]byte,x int,y int,last_is_ship bool) bool{
//	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0]) - 1{
//		return false
//	}
//	if board[x][y] == 'X'{
//
//	}
//
//	dfs_419(board,x + 1,y,board[x][y] == 'X')
//	dfs_419(board,x ,y + 1,board[x][y] == 'X')
//}
func point_is_ship(board [][]byte,x int,y int) bool{
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0]) - 1{
		return false
	}
	return board[x][y] == 'X'
}

func countBattleships(board [][]byte) int {
	var res int
	for i := 0; i < len(board);i++{
		for j := 0; j < len(board[0]);j++{
			if board[i][j] == 'X' && !point_is_ship(board,i+1,j) && !point_is_ship(board,i,j+1){
				res++
			}
		}
	}
	return res
}

//1104

func pathInZigZagTree(label int) []int {
	var res []int
	res = append(res, label)
	var layer_num int = math.Ilogb(float64(label))

	for layer_num > 1{
		var parent int = label/2
		layer_min := int(math.Pow(2,float64(layer_num - 1)))
		layer_max := int(math.Pow(2,float64(layer_num)) - 1)
		parent = layer_max - (parent - layer_min)
		label = parent
		layer_num--
		res = append(res, parent)
	}
	res = append(res, 1)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
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

//1252
//n = 2, m = 3, indices = [[0,1],[1,1]]
func oddCells(n int, m int, indices [][]int) int {
	var res int = 0
	var matrix [][]int = make([][]int,n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, m)
	}
	for i:= 0; i < len(indices);i++{
		row := indices[i][0]
		for c:= 0; c < m;c++{
			matrix[row][c]++
		}
		column := indices[i][1]
		for r := 0;r < n;r++{
			matrix[r][column]++
		}
	}
	for i := 0; i < n;i++{
		for j := 0;j < m;j++{
			if matrix[i][j]%2 == 1{
				res++
			}
		}
	}
	fmt.Println(matrix)
	return res
}


//1078
func findOcurrences(text string, first string, second string) []string {
	var res []string
	var words []string = strings.Split(text," ")
	if len(words) <= 1{
		return res
	}
	for slow := 0; slow < len(words) - 2;slow++{
		if words[slow] == first{
			if words[slow+1] == second{
				res = append(res, words[slow+2])
			}
		}
	}
	return res
}

//986
func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	index_a,index_b := 0,0
	for index_a < len(A) && index_b < len(B){
		if A[index_a][1] < B[index_b][0]{
			index_a++
		}else if B[index_b][1] < A[index_a][0]{
			index_b++
		}else{
			max_start := 0
			if A[index_a][0] > B[index_b][0]{
				max_start = A[index_a][0]
			}else{
				max_start = B[index_b][0]
			}
			min_end := 0
			if A[index_a][1] > B[index_b][1]{
				min_end = B[index_b][1]
				index_b++
			}else{
				min_end = A[index_a][1]
				index_a++
			}
			res = append(res, []int{max_start,min_end})
		}
	}
	return res
}

// 791
// S = "cba"
// T = "abcd"
// Output: "cbad"
func customSortString(S string, T string) string {
	var records map[int32]int = make(map[int32]int)
	for i,val := range S{
		records[val] = i
	}

	for i := 1;i < len(T) - 1;i++{
		for j := 0; j < len(T) - i;j++{
			var first_score,second_score int = -1,-1
			if _,ok := records[int32(T[j])];ok{
				first_score = records[int32(T[j])]
			}
			if _,ok := records[int32(T[j+1])];ok{
				second_score = records[int32(T[j+1])]
			}
			if first_score > second_score{
				c := []byte(T)  // 将字符串 s 转换为 []byte 类型
				c[j] = T[j+1]
				c[j+1] = T[j]
				T = string(c)
			}
		}
	}
	return T
}

//1043
func max_int_slice(nums []int) int{
	var max int = 0
	for _,val := range nums{
		if val > max{
			max = val
		}
	}
	return max
}

func max_sum(A []int,begin int,K int,records map[int]int) int{
	var max int = 0
	if begin >= len(A) {
		return max
	}
	for i := 1;i <= K;i++{
		if (begin + i)  > len(A) {
			break
		}
		var sum int = 0
		if val,ok := records[begin+i];ok{
			sum = val
		}else{
			sum = max_sum(A,begin + i,K,records)
			records[begin + i] = sum
		}
		cur_sum := max_int_slice(A[begin:begin+i]) * (i) + sum
		if cur_sum > max{
			max = cur_sum
		}
	}
	return max
}

func maxSumAfterPartitioning(A []int, K int) int {
	var records map[int]int = make(map[int]int,len(A))
	return max_sum(A,0,K,records)
}

//134
// [1,2,3,4,5]
// [3,4,5,1,2]

//	{2,3,4}
//	{3,4,3}
func canCompleteCircuit(gas []int, cost []int) int {
	var start_point int = -1
	var l = len(gas)
	for begin := 0; begin <= len(gas);begin++ {
		trace := begin + 1
		cur_gas := 0
		for begin !=  trace % l {
			cur_gas += gas[(trace-1) % l]
			if cur_gas < cost[(trace-1) % l]{
				break
			}else{
				cur_gas -= cost[(trace-1) % l]
			}
			trace++
		}
		cur_gas += gas[(trace-1) % l]
		if begin == trace % l && cur_gas >= cost[(trace-1) % l]{
			return begin
		}
	}
	return start_point
}

//841
// [[1],[2],[3],[]]
//  [[1,3,2],[2,3],[2,1,3,1],[]]
//  [[1,3],[3,0,1],[2],[0]]
func canVisitAllRooms(rooms [][]int) bool {
	var room_num int = len(rooms)
	var visited map[int]bool = make(map[int]bool,room_num)
	for i := 0; i < room_num;i++{
		visited[i] = false
	}
	var target int = (room_num - 1) * room_num / 2
	var cur_num int = 0
	l := list.New()
	l.PushBack(0)
	for l.Len() != 0{
		ele := l.Back()
		val := ele.Value.(int)
		for i := 0; i < len(rooms[val]);i++{
			if !visited[rooms[val][i]]{
				cur_num += rooms[val][i]
				visited[rooms[val][i]] = true
				l.PushBack(rooms[val][i])
			}
		}
		l.Remove(ele)
		if cur_num >= target{
			return true
		}
	}
	return false
}

//1219
func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func dfs_gold(grid [][]int,x int,y int,record [][]int)int{
	if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0{
		return 0
	}
	if grid[x][y] == 0 || record[x][y] != 0{
		return 0
	}
	record[x][y] = 1
	var dup_record1,dup_record2,dup_record3,dup_record4 [][]int = make([][]int,len(grid)),make([][]int,len(grid)),make([][]int,len(grid)),make([][]int,len(grid))
	for i := 0;i < len(grid);i++ {
		dup_record1[i] = make([]int,len(grid[i]))
		dup_record2[i] = make([]int,len(grid[i]))
		dup_record3[i] = make([]int,len(grid[i]))
		dup_record4[i] = make([]int,len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			dup_record1[i][j] = record[i][j]
			dup_record2[i][j] = record[i][j]
			dup_record3[i][j] = record[i][j]
			dup_record4[i][j] = record[i][j]

		}
	}
	var res int = 0
	res = max_int(grid[x][y] + dfs_gold(grid,x - 1,y,dup_record1),grid[x][y] + dfs_gold(grid,x,y-1,dup_record2))
	res = max_int(res,grid[x][y] + dfs_gold(grid,x + 1,y,dup_record3))
	res = max_int(res,grid[x][y] + dfs_gold(grid,x,y + 1,dup_record4))
	return res
}

func getMaximumGold(grid [][]int) int {
	var max int = 0
	for i := 0;i < len(grid);i++{
		for j := 0;j < len(grid[i]);j++{
			if grid[i][j] != 0{
				var record [][]int = make([][]int,len(grid))
				for row := 0; row < len(grid);row++{
					record[row] = make([]int,len(grid[i]))
				}
				max = max_int(max,dfs_gold(grid,i,j,record))
			}
		}
	}
	return max
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


//1277
// Input: matrix =
// [
//  [0,1,1,1],
//  [1,1,1,1],
//  [0,1,1,1]
// ]
// Output: 15
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

func countSquares(matrix [][]int) int {
	var cnt int = 0
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				continue
			}
			matrix[i][j] = 1 + min_int_number(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
	}
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			cnt += matrix[i][j]
		}
	}
	return cnt
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

func maximalSquare(matrix [][]byte) int {
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			matrix[i][j] -= '0'
		}
	}

	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				continue
			}
			matrix[i][j] = 1 + min_byte_number(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
	}
	var res int = 0
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			if int(matrix[i][j]) > res{
				res = int(matrix[i][j])
			}
		}
	}
	return res * res
}

//14
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0{
		return ""
	}else if l == 1{
		return strs[0]
	}
	var res string = strs[0]
	for i := 1;i < l;i++{
		for j:= 0;j < len(strs[i]) && j < len(res);j++{
			if strs[i][j] != res[j]{
				if j == 0{
					return ""
				}
				res = res[0:j]
				break
			}
		}
	}
	return res
}

//1143
//Input: text1 = "abcde", text2 = "ace"
//Output: 3
func longestCommonSubsequence(text1 string, text2 string) int {
	return 0
}

//1254
func closedIsland(grid [][]int) int {
	return 0
}

//931
func dp_minpath(A [][]int,i int,j int,memo [][]int)int{
	if i >= len(A) || i < 0 || j >= len(A[0]) || j < 0 {
		return 0
	}
	if memo[i][j] != 0{
		return memo[i][j]
	}
	var res int = math.MaxInt32
	res = int(math.Min(float64(A[i][j] + dp_minpath(A,i+1,j,memo)),float64(res)))
	if (j + 1) < len(A[0]) {
		res = int(math.Min(float64(A[i][j]+dp_minpath(A, i+1, j+1,memo)), float64(res)))
	}
	if (j - 1) >= 0 {
		res = int(math.Min(float64(A[i][j]+dp_minpath(A, i+1, j-1,memo)), float64(res)))
	}
	memo[i][j] = res
	return res
}

func minFallingPathSum(A [][]int) int {
	var res int
	var memo [][]int = make([][]int,len(A))
	for i := 0; i < len(A);i++{
		memo[i] = make([]int,len(A[0]))
	}
	for j := 0; j < len(A[0]);j++{
		res = int(math.Min(float64(dp_minpath(A,0,j,memo)),float64(math.MaxInt32)))
	}
	for _,v := range memo[0]{
		if v < res{
			res = v
		}
	}
	return res
}

//1260
// [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]] K = 4
func shiftGrid(grid [][]int, k int) [][]int {
	var row_num int = len(grid)
	var col_num int = len(grid[0])
	var total = row_num * col_num
	var res []int = make([]int,total)
	for i := 0;i < row_num;i++{
		for j := 0;j < col_num;j++{
			res[(i * col_num + j + k) % total] = grid[i][j]
		}
	}
	var grid2 [][]int = make([][]int,len(grid))
	for i := 0;i < row_num;i++{
		grid2[i] = make([]int,len(grid[0]))
		for j := 0;j < col_num;j++{
			grid2[i][j] = res[i * col_num + j]
		}
	}
	return grid2
}

//1072
func maxEqualRowsAfterFlips(matrix [][]int) int {
	var res int = 0
	var record map[string]int = make(map[string]int)
	for row := 0; row < len(matrix);row++{
		var s string = ""
		var reverse_s string = ""
		for column := 0; column < len(matrix[0]);column++{
			if matrix[row][column] == 1{
				reverse_s += "0"
			}else{
				reverse_s += "1"
			}
			s += strconv.Itoa(matrix[row][column])
		}
		if _,ok := record[reverse_s];ok{
			record[reverse_s]++
		}else{
			record[reverse_s] = 1
		}
		if _,ok := record[s];ok{
			record[s]++
		}else{
			record[s] = 1
		}
	}
	for _,val := range record{
		if val > res{
			res = val
		}
	}
	return res
}

//1249
// "())()((("
func reverse(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		tmp := s1[i]
		s1[i] = s1[len(s1)-1-i]
		s1[len(s1)-1-i] = tmp
	}
	return string(s1)
}

func minRemoveToMakeValid(s string) string {
	var res string
	var start_tag int = 0
	for i := 0; i < len(s);i++{
		if s[i] == '('{
			start_tag++
		}else if s[i] == ')'{
				if start_tag <= 0{
					continue
				}else{
					start_tag--
				}
		}
		res += string(s[i])
	}
	var res2 string
	i := len(res) - 1
	for ; i >= 0 && start_tag > 0;i--{
		if res[i] == '('{
			start_tag--
			continue
		}
		res2 += string(res[i])
	}
	res3 := res[0:i+1] + reverse(res2)
	return res3
}

//1074
//Input: matrix = [[1,-1],[-1,1]], target = 0
//Output: 5
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var res int = 0
	for row := 0;row < len(matrix);row++{
		for col := 0;col < len(matrix[0]);col++{

		}
	}
	return res
}

//260
//Input:  [1,2,1,3,2,5]
//Output: [3,5]
func singleNumber(nums []int) []int {
	var val int = 0
	for i := 0;i < len(nums);i++{
		val = val ^ nums[i]
	}
	var tag int = 1
	for i := 1;i < 32;i++{
		if val | tag == val{
			break
		}else{
			tag = tag << 1
		}
	}
	fmt.Println(tag)
	var group [][]int = make([][]int,2)
	for i := 0;i < len(nums);i++{
		if nums[i] | tag == nums[i]{
			group[0] = append(group[0],nums[i])
		}else{
			group[1] = append(group[1],nums[i])
		}
	}
	var first_num,second_num int = 0,0
	for i := 0;i < len(group[0]);i++{
		first_num ^= group[0][i]
	}
	for i := 0;i < len(group[1]);i++{
		second_num ^= group[1][i]
	}
	return []int{first_num,second_num}
}

//647
func try_palindromic(s string,start int,end int) int{
	var cnt int = 0
	for start >= 0 && end < len(s){
		if s[start] == s[end]{
			cnt++
		}else{
			break
		}
		start--
		end++
	}
	return cnt
}

//647
func countSubstrings(s string) int {
	var res int = 0
	for start := 0;start < len(s);start++{
		res += try_palindromic(s,start,start)
		fmt.Println(res)
		res += try_palindromic(s,start,start+1)
		fmt.Println(res)
	}
	return res
}

//413
func is_same_diff(A []int,left int,right int) (diff int,ret bool){
	if left < 0 || right >= len(A){
		diff = 0
		ret = false
		return
	}
	diff = A[left+1] - A[left]
	ret = true
    for i := left + 1;i < right;i++{
    	if (A[i+1] - A[i]) != diff{
    		ret = false
    		break
		}
	}
	return
}

func numberOfArithmeticSlices(A []int) int {
	var l int = len(A)
	if l < 3{
		return 0
	}
	var record map[int]int = make(map[int]int)
	for start,end := 0, 2;start < l && end < l;{
		_,ok := is_same_diff(A,start,end)
		if ok{
			end++
			if _,ok := record[start];ok{
				record[start]++
			}else{
				record[start] = 1
			}
		}else{
			start = end - 1
			end = start + 2
		}
	}
	res := 0
	for _,val := range record{
		for val > 0{
			res += val
			val--
		}
	}
	return res
}



//1281
func subtractProductAndSum(n int) int {
	var sum int = 0
	var times int = 1
	for n > 0{
		val := n % 10
		sum += val
		times *= val
		n = n/10
	}
	return times - sum
}

//1266
func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

func minTimeToVisitAllPoints(points [][]int) int {
	var res int = 0
	for pos := 1; pos < len(points);pos++{
		res += max_int(abs_int(points[pos][0] - points[pos-1][0]),abs_int(points[pos][1] - points[pos-1][1]))
	}
	return res
}

//1094
//func carPooling(trips [][]int, capacity int) bool {
//
//}

//395
func longestSubstring(s string, k int) int{
	l := len(s)
	var less_than_k map[uint8]int = make(map[uint8]int)
	for i := 0;i < l;i++{
		if _,ok := less_than_k[s[i]];ok{
			less_than_k[s[i]]++
		}else{
			less_than_k[s[i]] = 1
		}
	}
	for c,n := range less_than_k{
		if n >= k{
			delete(less_than_k,c)
		}
	}
	if len(less_than_k) == 0{
		return l
	}
	if len(less_than_k) == l{
		return 0
	}
	var trace []int = make([]int,l)
	for i := 0;i < l;i++{
		if _,ok := less_than_k[s[i]];ok{
			trace[i] = -1
		}
	}
	var max_len = 0
	start := 0
	for start < l {
		for start < l{
			if trace[start] != - 1{
				break
			}
			start++
		}
		if start == l{
			break
		}
		end := start
		for end < l{
			if trace[end] == -1{
				break
			}
			end++
		}
		res := longestSubstring(s[start:end],k)
		if res > max_len{
			max_len = res
		}
		start++
	}
	return max_len
}

//func longestSubstring(s string, k int) int {
//	var record [26]int
//	for _,v := range s{
//		record[v - 'a']++
//	}
//	var less_than_k map[byte]int = make(map[byte]int)
//	var i byte = 0
//	for ;i < 26;i++{
//		if record[i] != 0 && record[i] < k{
//			less_than_k[i] = record[i]
//		}
//	}
//	begin,end := 0,len(s) - 1
//	for begin < len(s) && begin < end{
//		if len(less_than_k) == 0{
//			break
//		}
//		_,okbegin := less_than_k[s[begin] - 'a']
//		_,okend := less_than_k[s[end] - 'a']
//		if !okbegin && !okend{
//			begin++
//			end--
//			continue
//		}
//		if _,ok := less_than_k[s[begin] - 'a'];ok{
//			less_than_k[s[begin] - 'a']--
//			if less_than_k[s[begin] - 'a'] == 0{
//				delete(less_than_k,s[begin] - 'a')
//			}
//			begin++
//		}
//		if begin < end{
//			if _,ok := less_than_k[s[end] - 'a'];ok{
//				less_than_k[s[end] - 'a']--
//				if less_than_k[s[end] - 'a'] == 0{
//					delete(less_than_k,s[end] - 'a')
//				}
//				end--
//			}
//		}
//		if len(less_than_k) == 0{
//			break
//		}
//	}
//	if len(less_than_k) > 0{
//		return 0
//	}
//	sub := s[begin:end+1]
//	var rest map[int32]int = make(map[int32]int)
//	for _,v := range sub{
//		if _,ok := rest[v];ok{
//			rest[v]++
//		}else{
//			rest[v] = 1
//		}
//	}
//	for _,v := range rest{
//		if v < k{
//			return 0
//		}
//	}
//	return end - begin + 1
//}

//326
func isPowerOfThree(n int) bool {
	if n > 1{
		for n > 1{
			if n % 3 != 0{
				return false
			}
			n = n / 3
		}
	}
	return n == 1
}

//179
func compare_pre_digit(a int,b int)bool{
	var s1 string = strconv.Itoa(a)
	var s2 string = strconv.Itoa(b)
	return strings.Compare(s1 + s2,s2 + s1) <= 0
}

func largestNumber(nums []int) string {
	for i := 0;i < len(nums);i++{
		for j := 1;j < len(nums) - i;j++{
			if compare_pre_digit(nums[j-1],nums[j]){
				nums[j-1],nums[j] = nums[j],nums[j-1]
			}
		}
	}
	var res string
	for _,v := range nums{
		res += strconv.Itoa(v)
	}
	res = strings.TrimLeft(res,"0")
	if res == ""{
		return "0"
	}
	return res
}

//287
func findDuplicate(nums []int) int {
	min := 0
	max := len(nums)-1
	for min < max{
		mid := min + (max - min)/2
		cnt := 0
		for _,v := range nums{
			if v <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			min = mid + 1
		}else{
			max = mid
		}
	}
	return min
}

//289
//Input:
//[
//  [0,1,0],
//  [0,0,1],
//  [1,1,1],
//  [0,0,0]
//]
//Output:
//[
//  [0,0,0],
//  [1,0,1],
//  [0,1,1],
//  [0,1,0]
//]
func get_status(board [][]int,i,j int)int{
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
		return 0
	}
	return board[i][j]
}

func gameOfLife(board [][]int)  {
	var neighbour_cnt [][]int = make([][]int,len(board))
	for i := 0;i < len(board);i++{
		neighbour_cnt[i] = make([]int,len(board[0]))
		for j := 0;j < len(board[0]);j++{
			neighbour_cnt[i][j] = get_status(board,i - 1,j - 1) + get_status(board,i - 1,j) + get_status(board,i - 1,j + 1) + get_status(board,i,j - 1) + get_status(board,i,j + 1)+ get_status(board,i + 1,j - 1) + get_status(board,i + 1,j) + get_status(board,i + 1,j + 1)
		}
	}
	for i := 0;i < len(board);i++{
		for j := 0;j < len(board[0]);j++{
			if board[i][j] == 1{
				if neighbour_cnt[i][j] < 2 || neighbour_cnt[i][j] > 3{
					board[i][j] = 0
				}
			}else{
				if neighbour_cnt[i][j] == 3{
					board[i][j] = 1
				}
			}
		}
	}
}

//1295
func findNumbers(nums []int) int {
	var res int = 0
	for i := 0;i < len(nums);i++{
		if len(strconv.Itoa(nums[i])) % 2 == 0{
			res++
		}
	}
	return res
}

//1275
func is_line(steps [][]int)bool{
	if len(steps) < 3{
		return false
	}
	var row_map map[int]int = make(map[int]int)
	var col_map map[int]int = make(map[int]int)
	var tilt_line_points int = 0
	var reverse_tilt_line_points int = 0
	for i := 0;i < len(steps);i++{
		if _,ok := row_map[steps[i][0]];ok{
			row_map[steps[i][0]] += steps[i][1] + 1
		}else{
			row_map[steps[i][0]] = steps[i][1] + 1
		}
		if _,ok := col_map[steps[i][1]];ok{
			col_map[steps[i][1]] += steps[i][0] + 1
		}else{
			col_map[steps[i][1]] = steps[i][0] + 1
		}
		if steps[i][0] == steps[i][1]{
			tilt_line_points++
		}
		if (steps[i][0] == 0 && steps[i][1] == 2) || (steps[i][0] == 1 && steps[i][1] == 1)||(steps[i][0] == 2 && steps[i][1] == 0)  {
			reverse_tilt_line_points++
		}
	}
	for _,val := range row_map{
		if val == 6{
			return true
		}
	}
	for _,val := range col_map{
		if val == 6{
			return true
		}
	}
	if tilt_line_points == 3 || reverse_tilt_line_points == 3{
		return true
	}
	return false
}

func tictactoe(moves [][]int) string {
	var a_steps [][]int
	var b_steps [][]int
	for i := 0;i < len(moves);i++{
		if i % 2 == 0{
			a_steps = append(a_steps,moves[i])
		}else{
			b_steps = append(b_steps,moves[i])
		}
	}
	a_win := is_line(a_steps)
	b_win := is_line(b_steps)
	if !a_win && !b_win{
		if len(moves) < 9{
			return "Pending"
		}else{
			return "Draw"
		}
	}else if a_win{
		return "A"
	}else {
		return "B"
	}
}

//200
//Input:
//11000
//11000
//00100
//00011
//
//Output: 3
func dfs_numIslands(grid [][]byte,row int,col int){
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != '1'{
		return
	}
	grid[row][col] = '0'
	dfs_numIslands(grid,row - 1,col)
	dfs_numIslands(grid,row + 1,col)
	dfs_numIslands(grid,row,col - 1)
	dfs_numIslands(grid,row,col + 1)
}

func numIslands(grid [][]byte) int {
	var res int = 0
	for i := 0;i < len(grid);i++{
		for j := 0;j < len(grid[0]);j++{
			if grid[i][j] == '1'{
				dfs_numIslands(grid,i,j)
				res++
			}
		}
	}
	return res
}

//300
//Input: [10,9,2,5,3,7,101,18]
//Output: 4
//Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
func lengthOfLIS(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	var dp []int = make([]int,len(nums)+1)
	dp[0] = 1
	for i := 1;i < len(nums);i++{
		max := 1
		for j := 0;j < i;j++{
			if nums[i] > nums[j]{
				if (dp[j] + 1) > max{
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max
	}
	var res int = 1
	for _,val := range dp{
		if val > res{
			res = val
		}
	}
	return res
}

//1299
//Input: arr = [17,18,5,4,6,1]
//Output: [18,6,6,6,1,-1]
func replaceElements(arr []int) []int {
	if len(arr) == 0{
		return arr
	}
	var biggest int = arr[len(arr) - 1]
	arr[len(arr) - 1] = -1
	for i := len(arr) - 2;i >= 0;i--{
		if arr[i] > biggest{
			tmp := arr[i]
			arr[i] = biggest
			biggest = tmp
		}else{
			arr[i] = biggest
		}
	}
	return arr
}

//1306
//Input: arr = [4,2,3,0,3,1,2], start = 5
//Output: true
//All possible ways to reach at index 3 with value 0 are:
//index 5 -> index 4 -> index 1 -> index 3
//index 5 -> index 6 -> index 4 -> index 1 -> index 3
func dfs_canReach(arr []int,start int,dp []int)bool{
	if start < 0 || start >= len(arr){
		return false
	}
	if dp[start] == 1{
		return false
	}
	if arr[start] == 0{
		return true
	}
	dp[start] = 1
	forward_res := dfs_canReach(arr,start + arr[start],dp)
	backward_res := dfs_canReach(arr,start - arr[start],dp)
	return forward_res || backward_res
}

func canReach(arr []int, start int) bool {
	var dp []int = make([]int,len(arr))
	return dfs_canReach(arr,start,dp)
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

//994
//Input: [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
func fill_near_orgin(grid [][]int,visited,r,c,rows,cols int){
	if r - 1 >= 0 && grid[r-1][c] == 1{
		grid[r-1][c] = 2
	}
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	var steps int = 0
	var unvisited map[string]bool = make(map[string]bool) //store unvisit good orange position
	var depth [][]int //store current step bad orange
	for row := 0;row < rows;row++{
		for col := 0;col < columns;col++{
			if grid[row][col] == 1{
				k := strconv.Itoa(row)+","+strconv.Itoa(col)
				unvisited[k] = true
			}
			if grid[row][col] == 2{
				depth = append(depth,[]int{row,col})
			}
		}
	}
	if len(unvisited) == 0{// If no good orange exist,return 0
		return 0
	}
	for len(depth) > 0{
		l := len(depth)
		var cur_depth [][]int = make([][]int,len(depth))
		copy(cur_depth,depth)
		depth = make([][]int,0)//clear last level bad orange
		for i := 0;i < l;i++{
			r := cur_depth[i][0]
			c := cur_depth[i][1]
			//check 4 direction's orange
			if r - 1 >= 0 && grid[r-1][c] == 1{
				grid[r-1][c] = 2
				k := strconv.Itoa(r-1)+","+strconv.Itoa(c)
				delete(unvisited,k)
				depth = append(depth, []int{r-1,c})
			}
			if r + 1 < rows && grid[r+1][c] == 1{
				grid[r+1][c] = 2
				k := strconv.Itoa(r+1)+","+strconv.Itoa(c)
				delete(unvisited,k)
				depth = append(depth, []int{r+1,c})
			}
			if c - 1 >= 0 && grid[r][c-1] == 1{
				grid[r][c-1] = 2
				k := strconv.Itoa(r)+","+strconv.Itoa(c-1)
				delete(unvisited,k)
				depth = append(depth, []int{r,c-1})
			}
			if c + 1 < columns && grid[r][c+1] == 1{
				grid[r][c+1] = 2
				k := strconv.Itoa(r)+","+strconv.Itoa(c+1)
				delete(unvisited,k)
				depth = append(depth, []int{r,c+1})
			}
		}
		steps++
		if len(unvisited) == 0{// when there is no good orange,return steps
			return steps
		}
	}
	if len(unvisited) > 0{
		return -1
	}
	return steps
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
//Input: coins = [1, 2, 5], amount = 11
//Output: 3
//Explanation: 11 = 5 + 5 + 1
//Top to bottom
//func dp_coinChange(coins []int,amount int,rest int,cur_index int,memo map[int]int)int{
//	if cur_index >= len(coins) || rest < 0 {
//		return math.MaxInt32
//	}
//	if rest == 0{
//		return 1
//	}
//	//if _,ok := memo[rest];ok{
//	//	return memo[rest]
//	//}
//	steps := math.MaxInt32
//	//choose current coin and continue
//	l1 := 1 + dp_coinChange(coins,amount,rest - coins[cur_index],cur_index,memo)
//	if l1 < steps{
//		steps = l1
//	}
//	//choose current coin and move to next
//	l2  := 1 + dp_coinChange(coins,amount,rest - coins[cur_index],cur_index + 1,memo)
//	if l2 < steps{
//		steps = l2
//	}
//	//skip current coin and move to next
//	l3 := dp_coinChange(coins,amount,rest,cur_index + 1,memo)
//	if l3 < steps{
//		steps = l3
//	}
//	if steps != math.MaxInt32 {
//		if val,ok := memo[rest];ok{
//			if steps < val{
//				memo[rest] = steps
//			}
//		}else{
//			memo[rest] = steps
//		}
//	}
//	return steps
//}
//
//func coinChange(coins []int, amount int) int {
//	sort.Ints(coins)
//	for i := 0;i < len(coins)/2;i++{
//		coins[i],coins[len(coins)-i-1] = coins[len(coins)-i-1],coins[i]
//	}
//	var memo map[int]int = make(map[int]int)
//	res :=  dp_coinChange(coins,amount,amount,0,memo)
//	return res
//}

//func coinChange(coins []int, amount int) int {
//	var memo []int = make([]int,amount + 1)//memo[i] means the min step from 0 to i
//	for i,_ := range memo{
//		memo[i] = math.MaxInt32
//	}
//	memo[0] = 0
//	// var steps int = 0
//	for i := 1;i <= amount;i++{
//		min_steps := math.MaxInt32
//		for _,c := range coins{
//			if (i - c) >= 0{
//				if memo[i-c] < min_steps{
//					min_steps = memo[i-c]
//				}
//			}
//		}
//		if min_steps != math.MaxInt32{
//			memo[i] = 1 + min_steps
//		}
//	}
//	if memo[amount] == math.MaxInt32{
//		return - 1
//	}
//	return memo[amount]
//}

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

//1337
//Input: mat =
//[[1,1,0,0,0],
//[1,1,1,1,0],
//[1,0,0,0,0],
//[1,1,0,0,0],
//[1,1,1,1,1]],
//k = 3
//Output: [2,0,3]
func kWeakestRows(mat [][]int, k int) []int {
	var rows = len(mat)
	if rows == 0{
		return []int{}
	}
	var columns = len(mat[0])
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < rows;i++{
		j := 0
		for ;j < columns;j++ {
			if mat[i][j] == 0 {
				break
			}
		}
		if _,ok := record[j];ok{
			record[j] = append(record[j], i)
		}else{
			record[j] = []int{i}
		}
	}
	var keys []int
	for l := range record {
		keys = append(keys, l)
	}
	sort.Ints(keys)
	var res []int
	var i int = 0
	for _,num := range keys{
		for _,row := range record[num]{
			res = append(res,row)
			i++
			if i == k{
				return res
			}
		}
	}
	return res
}

//394
//s = "3[a]2[bc]", return "aaabcbc".
//s = "3[a2[c]]", return "accaccacc".
//s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
func decodeString(s string) string {
	var word_record []string
	var cnt_record []int
	var visit int = 0
	var word_num int
	var cur_string string
	for visit < len(s){
		if s[visit] == '['{
			cnt_record = append(cnt_record, word_num)
			word_record = append(word_record,cur_string)
			cur_string = ""
			word_num = 0
		}else if s[visit]==']'{
			pre := word_record[len(word_record) - 1]
			cnt := cnt_record[len(cnt_record) - 1]
			for i := 0;i < cnt;i++{
				pre += cur_string
			}
			cur_string = pre
			cnt_record = cnt_record[:len(word_record) - 1]
			word_record = word_record[:len(word_record) - 1]
		}else{
			if s[visit] < '0' || s[visit] > '9'{
				cur_string += string(s[visit])
			}else{
				c,_ := strconv.Atoi(string(s[visit]))
				word_num = word_num * 10 + c
			}
		}
		visit++
	}
	return cur_string
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

//152
//Input: [-2,0,-1]
//Output: 0
//[2,3,-2,4,-2,4,6,-9,3]
func min_int(a,b int)int{
	if a < b {
		return a
	}else{
		return b
	}
}

func dp_maxProduct(nums []int,begin int,memo map[int][]int)(min int,max int){
	if begin < 0{
		return 1,1
	}
	if nums[begin] == 0{
		return 0,0
	}
	if _,ok := memo[begin];ok{
		return memo[begin][0],memo[begin][1]
	}
	s,b := dp_maxProduct(nums,begin - 1,memo)
	big := max_int(max_int(b * nums[begin],s * nums[begin]),nums[begin])
	small := min_int(min_int(b * nums[begin],s * nums[begin]),nums[begin])
	memo[begin] = make([]int,2)
	memo[begin][0] = small
	memo[begin][1] = big
	return small,big
}

func maxProduct(nums []int) int {
	var res int = math.MinInt32
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < len(nums) ;i++{
		_,big :=  dp_maxProduct(nums,i,record)
		if big > res{
			res = big
		}
	}
	return res
}

func maxProduct2(nums []int) int{
	l := len(nums)
	if l == 0{
		return 0
	}
	if l == 1{
		return nums[0]
	}
	var dp_min []int = make([]int,l)
	var dp_max []int = make([]int,l)
	dp_min[0] = nums[0]
	dp_max[0] = nums[0]
	var max int = nums[0]
	for i := 1;i < l;i++{
		product1 := nums[i] * dp_min[i-1]
		product2 := nums[i] * dp_max[i-1]
		dp_min[i] = min_int_number(nums[i],product1,product2)
		dp_max[i] = max_int_number(nums[i],product1,product2)
		if dp_max[i] > max{
			max = dp_max[i]
		}
	}
	return max
}

//560
//Input:nums = [1,1,1], k = 2
//Output: 2
func subarraySum(nums []int, k int) int {
	var res int = 0
	var record map[int]int = make(map[int]int)
	var sum int = 0
	for _,n := range nums{
		sum += n
		if sum == k{
			res += 1
		}
		if _,ok := record[sum - k];ok{
			res += record[sum - k]
		}
		if _,ok := record[sum];ok{
			record[sum]++
		} else{
			record[sum] = 1
		}
	}
	return res
}

//309
//Input: [1,2,3,0,2]
//Output: 3
//Explanation: transactions = [buy, sell, cooldown, buy, sell]
func dp_maxProfit(prices []int,cur_pos int,has_stock bool,is_cooldown bool,nostock_memo map[int]int,hasstock_memo map[int]int)int{
	if cur_pos >= len(prices){
		return 0
	}

	var max int = math.MinInt32
	if is_cooldown{
		max = dp_maxProfit(prices,cur_pos + 1,has_stock,false,nostock_memo,hasstock_memo)
	}else{
		if has_stock{
			if _,ok := hasstock_memo[cur_pos];ok{
				return hasstock_memo[cur_pos]
			}
			max = max_int(prices[cur_pos] + dp_maxProfit(prices,cur_pos + 1,false,true,nostock_memo,hasstock_memo),
				dp_maxProfit(prices,cur_pos + 1,true,false,nostock_memo,hasstock_memo))
			hasstock_memo[cur_pos] = max
		}else{
			if _,ok := nostock_memo[cur_pos];ok{
				return nostock_memo[cur_pos]
			}
			max = max_int(dp_maxProfit(prices,cur_pos + 1,true,false,nostock_memo,hasstock_memo) - prices[cur_pos],
				dp_maxProfit(prices,cur_pos + 1,false,false,nostock_memo,hasstock_memo))
			nostock_memo[cur_pos] = max
		}
	}
	return max
}

func maxProfit3(prices []int) int {
	var nostock_record map[int]int = make(map[int]int)
	var hasstock_record map[int]int = make(map[int]int)
	return dp_maxProfit(prices,0,false,false,nostock_record,hasstock_record)
}

//1314
//Input: nums = [1,2,3,4]
//Output: [2,4,4,4]
//Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
//The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
func decompressRLElist(nums []int) []int {
	l := len(nums)
	var total_length int = 0
	for i := 0;i < l;i += 2{
		total_length += nums[i]
	}
	var res []int = make([]int,total_length)
	var cur_pos int = 0
	for i := 0;i < l;i += 2{
		cnt := nums[i]
		val := nums[i + 1]
		for j := 0;j < cnt;j++{
			res[cur_pos] = val
			cur_pos++
		}
	}
	return res
}

//416
func dp_canPartition(nums []int,cur_sum int,cur_pos int,target int,memo map[int]bool)int{
	if cur_pos >= len(nums) || cur_sum > target{
		return 0
	}
	if _,ok := memo[cur_sum];ok{
		return 0
	}
	if cur_sum == target{
		return 1
	}
	cnt := dp_canPartition(nums,cur_sum + nums[cur_pos],cur_pos + 1,target,memo) + dp_canPartition(nums,cur_sum,cur_pos + 1,target,memo)
	if cnt == 0{
		memo[cur_sum] = false
	}
	return cnt
}

func canPartition(nums []int) bool {
	var sum int = 0
	for _,v := range nums{
		sum += v
	}
	if sum % 2 == 1{
		return false
	}
	var record map[int]bool = make(map[int]bool)
	return dp_canPartition(nums,0,0,sum / 2,record) > 0
}

//148
//Input: 4->2->1->3
//Output: 1->2->3->4
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5
func merge_sortList(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val{
		head = l1
		l1 = l1.Next
	}else{
		head = l2
		l2 = l2.Next
	}
	var visit *ListNode = head
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			visit.Next = l1
			visit = l1
			l1 = l1.Next
		}else{
			visit.Next = l2
			visit = l2
			l2 = l2.Next
		}
	}
	if l1 == nil{
		visit.Next = l2
	}else if l2 == nil{
		visit.Next = l1
	}
	return head
}
//-1->5->3->4->0
func sortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next{
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
		if fast != nil{
			fast = fast.Next
		}
	}
	head2 := slow.Next
	slow.Next = nil
	newhead1 := sortList(head)
	newhead2 := sortList(head2)
	return merge_sortList(newhead1,newhead2)
}

//139
//Input: s = "applepenapple", wordDict = ["apple", "pen"]
//Output: true
//Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//Note that you are allowed to reuse a dictionary word.
//Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//Output: false
func dfs_wordBreak(s string,wordDict []string,cur_pos int,memo map[int]bool)bool{
	if cur_pos >= len(s){
		return true
	}
	if _,ok := memo[cur_pos];ok{
		return false
	}
	sub := s[cur_pos:]
	for _,w := range wordDict{
		if strings.HasPrefix(sub,w){
			if dfs_wordBreak(s,wordDict,cur_pos + len(w),memo){
				return true
			}
		}
	}
	memo[cur_pos] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	var record map[int]bool = make(map[int]bool)
	return dfs_wordBreak(s,wordDict,0,record)
}

//146
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type DataNode struct{
	key int
	value int
}

type LRUCache struct {
	data []DataNode
	capacity int
	cur_size int
}

func Constructor146(capacity int) LRUCache {
	var obj LRUCache
	obj.data = make([]DataNode,capacity)
	obj.capacity = capacity
	obj.cur_size = 0
	return obj
}

func (this *LRUCache) Get(key int) int {
	for i := this.cur_size - 1; i >= 0;i--{
		if this.data[i].key == key{
			tmp := this.data[i]
			for j := i;j < this.cur_size - 1;j++{
				this.data[j] = this.data[j + 1]
			}
			this.data[this.cur_size - 1] = tmp
			return tmp.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	for i := this.cur_size - 1; i >= 0;i--{
		if this.data[i].key == key{
			this.data[i].value = value
			var tmp DataNode
			tmp.key = this.data[i].key
			tmp.value = value
			for j := i;j < this.cur_size - 1;j++{
				this.data[j] = this.data[j + 1]
			}
			this.data[this.cur_size - 1] = tmp
			return
		}
	}

	if this.cur_size == this.capacity{
		this.data = append(this.data[:0],this.data[1:]...)
		this.data = append(this.data,DataNode{key,value})
	}else{
		this.cur_size++
		this.data[this.cur_size - 1] = DataNode{key,value}
	}
}

//1356

type sortBit []int

func (s sortBit) Less(i, j int) bool {
	var bit_cnt1 int = 0
	var bit_cnt2 int = 0
	var move int = 1
	for s[i] >= move{
		if (s[i] | move) == s[i]{
			bit_cnt1++
		}
		move = move << 1
	}
	move = 1
	for s[j] >= move{
		if (s[j] | move) == s[j]{
			bit_cnt2++
		}
		move = move << 1
	}
	if bit_cnt1 == bit_cnt2{
		return s[i] < s[j]
	}
	return bit_cnt1 < bit_cnt2
}

func (s sortBit) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBit) Len() int {
	return len(s)
}

func sortByBits(arr []int) []int {
	sort.Sort(sortBit(arr))
	return arr
}

//334
//Return true if there exists i, j, k
//such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
//Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.
func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l <= 2{
		return false
	}
	var min int = nums[0]
	var mid int = math.MaxInt32
	for i := 1;i < len(nums);i++ {
		if nums[i] > mid{
			return true
		}
		if nums[i] > min{
			mid = nums[i]
		}else{
			min = nums[i]
		}
	}
	return false;
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

//1370
//Input: s = "aaaabbbbcccc"
//Output: "abccbaabccba"
func sortString(s string) string {
	var record [26]int
	for _,c := range s{
		record[c-'a']++
	}
	var head_to_tail bool = true
	var res string
	head,tail := 0,25;
	for head <= tail {
		if record[head] == 0{
			head++
			continue
		}
		if record[tail] == 0{
			tail--
			continue
		}
		i := head
		j := tail
		for i <= j{
			if head_to_tail{
				if record[i] == 0{
					i++
					continue
				}
			}else{
				if record[j] == 0{
					j--
					continue
				}
			}
			if head_to_tail{
				res += string(i + 'a')
				record[i]--
				i++
			}else{
				res += string(j + 'a')
				record[j]--
				j--
			}
		}
		head_to_tail = !head_to_tail
	}
	return res
}

//1374
func generateTheString(n int) string {
	var res string
	if n % 2 == 1{
		for i := 0;i < n;i++{
			res += "a"
		}
		return res
	}else{
		for i := 0;i < n - 1;i++{
			res += "a"
		}
		res += "b"
		return res
	}
}

//23
/**
 * Definition for singly-linked list_queue.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge_sort(list1 *ListNode,list2 *ListNode)*ListNode{
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}
	var head *ListNode = nil
	var visit *ListNode = nil
	for list1 != nil && list2 != nil{
		if list1.Val < list2.Val{
			if head == nil{
				head = list1
				visit = list1
			}else{
				visit.Next = list1
				visit = list1
			}
			list1 = list1.Next
		}else{
			if head == nil{
				head = list2
				visit = list2
			}else{
				visit.Next = list2
				visit = list2
			}
			list2 = list2.Next
		}
	}
	for list1 != nil{
		visit.Next = list1
		visit = list1
		list1 = list1.Next
	}
	for list2 != nil{
		visit.Next = list2
		visit = list2
		list2 = list2.Next
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	var prev_lists []*ListNode = make([]*ListNode,len(lists))
	copy(prev_lists,lists)
	for len(prev_lists) > 1{
		var lists_merge []*ListNode
		l := len(prev_lists)
		for i := 0;i < l - 1;i += 2{
			new_list := merge_sort(prev_lists[i],prev_lists[i+1])
			lists_merge = append(lists_merge, new_list)
		}
		if l % 2 == 1{
			lists_merge = append(lists_merge,prev_lists[l - 1])
		}
		prev_lists = lists_merge
	}
	return prev_lists[0]
}

//1282
//Input: groupSizes = [3,3,3,3,3,1,3]
//Output: [[5],[0,1,2],[3,4,6]]
//Explanation:
//Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
func groupThePeople(groupSizes []int) [][]int {
	//save the array index by size
	var record map[int][] int = make(map[int][]int)
	for i,s := range groupSizes{
		if _,ok := record[s];ok{
			record[s] = append(record[s], i)
		}else{
			record[s] = []int{i}
		}
	}
	var res [][]int
	for length,nums := range record{
		var collection []int
		for i := 0;i < len(nums);i++{
			collection = append(collection, nums[i])
			if i % length == length - 1{
				res = append(res, collection)
				collection = []int{}
			}
		}
	}
	return res
}

//1329
func diagonalSort(mat [][]int) [][]int {
	rows := len(mat)
	if rows <= 1{
		return mat
	}
	columns := len(mat[0])
	if columns <= 1{
		return mat
	}
	for r := rows - 1;r >= 0;r--{
		for i := 0;i < r;i++{
			for j := 0;j < columns - 1;j++{
				if mat[i][j] > mat[i + 1][j + 1]{
					mat[i][j], mat[i + 1][j + 1] = mat[i + 1][j + 1],mat[i][j]
				}
			}
		}
	}
	return mat
}

//
func luckyNumbers (matrix [][]int) []int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var res []int
	for i := 0;i < rows;i++{
		var min_num_column int = 0
		var cur_row_min int = matrix[i][0]
		for j := 0;j < columns;j++{
			if matrix[i][j] < cur_row_min{
				cur_row_min = matrix[i][j]
				min_num_column = j
			}
		}

		var pass bool = true
		for j := 0;j < rows;j++{
			if j == i{
				continue
			}
			if matrix[j][min_num_column] > matrix[i][min_num_column]{
				pass = false
				break
			}
		}
		if pass{
			res = append(res, matrix[i][min_num_column])
		}
	}
	return res
}

//
type CustomStack struct {
	data []int
	max_size int
	cur_size int
}


func Constructor(maxSize int) CustomStack {
	var obj CustomStack
	obj.max_size = maxSize
	return obj
}

func (this *CustomStack) Push(x int)  {
	if this.cur_size < this.max_size{
		this.data = append(this.data,x)
		this.cur_size++
	}
}


func (this *CustomStack) Pop() int {
	if this.cur_size > 0{
		res := this.data[this.cur_size - 1]
		this.cur_size--
		this.data = this.data[:this.cur_size]
		return res
	}
	return -1
}


func (this *CustomStack) Increment(k int, val int)  {
	for i := 0;i < k && i < this.cur_size;i++{
		this.data[i] += val
	}
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

//3
func lengthOfLongestSubstring(s string) int {
	var start int = 0
	var end int = 0
	var max int = 0
	var cur_len int = 0
	var record map[uint8]int = make(map[uint8]int)
	for end < len(s) && start < len(s){
		if _,ok := record[s[end]];ok{
			start = max_int(record[s[end]] + 1,start)
			record[s[end]] = end
			if cur_len > max{
				max = cur_len
			}
			cur_len = end - start + 1
			end++
		}else{
			cur_len++
			record[s[end]] = end
			end++
		}
	}
	if (end - start) > max{
		max = (end - start)
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

//969
//Input: [3,2,4,1]
//Output: [4,2,4,3]
//Explanation:
//We perform 4 pancake flips, with k values 4, 2, 4, and 3.
//Starting state: A = [3, 2, 4, 1]
//After 1st flip (k=4): A = [1, 4, 2, 3]
//After 2nd flip (k=2): A = [4, 1, 2, 3]
//After 3rd flip (k=4): A = [3, 2, 1, 4]
//After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted.
func reverse_pancakeSort(A []int,begin int,end int){
	for begin < end{
		A[begin],A[end] = A[end],A[begin]
		begin++
		end--
	}
}

func pancakeSort(A []int) []int {
	l := len(A)
	var res []int
	for i := l - 1;i >= 0;i--{
		for j := 0;j <= i;j++{
			if A[j] == i + 1{
				if j != 0{
					res = append(res,j + 1)
					reverse_pancakeSort(A,0,j)
				}
				if i != 0{
					reverse_pancakeSort(A,0,i)
					res = append(res,i + 1)
				}
				break
			}
		}
	}
	return res
}

//1343
//Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
//Output: 3
//Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively.
//All other sub-arrays of size 3 have averages less than 4 (the threshold).
func numOfSubarrays(arr []int, k int, threshold int) int {
	l := len(arr)
	if l < k {
		return 0
	}
	var target int = threshold * k
	var sum int = 0
	for i := 0;i < k;i++{
		sum += arr[i]
	}
	var res int = 0
	begin := 0
	end := begin + k - 1
	for end < l{
		if sum >= target{
			res++
		}
		end++
		if end < l{
			sum -= arr[begin]
			sum += arr[end]
		}
		begin++
	}
	return res
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