package diagram

const (
	North = 0
	South = 1
	East = 2
	West = 3
)

var dirs []int = []int{North,South,East,West}
func OrderOfLargestPlusSign(N int, mines [][]int) int {
	var dis [][][]int = make([][][]int,N)
	var grid [][]int = make([][]int,N)
	for i := 0;i < N;i++{
		grid[i] = make([]int,N)
		dis[i] = make([][]int,N)
		for j := 0;j < N;j++{
			grid[i][j] = 1
			dis[i][j] = make([]int,4)
			for k := 0;k < 4;k++{
				dis[i][j][k] = 1
			}
		}
	}
	for _,m := range mines{
		grid[m[0]][m[1]] = 0
		for i := 0;i < 4;i++{
			dis[m[0]][m[1]][i] = 0
		}
	}

	var res int = 0
	//north
	for i := 1;i < N;i++{
		for j := 0;j < N;j++{
			if grid[i][j] == 0{
				continue
			}
			dis[i][j][0] += dis[i - 1][j][0]
		}
	}
	//south
	for i := N - 2;i >= 0;i--{
		for j := 0;j < N;j++{
			if grid[i][j] == 0{
				continue
			}
			dis[i][j][1] += dis[i + 1][j][1]
		}
	}
	//east
	for i := 1;i < N;i++{
		for j := 0;j < N;j++{
			if grid[j][i] == 0{
				continue
			}
			dis[j][i][2] += dis[j][i - 1][2]
		}
	}
	//west
	for i := N - 2;i >= 0;i--{
		for j := 0;j < N;j++{
			if grid[j][i] == 0{
				continue
			}
			dis[j][i][3] += dis[j][i + 1][3]
		}
	}
	for i := 0;i < N;i++{
		for j := 0;j < N;j++{
			d := min_int_number(dis[i][j][0],dis[i][j][1],dis[i][j][2],dis[i][j][3])
			if d > res{
				res = d
			}
		}
	}
	return res
}