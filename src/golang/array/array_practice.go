package array

import (
	"container/list"
	"math"
	"strconv"
)

func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

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

func max_int64(a,b int64)int64{
	if a > b {
		return a
	}else{
		return b
	}
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