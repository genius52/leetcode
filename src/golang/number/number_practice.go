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

//1108
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1 )
}

//1122
// Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// Output: [2,2,2,1,4,3,3,9,6,7,19]
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var res []int
	for _,val2 := range arr2{		for _,val1 := range arr1{
		if val1 == val2{
			res = append(res, val1)
		}
	}
	}
	var rest []int
	for _,val1 := range arr1{
		var find bool = false
		for _,val2 := range arr2{
			if val1 == val2{
				find = true
				break
			}
		}
		if !find {
			rest = append(rest,val1)
		}
	}
	sort.Ints(rest)
	res = append(res,rest...)
	return res
}

//Input: s1 = "ab" s2 = "eidbaooo"
//Output: True
//Explanation: s2 contains one permutation of s1 ("ba").

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

//1323
//Input: num = 9669
//Output: 9969
func maximum69Number (num int) int {
	var rest int= 0
	if 1000 < num && num < 10000{
		if (num / 1000) == 6{
			return num + 3000
		}
		rest = num % 1000
	}else if 100 < num && num < 1000{
		if (num /100) == 6 {
			return num + 300
		}
		rest = num % 100
	}else if 10 < num && num < 100{
		if (num / 10 ) == 6{
			return num + 30
		}
		rest = num % 10
	}else{
		if num == 6{
			return 9
		}
		return num
	}
	return num - rest + maximum69Number(rest)
}

func check_zero(num int)bool{
	for num != 0{
		if (num % 10) == 0{
			return false
		}
		num = num / 10
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	low := 1
	high := n - 1

	for low <= high{
		if check_zero(low) && check_zero(high){
			return []int{low,high}
		}
		low++
		high--
	}
	return []int{0,0}
}

//1137
func dp_tribonacci(n int,record map[int]int)int {
	if n == 1 || n == 2{
		return 1
	}
	if n <= 0{
		return 0
	}
	if _,ok := record[n];ok{
		return record[n]
	}
	sum := 0
	for i := 1;i <= 3 ;i++{
		sum += tribonacci(n - i)
	}
	record[n] = sum
	return sum
}

func tribonacci(n int) int {
	if n == 1 || n == 2{
		return 1
	}
	if n <= 0{
		return 0
	}
	n1 := 0
	n2 := 1
	n3 := 1
	res := 0
	for i := 3;i <= n;i++{
		res = n3 + n2 + n1
		n1 = n2
		n2 = n3
		n3 = res
	}
	return res
}


//1331
func arrayRankTransform(arr []int) []int {
	l := len(arr)
	var res []int = make([]int,l)
	var record map[int][]int = make(map[int][]int)
	for i,v := range arr{
		if _,ok := record[v];ok{
			record[v] = append(record[v], i)
		}else{
			record[v] = make([]int,1)
			record[v][0] = i
		}
	}
	sort.Ints(arr)
	rank := 1
	for i := 0;i < l;i++{
		if i > 0 && arr[i] == arr[i-1]{
			continue
		}
		cnt := 1
		for _,pos := range record[arr[i]]{
			res[pos] = rank
		}
		rank += cnt
	}
	return res
}

//38
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//func countAndSay(n int) string {
//	if n <= 1{
//		return "1"
//	}
//	var last string = "1"
//	for i := 2;i <= n;i++{
//		slow := 0
//		fast := slow + 1
//		var cur string
//		l := len(last)
//		for fast <= l{
//			if fast < len(last) && last[slow] == last[fast]{
//				fast++
//			}else{
//				l := fast - slow
//				cnt := strconv.Itoa(l)
//				num := string(last[slow])
//				cur += cnt + num
//				slow = fast
//				fast++
//			}
//		}
//		last = cur
//		if i == n{
//			return last
//		}
//	}
//	return last
//}

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


//1346
//Input: arr = [10,2,5,3]
//Output: true
func checkIfExist(arr []int) bool {
	var record map[int]bool = make(map[int]bool)
	for _,n := range arr{
		if n % 2 == 0{
			small := n / 2
			if _,ok := record[small];ok{
				return true
			}
		}
		big := n * 2
		if _,ok := record[big];ok{
			return true
		}
		record[n] = true
	}
	return false
}

//1342
func numberOfSteps(num int) int{
	if num <= 1{
		return num
	}
	var rest int = num % 2
	var step int = 0
	if rest == 1{
		step += 1
	}
	return step + 1 + numberOfSteps(num >> 1)
}

//1362
//Input: num = 8
//Output: [3,3]
//Input: num = 123
//Output: [5,25]
//func dp_closestDivisors(target int,low int,high int) (int,int){
//	if low > high{
//		return 1,target
//	}
//	product := low * high
//	if product == target{
//		l,h := dp_closestDivisors(target,low + 1,high - 1)
//		if (h - l) < (high - low){
//			return l,h
//		}else{
//			return low,high
//		}
//	}else if product > target{
//		new_high := math.Ceil(float64(high * target)/(float64)(product))
//		return dp_closestDivisors(target,low,int(new_high))
//	}else{
//		return dp_closestDivisors(target,low + 1,high)
//	}
//}
//
//func closestDivisors(num int) []int {
//	l1,h1 := dp_closestDivisors(num + 1,1,num + 1)
//	l2,h2 := dp_closestDivisors(num + 2,1,num + 2)
//	if (h1 - l1) < (h2 - l2){
//		return []int{l1,h1}
//	}else{
//		return []int{l2,h2}
//	}
//}

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

//1338
//Input: arr = [3,3,3,3,5,5,5,2,2,7]
//Output: 2
func minSetSize(arr []int) int {
	var record map[int]int = make(map[int]int)
	for _,n := range arr{
		if _,ok := record[n];ok{
			record[n]++
		}else{
			record[n] = 1
		}
	}
	var record_lengths []int
	for _, l := range record{
		record_lengths = append(record_lengths, l)
	}
	sort.Ints(record_lengths)
	target := int(math.Ceil(float64(len(arr))/2))
	total := 0
	for i := 0;i < len(record_lengths);i++{
		total += record_lengths[len(record_lengths) - i - 1]
		if total >= target{
			return i + 1
		}
	}
	return len(record_lengths)
}

//1347
//Input: s = "leetcode", t = "practice"
//Output: 5
//Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.
func minSteps(s string, t string) int {
	var record map[int32]int = make(map[int32]int)
	for _,c := range s{
		if _,ok := record[c];ok{
			record[c]++
		}else{
			record[c] = 1
		}
	}
	for _,c := range t{
		if _,ok := record[c];ok{
			record[c]--
			if record[c] == 0{
				delete(record,c)
			}
		}
	}
	var res int = 0
	for _,l := range record{
		res += l
	}
	return res
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

//739
//For example, given the list_queue of temperatures
//T = [73, 74, 75, 71, 69, 72, 76, 73],
//your output should be [1, 1, 4, 2, 1, 1, 0, 0].
func dailyTemperatures(T []int) []int {
	l := len(T)
	var s []int
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		if len(s) == 0{
			s = append(s,i)
			continue
		}
		for len(s) > 0 && T[i] > T[s[len(s) - 1]]{
			res[s[len(s) - 1]] = i - s[len(s) - 1]
			s = s[0:len(s) - 1]
		}
		s = append(s,i)
	}
	return res
}

//1385
//Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
//Output: 2
//Explanation:
//For arr1[0]=4 we have:
//|4-10|=6 > d=2
//|4-9|=5 > d=2
//|4-1|=3 > d=2
//|4-8|=4 > d=2
//For arr1[1]=5 we have:
//|5-10|=5 > d=2
//|5-9|=4 > d=2
//|5-1|=4 > d=2
//|5-8|=3 > d=2
//For arr1[2]=8 we have:
//|8-10|=2 <= d=2
//|8-9|=1 <= d=2
//|8-1|=7 > d=2
//|8-8|=0 <= d=2
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int = 0
	for i := 0;i < len(arr1);i++{
		var pass bool = true
		for j:= 0;j < len(arr2);j++{
			if int(math.Abs(float64(arr1[i] - arr2[j]))) <= d{
				pass = false
				break
			}
		}
		if pass{
			res++
		}
	}
	return res
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

//xiaoxiaole
//	1,2,3,3,3,2,2
//func xiaoxiaole(nums []int)[]int{
//	l := len(nums)
//	var res []int
//	for i := 0;i < l;i++{
//
//	}
//}


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

//491
func dfs_findSubsequences(nums []int,cur_pos int,cur_nums *[]int,record map[string]bool){
	l := len(*cur_nums)
	if l >= 2{
		var data string
		for _,n := range *cur_nums{
			if len(data) != 0{
				data += ","
			}
			data += strconv.Itoa(n)
		}
		record[data] = true
	}

	if cur_pos >= len(nums){
		return
	}
	if l > 0 {
		if nums[cur_pos] >= (*cur_nums)[l - 1]{
			var add_cur_nums []int
			add_cur_nums  = make([]int,l)
			copy(add_cur_nums,*cur_nums)
			add_cur_nums = append(add_cur_nums,nums[cur_pos])
			dfs_findSubsequences(nums,cur_pos+1,&add_cur_nums,record)
		}
		dfs_findSubsequences(nums,cur_pos+1,cur_nums,record)
	}else{
		var start_cur_nums []int
		start_cur_nums = append(start_cur_nums,nums[cur_pos])
		dfs_findSubsequences(nums,cur_pos+1,&start_cur_nums,record)
		var ignore_cur_nums []int
		dfs_findSubsequences(nums,cur_pos+1,&ignore_cur_nums,record)
	}
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	if len(nums) <= 0{
		return res
	}
	var cur_nums []int
	var record map[string]bool = make(map[string]bool)
	dfs_findSubsequences(nums,0,&cur_nums,record)
	for  r,_ := range record{
		var data_list []string = strings.Split(r,",")
		var tmp []int
		for i := 0;i < len(data_list);i++{
			d , err := strconv.ParseInt(data_list[i],10,64)
			if err == nil{
				tmp = append(tmp, int(d))
			}
		}
		res = append(res, tmp)
	}
	return res
}

//1332
//Input: s = "baabb"
//Output: 2
//Explanation: "baabb" -> "b" -> "".
//Remove palindromic subsequence "baab" then "b".
func is_palindromic(s string)bool{
	l := len(s)
	begin := 0
	end := l - 1
	for begin < end{
		if s[begin] != s[end]{
			return false;
		}
	}
	return true
}
func removePalindromeSub(s string) int {
	if len(s) == 0{
		return 0
	}
	if is_palindromic(s){
		return 1
	}
	return 2
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