package diagram

import (
	"math"
)

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

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

//621
//func leastInterval(tasks []byte, n int) int {
//	var record map[byte]int = make(map[byte]int)
//	for _,c := range tasks{
//		if _,ok := record[c];ok{
//			record[c]++
//		}else{
//			record[c] = 1
//		}
//	}
//	var steps int = 0
//	i := len(tasks)
//	for i > 0{
//		cnt := 0
//		//l := len(record)
//		for t,_ := range record{
//			record[t]--
//			i--
//			if record[t] == 0{
//				delete(record,t)
//			}
//
//			cnt++
//			steps++
//		}
//		if cnt < n + 1{
//			steps += n + 1 - cnt
//		}
//	}
//	return steps
//}

//10
//Input:
//s = "ab"
//p = ".*"
//Output: true
//Input:
//s = "aab"
//p = "c*a*b"
//Output: true
func dp_isMatch(s string,p string,cur_s int,cur_p int,last_p uint8)bool{
	if cur_s >= len(s){
		return true
	}
	if cur_p >= len(p){
		return false
	}

	c := p[cur_p]
	if s[cur_s] == c{
		return dp_isMatch(s,p,cur_s + 1,cur_p + 1,p[cur_p])
	}else if c == '*'{
		if s[cur_s] != last_p{
			return dp_isMatch(s,p,cur_s,cur_p + 1,last_p)
		}else{
			return dp_isMatch(s,p,cur_s + 1,cur_p,last_p)
		}
	}else if c == '.'{
		if cur_p < len(p) - 1 {

		}
		return dp_isMatch(s,p,cur_s + 1,cur_p + 1,s[cur_s])
	}else{
		if cur_p < len(p) - 1{
			if p[cur_p + 1] == '*'{
				return dp_isMatch(s,p,cur_s,cur_p + 2,last_p)
			}
		}
		return false
	}
}

func isMatch(s string, p string) bool {
	return dp_isMatch(s,p,0,0,0)
}

//1284
func minFlips(mat [][]int) int {
	//rows := len(mat)
	//columns := len(mat[0])
	//for i := 0;i < rows;i++{
	//	for j := 0;j < columns;j++{
	//		fmt.Printf("%d行%d列是%d\r\n",i,j,mat[i][j])
	//	}
	//}
	return 0
}