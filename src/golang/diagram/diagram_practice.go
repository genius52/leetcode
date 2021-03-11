package diagram

import (
	"math"
)
func min_int_number(nums ...int)int{
	var min int = math.MaxInt32
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

//85
//Input:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//Output: 6
const VERTICAL = 0
const HORIZONTAL = 1
func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0{
		return 0
	}
	columns := len(matrix[0])
	var dp [][][]int = make([][][]int,rows)
	for i:= 0;i < rows;i++{
		dp[i] = make([][]int,columns)
		for j := 0;j < columns;j++{
			dp[i][j] = make([]int,2)
		}
	}
	//0: vertical direction 垂直方向
	//1: horizon direction 水平方向
	if matrix[0][0] == '0'{
		dp[0][0][VERTICAL] = 0
		dp[0][0][HORIZONTAL] = 0
	}else{
		dp[0][0][VERTICAL] = 1
		dp[0][0][HORIZONTAL] = 1
	}
	for i := 1;i < rows;i++{
		dp[i][0][HORIZONTAL] = 0
		if matrix[i][0] == '0'{
			dp[i][0][VERTICAL] = 0
		}else{
			dp[i][0][VERTICAL] = 1 + dp[i - 1][0][VERTICAL]
		}
	}
	for j := 1;j < columns;j++{
		dp[0][j][VERTICAL] = 0
		if matrix[0][j] == '0'{
			dp[0][j][HORIZONTAL] = 0
		}else{
			dp[0][j][HORIZONTAL] = 1 + dp[0][j - 1][HORIZONTAL]
		}
	}
	var max int = 0
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if matrix[i][j] == '0'{
				continue
			}
			if (dp[i-1][j][HORIZONTAL] == 0 || dp[i-1][j][VERTICAL] == 0) && (dp[i][j-1][HORIZONTAL] == 0 ||dp[i][j-1][VERTICAL] == 0){
				dp[i][j][VERTICAL] = 1
				dp[i][j][HORIZONTAL] = 1
				if max < 1{
					max = 1
				}
				continue
			}
			if dp[i-1][j-1][HORIZONTAL] == 0 || dp[i-1][j-1][VERTICAL] == 0 {
				dp[i][j][VERTICAL] = 1 + dp[i][j - 1][VERTICAL]
				dp[i][j][HORIZONTAL] = 1 + dp[i - 1][j][HORIZONTAL]
			}else{
					dp[i][j][VERTICAL] = 1 + dp[i][j-1][VERTICAL]
				dp[i][j][HORIZONTAL] = 1 + dp[i-1][j][VERTICAL]
			}

			//else if dp[i - 1][j] != 0 && dp[i][j-1] != 0{
			//	dp[i][j] = 1 + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
			//}else{
			//	dp[i][j] = max_int(dp[i-1][j],dp[i][j-1]) - dp[i-1][j-1] + 1
			//}
			//if dp[i][j] > max{
			//	max = dp[i][j]
			//}
		}
	}
	//fmt.Println(dp)
	return max
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

//980
//1 represents the starting square.  There is exactly one starting square.
//2 represents the ending square.  There is exactly one ending square.
//0 represents empty squares we can walk over.
//-1 represents obstacles that we cannot walk over.
func dfs_uniquePathsIII(grid [][]int,row int,column int,visited [][]int, total *int){
	rows := len(grid)
	columns := len(grid[0])
	if row < 0 || row >= rows || column < 0 || column >= columns || visited[row][column] != 0{
		return
	}
	if grid[row][column] == 2{
		for i := 0;i < rows;i++{
			for j := 0;j < columns;j++{
				if i == row && j == column{
					continue
				}
				if visited[i][j] == 0{
					return
				}
			}
		}
		*total++
		return
	}
	if grid[row][column] == -1{
		visited[row][column] = -1
		return
	}
	visited[row][column] = 1
	dfs_uniquePathsIII(grid,row - 1,column,visited,total)
	dfs_uniquePathsIII(grid,row,column - 1,visited,total)
	dfs_uniquePathsIII(grid,row + 1,column,visited,total)
	dfs_uniquePathsIII(grid,row,column + 1,visited,total)
	visited[row][column] = 0
}

func reset_visit(visited [][]int){
	for i := 0;i < len(visited);i++{
		for j := 0;j < len(visited[0]);j++ {
			if visited[i][j] == 1{
				visited[i][j] = 0
			}
		}
	}
}

func uniquePathsIII(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	var visited [][]int = make([][]int,rows)
	for ii := 0;ii < rows;ii++{
		visited[ii] = make([]int,columns)
	}
	var total int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] != 1{
				continue
			}
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i - 1,j,visited,&total)
			reset_visit(visited)
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i + 1,j,visited,&total)
			reset_visit(visited)
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i,j - 1,visited,&total)
			reset_visit(visited)
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i,j + 1,visited,&total)
			break
		}
	}
	return total
}

//1392
//Input: s = "ababab"
//Output: "abab"
func longestPrefix(s string) string {
	l := len(s)
	var mod int = 10e9 + 7
	var res string
	var head int = 0
	var tail int = 0
	var times int = 1
	for i := 0;i < l - 1;i++{
		var s1 int = int((s[i] - 'a') * 26)
		var s2 int = int((s[l - 1 - i] - 'a') * 26)
		head = (head * 26 + s1) % mod
		tail = (s2 * times + tail) % mod
		times = times * 26 % mod
		if head == tail{
			res = s[:i+1]
		}
	}
	return res
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