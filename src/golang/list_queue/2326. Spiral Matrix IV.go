package list_queue

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	var res [][]int = make([][]int,m)
	for i := 0;i < m;i++{
		res[i] = make([]int,n)
		for j := 0;j < n;j++{
			res[i][j] = -1
		}
	}
	var dirs [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	var idx int = 0
	var up int = 0
	var down int = m - 1
	var left int = 0
	var right int = n - 1
	var r int = 0
	var c int = 0
	var cnt int = m * n
	for cnt >= 0 && head != nil{
		for r >= up && r <= down && c >= left && c <= right && head != nil{
			res[r][c] = head.Val
			head = head.Next
			r += dirs[idx][0]
			c += dirs[idx][1]
			cnt--
		}
		r -= dirs[idx][0]
		c -= dirs[idx][1]
		if idx == 0{
			up++
		}else if idx == 1{
			right--
		}else if idx == 2{
			down--
		}else if idx == 3{
			left++
		}
		idx = (idx + 1) % 4
		r += dirs[idx][0]
		c += dirs[idx][1]
	}
	return res
}