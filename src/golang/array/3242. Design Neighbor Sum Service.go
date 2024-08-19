package array

type neighborSum struct {
	data [][]int
}

func Constructor(grid [][]int) neighborSum {
	var obj neighborSum
	obj.data = grid
	return obj
}

func (this *neighborSum) AdjacentSum(value int) int {
	var dirs [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var sum int = 0
	var rows int = len(this.data)
	var columns int = len(this.data[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if this.data[r][c] == value {
				for _, dir := range dirs {
					var next_r int = r + dir[0]
					var next_c int = c + dir[1]
					if next_r >= 0 && next_r < rows && next_c >= 0 && next_c < columns {
						sum += this.data[next_r][next_c]
					}
				}
				break
			}
		}
	}
	return sum
}

func (this *neighborSum) DiagonalSum(value int) int {
	var dirs [][]int = [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	var sum int = 0
	var rows int = len(this.data)
	var columns int = len(this.data[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if this.data[r][c] == value {
				for _, dir := range dirs {
					var next_r int = r + dir[0]
					var next_c int = c + dir[1]
					if next_r >= 0 && next_r < rows && next_c >= 0 && next_c < columns {
						sum += this.data[next_r][next_c]
					}
				}
				break
			}
		}
	}
	return sum
}

/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
