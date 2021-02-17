package array

func GenerateMatrix(n int) [][]int {
	var res [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		res[i] = make([]int,n)
	}
	var dir [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	var dir_index = 0
	var num int = 1
	var steps int = n
	var x int = 0
	var y int = 0
	for num <= n * n{
		for i := 0;i < steps;i++{
			res[x][y] = num
			num++
			x += dir[dir_index][0]
			y += dir[dir_index][1]
		}
		if dir_index == 0 || dir_index == 2{
			steps--
		}
		x -= dir[dir_index][0]
		y -= dir[dir_index][1]
		dir_index = (dir_index + 1) % 4
		x += dir[dir_index][0]
		y += dir[dir_index][1]
	}
	return res
}