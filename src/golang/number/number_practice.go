package number

import (
	"bytes"
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

func toHex(num int) string {
	//var m []string = make([]string,16,0)
	char_map := "0123456789abcdef"
	var res string
	for num > 0{
		mod := char_map[num % 16]
		res = string(mod) + res
		num = num/16
	}
	return res
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

func isPerfectSquare(num int)bool {
	var i int = 1;
	for num > 0 {
		num -= i;
		i += 2;
	}
	return num == 0;
}

func mySqrt(x int) int {
	if x == 0{
		return 0
	}
	low := 1
	high := x
	for low <= high{
		mid := low + (high-low)/2
		val := mid*mid
		if val == x{
			return mid
		}else if val < x{
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return high
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for _,house_pos := range houses{
 		min_diff := math.MaxInt32
		for _,heater_pos := range heaters{
			diff := math.Abs(float64(house_pos - heater_pos))
			if diff < float64(min_diff){
				min_diff = int(diff)
			}
		}
		if min_diff > res{
			res = min_diff
		}
	}
	return res
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

//Given a positive integer, return its corresponding column title as appear in an Excel sheet.
//
//For example:
//
//    1 -> A
//    2 -> B
//    3 -> C
//    ...

//    26 -> Z
//    27 -> AA
//    28 -> AB
//    ...
//Example 1:
//
//Input: 1
//Output: "A"
//Example 2:
//
//Input: 28
//Output: "AB"
//Example 3:
//
//Input: 701
//Output: "ZY"
func convertToTitle(n int) string {
	var res string
	if n < 1{
		return res
	}

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

//279
func numSquares(n int) int {
	var record []int = make([]int,n + 1)
	record[0] = 0
	for i := 1;i <= n;i++{
		record[i] = math.MaxInt32
		for j := 1;(j * j) <= i;j++{
			record[i] = min_int_number(record[i],record[i - j*j] + 1)
		}
	}
	return record[n]
}

//43
//Input: num1 = "123", num2 = "456"
//Output: "56088"
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0"{
		return "0"
	}
	var long_num string
	var short_num string
	l1 := len(num1)
	l2 := len(num2)
	if l1 >=l2{
		long_num = num1
		short_num = num2
	}else{
		long_num = num2
		short_num = num1
	}
	var tmp [][]int = make([][]int,len(short_num))
	for i := 0;i < len(short_num);i++ {
		tmp[i] = make([]int,len(long_num))
		for j := 0;j < len(long_num);j++ {
			v1,err1 := strconv.Atoi(string(short_num[len(short_num) - 1 - i]))
			v2,err2 := strconv.Atoi(string(long_num[len(long_num) - 1 - j]))
			if nil == err1 && nil == err2{
				tmp[i][j] = v1 * v2
			}
		}
	}
	var bit_sum []int = make([]int,len(short_num) + len(long_num))
	for i := 0;i <(len(long_num) + len(short_num));i++{//i表示 个、十、百、千、万
		for j := 0;j < len(short_num) && (i - j) >=0 ;j++{//j表示
			if i - j >= len(long_num){
				continue
			}
			bit_sum[i] += tmp[j][i - j]
		}
	}
	var res string
	var upgrade int = 0
	if bit_sum[0] >= 10{
		upgrade = bit_sum[0] / 10
		res += strconv.Itoa(bit_sum[0] % 10)
	}else{
		res += strconv.Itoa(bit_sum[0])

	}

	for i := 1;i < len(bit_sum);i++{
		if i == (len(bit_sum)-1) && bit_sum[i] == 0{
			break
		}
		val := bit_sum[i] + upgrade
		res = strconv.Itoa(val % 10) + res
		upgrade = val / 10
	}
	if upgrade > 0{
		res = strconv.Itoa(upgrade) + res
	}
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

//696
//Input: "00110011"
//Output: 6
//Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
func is_binary(s string,start,end int,l int)int{
	if start < 0 || end >= l{
		return 0
	}
	if s[start] == s[end]{
		return 0
	}
	var cnt int = 1
	head_val := s[start]
	tail_val := s[end]
	start--
	end++
	for start >= 0 && end < l{
		if s[start] != head_val || s[end] != tail_val{
			return cnt
		}
		start--
		end++
		cnt++
	}
	return cnt
}

func countBinarySubstrings(s string) int {
	var res int = 0
	l := len(s)
	for i := 0;i < (l - 1);i++{
		res += is_binary(s,i,i+1,l)
	}
	return res
}

//119
//1
//11
//121
//1331
//14641
func getRow(rowIndex int) []int {
	if rowIndex < 0{
		return []int{}
	}
	var triangle [][]int = make([][]int,rowIndex + 1)
	triangle[0] = make([]int,1)
	triangle[0][0] = 1
	level := 1
	for level <= rowIndex{
		triangle[level] = make([]int,level + 1)
		triangle[level][0] = 1
		triangle[level][level] = 1
		pos := 1
		for pos <= level/2{
			triangle[level][pos] = triangle[level-1][pos] + triangle[level-1][pos - 1]
			triangle[level][level - pos] = triangle[level-1][level - pos] + triangle[level-1][level - pos - 1]
			pos++
		}
		level++
	}
	return triangle[level-1]
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


//299
//Input: secret = "1807", guess = "7810"
//Output: "1A3B"
//Input: secret = "1123", guess = "0111"
//Output: "1A1B"
func min_int(a,b int)int{
	if a < b{
		return a
	}
	return b
}
func getHint(secret string, guess string) string {
	var secret_record [10]int
	var guess_record [10]int
	var a_cnt int = 0
	for i,_ := range secret{
		if secret[i] == guess[i]{
			a_cnt++
		}else{
			secret_record[secret[i] - '0']++
			guess_record[guess[i] - '0']++
		}
	}
	var b_cnt int = 0
	for i := 0;i < 10;i++{
		b_cnt += min_int(secret_record[i],guess_record[i])
	}
	res := strconv.Itoa(a_cnt) + "A" + strconv.Itoa(b_cnt) + "B"
	return res
}

//303
//obj := Constructor(nums);
//param_1 := obj.SumRange(i,j);
type NumArray struct {
	data []int
	dp map[string]int
}

func Constructor3(nums []int) NumArray {
	var obj NumArray
	obj.data = make([]int,len(nums))
	copy(obj.data,nums)
	obj.dp = make(map[string]int)
	//sort.Ints(obj.data)
	return obj
}


func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || j >= len(this.data){
		return 0
	}
	var k string = strconv.Itoa(i)+ "," + strconv.Itoa(j)
	if _,ok := this.dp[k];ok{
		return this.dp[k]
	}
	var res int = 0
	for begin := i; begin <= j;begin++{
		res += this.data[begin]
	}
	this.dp[k] = res
	return res
}

//447
func numberOfBoomerangs(points [][]int) int {
	var res int = 0
	for i := 0;i < len(points);i++{
		target := points[i]
		var dis_record map[int]int = make(map[int]int)
		for j := 0;j < len(points);j++{
			if i == j{
				continue
			}
			dis := (target[0] - points[j][0]) * (target[0] - points[j][0]) + (target[1] - points[j][1]) * (target[1] - points[j][1])
			if _,ok := dis_record[dis];ok{
				dis_record[dis]++
			}else{
				dis_record[dis] = 1
			}
		}
		for _,v := range dis_record{
			if v <= 1{
				continue
			}
			res += v * (v - 1)
		}
	}
	return res
}

//494
//Input: nums is [1, 1, 1, 1, 1], S is 3.
//Output: 5
//Explanation:
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
func dp_findTargetSumWays(nums []int,cur_pos int,cur_sum int,S int)int{
	if cur_pos > len(nums){
		return 0
	}
	if cur_pos == len(nums){
		if cur_sum == S{
			return 1
		}else{
			return 0
		}
	}
	cnt := dp_findTargetSumWays(nums,cur_pos + 1,cur_sum + nums[cur_pos],S)
	cnt += dp_findTargetSumWays(nums,cur_pos + 1,cur_sum - nums[cur_pos],S)
	return cnt
}

func findTargetSumWays(nums []int, S int) int {
	return dp_findTargetSumWays(nums,0,0,S)
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

//402
func dp_removeKdigits(num string,cur_num string,target_len int,cur_pos int)int{
	if len(cur_num) == target_len{
		val,_ := strconv.Atoi(cur_num)
		return val
	}
	if cur_pos >= len(num){
		val,_ := strconv.Atoi(num)
		return val
	}
	n1 := dp_removeKdigits(num,cur_num,target_len,cur_pos + 1)
	n2 := dp_removeKdigits(num,cur_num + string(num[cur_pos]),target_len,cur_pos + 1)
	return min_int(n1,n2)
}

func removeKdigits2(num string, k int) string {
	if len(num) <= k{
		return "0"
	}
	res := strconv.Itoa(dp_removeKdigits(num,"",len(num) - k,0))
	return res
}

func removeKdigits(num string, k int) string {
	if len(num) <= k{
		return "0"
	}
	var remove_cnt int = 0
	for remove_cnt < k && len(num) > 0{
		var delete bool = false
		for i := 0;i < len(num) - 1;i++{
			if num[i] > num[i+1]{
				num = num[:i] + num[i+1:]
				delete = true
				break;
			}
		}
		if !delete{
			num = num[:len(num)-1]
		}
		remove_cnt++
	}
	res := strings.TrimLeft(num,"0")
	if len(res) == 0{
		return "0"
	}
	return res
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

//312
//Input: [3,1,5,8]
//Output: 167
//Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
//func dp_maxCoins(nums []int,memo [][]int)int{
//	if len(nums) < 0 {
//		return 0
//	}
//	var res int = 0
//	for mid := 0; mid < len(nums);mid++{
//		var left_num int = 1
//		var right_num int = 1
//		if mid != 0{
//			left_num = nums[mid-1]
//		}
//		if mid != len(nums) - 1{
//			right_num = nums[mid + 1]
//		}
//		cur_val := left_num * nums[mid] * right_num
//		var new_nums []int = make([]int,len(nums))
//		copy(new_nums,nums)
//		new_nums = append(new_nums[:mid],new_nums[mid+1:]...)
//		rest_val := dp_maxCoins(new_nums,memo)
//		sum := rest_val + cur_val
//		if sum > res{
//			res = sum
//		}
//	}
//	return res
//}
//
//func maxCoins(nums []int) int {
//	var record [][]int = make([][]int,len(nums))
//	for i := 0;i < len(nums);i++{
//		record[i] = make([]int,len(nums))
//	}
//	return dp_maxCoins(nums,record)
//}

func maxCoins(nums []int) int {
	new_len := len(nums) + 2
	var nums2 []int = make([]int,new_len)
	for i,n := range nums{
		nums2[i + 1] = n
	}
	nums2[0] = 1
	nums2[new_len - 1] = 1
	var dp [][]int = make([][]int,new_len)	//dp[i][j] means  maximum coins if we burst ballons from postion i to j
	for i := 0;i < new_len;i++{
		dp[i] = make([]int,new_len)
	}
	for span := 1; span <= len(nums);span++{
		for left := 1;left <= len(nums) - span + 1;left++{//ignore left = 0
			right := left + span - 1
			for k := left;k <= right;k++{
				dp[left][right] = max_int(dp[left][right],nums2[left - 1] * nums2[k] * nums2[right + 1] + dp[left][k - 1] + dp[k + 1][right])
			}
		}
	}
	return dp[1][len(nums)]
}

//227
//Implement a basic calculator to evaluate a simple expression string.
//The expression string contains only non-negative integers, +, -, *, / operators and empty spaces .
//The integer division should truncate toward zero.
//Input: " 3+5 / 2 "
//Output: 5
func calculate(s string) int {
	var q list.List
	var cur_num int = 0
	var last_op int32
	s = "+" + s + "+"
	for _,c := range s{
		if c >= '0' && c <= '9' {
			n , _ := strconv.Atoi(string(c))
			cur_num = cur_num * 10 + n
			continue
		}
		if c == ' '{
			continue
		}
		if last_op == '+'{
			q.PushBack(cur_num)
		}else if last_op == '-'{
			q.PushBack(-cur_num)
		}else if last_op == '*'{
			last_num := q.Back().Value.(int)
			q.Remove(q.Back())
			q.PushBack(last_num * cur_num)
		}else if last_op == '/'{
			last_num := q.Back().Value.(int)
			q.Remove(q.Back())
			q.PushBack(last_num/cur_num)
		}
		last_op = c
		cur_num = 0
	}
	var sum int = 0
	for item := q.Front();nil != item ;item = item.Next() {
		n := item.Value.(int)
		sum += n
	}
	return sum
}

//295
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2


//162
func findPeakElement(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	if nums[0] >= nums[1] {
		return 0
	}
	if nums[l - 1] >= nums[l - 2]{
		return l - 1
	}
	low := 0
	high := l - 1
	for low < high{
		mid := (low + high)/2
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1]{
			return mid
		}
		if nums[mid-1] > nums[mid]{
			high = mid
		}else{
			low = mid + 1
		}
	}
	return 0
}

//933
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
type RecentCounter struct {
	record []int
}

func Constructor933() RecentCounter {
	var obj RecentCounter
	return obj
}

func (this *RecentCounter) Ping(t int) int {
	this.record = append(this.record,t)
	low := 0
	l := len(this.record)
	high := l - 1
	target := t - 3000
	mid := 0
	//find the postion of first num large or equal the target
	for low <= high{
		mid = (low + high)/2
		if this.record[mid] == target{
			break
		}
		if this.record[mid] < target{
			low = mid + 1
			if mid < (l - 1) && this.record[mid + 1] >= target{
				return len(this.record) - mid - 1
			}
		}else{
			high = mid - 1
			if mid > 0 && this.record[mid - 1] < target{
				return len(this.record) - mid
			}
		}
	}
	return len(this.record) - mid
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
//For example, given the list of temperatures
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

//20
//Input: 2.00000, -2
//Output: 0.25000
func myPow(x float64, n int) float64 {
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	if n < 0{
		return myPow(1 / x,-n)
	}
	if n == 2{
		return x * x
	}
	if n % 2 == 0{
		return myPow(myPow(x,n/2),2)
	}else{
		return x * myPow(myPow(x,n/2), 2)
	}
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

//959
func dfs_regionsBySlashes(visited [][]bool,i int,j int)bool{
	l := len(visited)
	if i < 0 || j < 0 || i >= l || j >= l || visited[i][j]{
		return false
	}
	visited[i][j] = true
	dfs_regionsBySlashes(visited,i - 1,j)
	dfs_regionsBySlashes(visited,i + 1,j)
	dfs_regionsBySlashes(visited,i,j - 1)
	dfs_regionsBySlashes(visited,i,j + 1)
	return true
}

func regionsBySlashes(grid []string) int {
	l := len(grid)
	var visited [][]bool = make([][]bool,l * 3)
	for i := 0;i < l;i++ {
		visited[i * 3] = make([]bool,l * 3)
		visited[i * 3 + 1] = make([]bool,l * 3)
		visited[i * 3 + 2] = make([]bool,l * 3)
		for j := 0;j < l;j++{
			if grid[i][j] == '\\'{
				visited[i * 3][j * 3] = true
				visited[i * 3 + 1][j * 3 + 1] = true
				visited[i * 3 + 2][j * 3 + 2] = true
			}else if grid[i][j] == '/'{
				visited[i * 3][j * 3 + 2] = true
				visited[i * 3 + 1][j * 3 + 1] = true
				visited[i * 3 + 2][j * 3] = true
			}
		}
	}
	var res int = 0
	for i := 0;i < l * 3;i++{
		for j := 0;j < l * 3;j++{
			if dfs_regionsBySlashes(visited,i,j){
				res++
			}
		}
	}
	return res
}

//53
//[-2,1,-3,4,-1,2,1,-5,4],
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	//var dp []int = make([]int,l)
	//dp[0] = 0
	var last_sum int = 0
	var cur_sum int = 0
	var max int = math.MinInt32
	for i := 0;i < l;i++{
		if i == 0{
			cur_sum = nums[i]
		}else{
			cur_sum = max_int(nums[i], last_sum + nums[i])
		}
		last_sum = cur_sum
		if cur_sum > max{
			max = cur_sum
		}
	}
	return max
}

//918
//Input: [3,-1,2,-1]
//Output: 4
//Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4
//Input: [5,-3,5]
//Output: 10
//Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10
func maxSubarraySumCircular(A []int) int {
	l := len(A)
	var max int = math.MinInt32
	var min int = math.MaxInt32
	var last_sum_max int
	var cur_sum_max int
	var last_sum_min int
	var cur_sum_min int
	var sum int = 0
	for i := 0;i < l;i++{
		sum += A[i]
		if i == 0{
			cur_sum_max = A[i]
			cur_sum_min = A[i]
		}else{
			cur_sum_max = max_int(A[i],last_sum_max + A[i])
			cur_sum_min = min_int(A[i],last_sum_min + A[i])
		}
		if cur_sum_max > max{
			max = cur_sum_max
		}
		if cur_sum_min < min{
			min = cur_sum_min
		}
		last_sum_max = cur_sum_max
		last_sum_min = cur_sum_min
	}
	if sum != min{
		return max_int(max, sum - min)
	}
	return max
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