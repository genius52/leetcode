package number

import (
	"bytes"
	"container/list"
	"math"
	"sort"
	"strconv"
	"strings"
)

func min_int(a,b int)int{
	if a < b{
		return a
	}
	return b
}

func max_int(a,b int)int{
	if a > b {
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

func canThreePartsEqualSum(A []int) bool {
	var total int = 0
	for _, val := range A{
		total += val
	}
	if total % 3 != 0{
		return false
	}
	var target int = total/3
	var sum int = 0
	var find_cnt = 0
	for _,val := range A{
		sum += val
		if(sum == target){
			find_cnt++
			if(find_cnt == 3){
				return true
			}
			sum = 0
		}
	}
	return false;
}

func perm(s string,begin int,end int,target string) bool {
	if(begin > end){
		if strings.Contains(target,s){
			return true
		}
		return false
	}
	c := []byte(s)
	for i := begin;i <= end; i++{
		tmp := c[i]
		c[i] = c[begin]
		c[begin] = tmp
		s2 := string(c)
		if perm(s2,begin + 1,end,target){
			return true
		}
	}
	return false
}

//Input: s1 = "ab" s2 = "eidbaooo"
//Output: True
//Explanation: s2 contains one permutation of s1 ("ba").
func checkInclusion(s1 string, s2 string) bool {
	var len1 int = len(s1)
	var len2 int = len(s2)
	if len1 > len2{
		return false
	}
	return perm(s1,0,len1 - 1,s2)
}

//39
//Input: candidates = [2,3,6,7], target = 7,
//A solution set is:
//[
//  [7],
//  [2,2,3]
//]
func dfs_sum(candidates []int,cur_pos int,cur_sum int,target int,cur_nums []int,res *[][]int){
	if cur_sum > target{
		return
	}
	if cur_sum == target{
		var b = make([]int, len(cur_nums))
		copy(b, cur_nums)
		*res = append(*res,b)
		return
	}
	if cur_pos >= len(candidates){
		return
	}
	//add current element
	cur_nums = append(cur_nums,candidates[cur_pos])
	dfs_sum(candidates,cur_pos,cur_sum + candidates[cur_pos],target,cur_nums,res)
	cur_nums = cur_nums[:len(cur_nums)-1]
	//ignore current element
	dfs_sum(candidates,cur_pos+1,cur_sum,target,cur_nums,res)
}

func combinationSum(candidates []int, target int) [][]int {
	for i := 0;i < len(candidates);i++{
		for j := 1;j < len(candidates) - i;j++{
			if candidates[j] < candidates[j-1]{
				candidates[j] , candidates[j-1] = candidates[j-1],candidates[j]
			}
		}
	}
	var res [][]int
	var cur_nums []int
	dfs_sum(candidates,0,0,target,cur_nums,&res)
	return res
}

//38
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
func countAndSay(n int) string {
	if n <= 1{
		return "1"
	}
	var last string = "1"
	for i := 2;i <= n;i++{
		slow := 0
		fast := slow + 1
		//var cur string
		var cur bytes.Buffer
		l := len(last)
		for fast <= l{
			if fast < len(last) && last[slow] == last[fast]{
				fast++
			}else{
				//l := fast - slow
				//cnt := strconv.Itoa(l)
				//num := string(last[slow])
				//cur += cnt + num
				cur.WriteString(strconv.Itoa(fast - slow))
				cur.WriteByte(last[slow])
				slow = fast
				fast++
			}
		}
		last = cur.String()
		if i == n{
			return last
		}
	}
	return last
}

//1362
//Input: num = 8
//Output: [3,3]
//Input: num = 123
//Output: [5,25]
func dp_closestDivisors(num int,low int,high int)(int,int){
	if low <= 0 || high > num{
		return 1,high
	}
	product := low * high
	if product == num + 1 || product == num + 2{
		return low,high
	}else if product > num{
		return dp_closestDivisors(num,low - 1,high)
	}else{
		return dp_closestDivisors(num,low,high + 1)
	}
}

func closestDivisors(num int) []int {
	n := int(math.Ceil(math.Sqrt(float64(num))))
	//l,h := dp_closestDivisors(num,n,n)
	low := n
	high := n
	for low > 0 && high <= num + 2{
		if low <= 0 || high > num{
			return []int{1,high}
		}
		product := low * high
		if product == num + 1 || product == num + 2{
			return []int{low,high}
		}else if product > num{
			low--
		}else{
			high++
		}
	}
	return []int{1,num}
}

//1314
func matrixBlockSum(mat [][]int, K int) [][]int {
	rows := len(mat)
	columns := len(mat[0])
	var sum [][]int = make([][]int,rows)//sum[i][j] means sum numbers from "0,0" to "i,j"
	for i := 0;i < rows;i++{
		sum[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			upper_sum := 0
			if i - 1 >= 0{
				upper_sum =  sum[i - 1][j]
			}
			left_sum := 0
			if j - 1 >= 0{
				left_sum = sum[i][j - 1]
			}
			dup_sum := 0
			if j - 1 >= 0 && i - 1 >= 0{
				dup_sum = sum[i - 1][j - 1]
			}
			sum[i][j] = upper_sum + left_sum - dup_sum + mat[i][j]
		}
	}
	var res [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
		for j := 0;j < columns;j++ {
			var top_left_region int = 0
			var left_region int = 0
			var top_region int = 0
			var top int = i - K
			var left int = j - K
			if top - 1 >= 0 && left - 1 >= 0{
				top_region = 0
				top_left_region = sum[top - 1][left - 1]
			}
			var bottom int = i + K
			var right int = j + K
			if bottom >= rows{
				bottom = rows - 1
			}
			if right >= columns{
				right = columns - 1
			}
			if top - 1 >= 0{
				top_region = sum[top - 1][right]
			}
			if left - 1 >= 0{
				left_region = sum[bottom][left - 1]
			}
			res[i][j] = sum[bottom][right] - left_region - top_region + top_left_region
		}
	}
	return res
}

//315
//Input: [5,2,6,1]
//Output: [2,1,1,0]
func countSmaller(nums []int) []int {
	l := len(nums)
	var dp []int = make([]int,l)
	dp[l-1] = 0
	for i := 1;i < l;i++{
		var max int = 0
		cur_num := nums[l - 1 - i]
		for j := l - i;j < l;j++{
			if cur_num > nums[j]{
				if dp[j] >= max{
					max = dp[j] + 1
				}
			}
		}
		dp[l - 1 - i] = max
	}
	return dp
}

//84
//Input: [2,1,5,6,2,3]
//Output: 10
////[4,2,0,3,2,5]
func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 1{
		return heights[0]
	}
	var s []int
	var max int = 0
	heights = append(heights, 0)
	l++
	for i := 0;i < l;i++{
		if len(s) == 0 || heights[i] > heights[s[len(s) - 1]]{
			s = append(s,i)
			continue
		}
		for len(s) > 0 && heights[i] < heights[s[len(s) - 1]] {
			h := heights[s[len(s) - 1]]
			s = s[:len(s) - 1]
			var res int = 0
			if len(s) == 0{
				res = h * i
			}else{
				res = h * (i - 1 - s[len(s) - 1])
			}
			if res > max{
				max = res
			}
		}
		s = append(s,i)
	}
	return max
}

//1387
func dfs(num int,record map[int]int)int{
	if num == 1{
		return 0
	}
	if _,ok := record[num];ok{
		return record[num]
	}
	var res int = 0
	if num % 2 == 1{
		res = 1 + dfs(num * 3 + 1,record)
	}else{
		res = 1 + dfs(num/2,record)
	}
	record[num] = res
	return res;
}

func getKth(lo int, hi int, k int) int {
	var record map[int][]int = make(map[int][]int,hi - lo + 1)
	var dp map[int]int = make(map[int]int)
	dp[1] = 0
	for i := lo;i <= hi;i++{
		dfs(i,dp)
	}
	for i,v := range dp{
		if i >= lo && i <= hi{
			if _,ok := record[v];ok{
				record[v] = append(record[v],i)
			}else{
				record[v] = []int{i}
			}
		}
	}
	var keys []int
	for ke := range record {
		keys = append(keys, ke)
	}
	sort.Ints(keys)
	for _, ke := range keys{
		sort.Ints(record[ke])
		for _,n := range record[ke]{
			k--
			if k == 0{
				return n
			}
		}
	}
	return 0
}

//41
//1,2,3
func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0;i < l;i++{
		if nums[i] <= 0 || nums[i] > l{
			nums[i] = 0
		}
	}
	for i := 0;i < l;{
		if nums[i] != 0{
			if nums[i] != nums[nums[i] - 1]{// i 应当等于nums[i] - 1,也就是nums[i] - 1应当等于 i
				nums[nums[i] - 1], nums[i] = nums[i],nums[nums[i] - 1]
			}else{
				i++
			}
		}else{
			i++
		}
	}
	i := 0
	for ;i < l;i++{
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return i + 1
}

//32
//Input: "(()"
//Output: 2
//Explanation: The longest valid parentheses substring is "()"
//Input: ")()())"
//Output: 4
//Explanation: The longest valid parentheses substring is "()()"
func longestValidParentheses(s string) int{
	l := len(s)
	var stack list.List
	var max int = 0
	var start int = 0
	for i := 0;i < l;i++{
		if s[i] == ')' {
			if stack.Len() == 0{
				start = i + 1
			}else{
				top := stack.Back()
				stack.Remove(top)
				if stack.Len() == 0{//)()()()()
					max = max_int(i - start + 1,max)//start will not move,if string is valid
				}else{
					last_left := stack.Back().Value.(int)
					max = max_int(i - last_left,max)//(  () ()  length = i - last_open_ position
				}
			}
		}else{
			stack.PushBack(i)
		}
	}
	return max
}

func dp_longestValidParentheses(s string) int{
	l := len(s)
	if l <= 1{
		return 0
	}
	var dp []int = make([]int,l)
	dp[0] = 0
	var max int = 0
	for i := 1;i < l;i++{
		if s[i] == ')'{
			head := i - dp[i-1] - 1
			if head >= 0 && s[head] == '('{
				dp[i] = dp[i - 1] + 2
				if i - dp[i-1] - 2 > 0{
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > max{
				max = dp[i]
			}
		}
	}
	return max
}

//166
//Input: numerator = 2, denominator = 3
//Output: "0.(6)"
func fractionToDecimal(numerator int, denominator int) string {
	var integer int = numerator / denominator
	var rest = numerator - integer
	if rest == 0 {
		return strconv.Itoa(integer)
	}
	var result string = strconv.FormatFloat(float64(rest)/float64(denominator), 'f', 6, 64)
	if result[len(result) - 1] == '0'{
		return strconv.FormatFloat(float64(numerator)/float64(denominator),'g',-1,64)
	}
	i := len(result) - 3

	for i >= 0{
		if result[i] == result[i + 1]{
			i--
		}else{
			result = result[1:i+1] + "(" + string(result[len(result) - 2]) + ")"
			break
		}
	}
	return strconv.Itoa(integer) + result
}

//6 ZigZag Conversion
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
func convert(s string, numRows int) string {
	l := len(s)
	if numRows <= 1{
		return s
	}
	var arr [][]uint8 = make([][]uint8,numRows)
	for i := 0;i < numRows;i++{
		arr[i] = make([]uint8,l)
	}
	i := 0
	col := 0
	for i < l{
		for row := 0;row < numRows;row++{
			arr[row][col] = s[i]
			i++
			if i >= l{
				goto out
			}
		}
		col++
		for row := numRows - 2;row > 0;row--{
			arr[row][col] = s[i]
			col++
			i++
			if i >= l{
				goto out
			}
		}
	}
out:
	var res string
	for r := 0;r < numRows;r++ {
		for c := 0; c < l; c++ {
			if arr[r][c] != 0 {
				res += string(arr[r][c])
			}
		}
	}
	return res
}