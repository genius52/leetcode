package diagram

func MinimumSum(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = rows * columns
	//横切3个
	for r1 := 0; r1 < rows-2; r1++ {
		top := get_minimumArea(grid, 0, columns-1, 0, r1)
		if top <= 0 {
			continue
		}
		for r2 := r1 + 1; r2 < rows-1; r2++ {
			mid := get_minimumArea(grid, 0, columns-1, r1+1, r2)
			down := get_minimumArea(grid, 0, columns-1, r2+1, rows-1)
			if mid > 0 && down > 0 {
				cur := top + mid + down
				if cur < res {
					res = cur
				}
			}
		}
	}
	//竖切3个
	for c1 := 0; c1 < columns-2; c1++ {
		left := get_minimumArea(grid, 0, c1, 0, rows-1)
		if left <= 0 {
			continue
		}
		for c2 := c1 + 1; c2 < columns-1; c2++ {
			mid := get_minimumArea(grid, c1+1, c2, 0, rows-1)
			right := get_minimumArea(grid, c2+1, columns-1, 0, rows-1)
			if mid > 0 && right > 0 {
				cur := left + mid + right
				if cur < res {
					res = cur
				}
			}
		}
	}
	//上面1个，下面左右2个
	for r := 0; r < rows-1; r++ {
		top := get_minimumArea(grid, 0, columns-1, 0, r)
		if top <= 0 {
			continue
		}
		for c := 0; c < columns-1; c++ {
			left := get_minimumArea(grid, 0, c, r+1, rows-1)
			right := get_minimumArea(grid, c+1, columns-1, r+1, rows-1)
			if left > 0 && right > 0 {
				cur := top + left + right
				if cur < res {
					res = cur
				}
			}
		}
	}
	//下面1个，上面左右2个
	for r := 1; r < rows; r++ {
		bottom := get_minimumArea(grid, 0, columns-1, r, rows-1)
		if bottom <= 0 {
			continue
		}
		for c := 0; c < columns-1; c++ {
			left := get_minimumArea(grid, 0, c, 0, r-1)
			right := get_minimumArea(grid, c+1, columns-1, 0, r-1)
			if left > 0 && right > 0 {
				cur := bottom + left + right
				if cur < res {
					res = cur
				}
			}
		}
	}
	//左边1个，右边上下2个
	for c := 0; c < columns-1; c++ {
		left := get_minimumArea(grid, 0, c, 0, rows-1)
		if left <= 0 {
			continue
		}
		for r := 0; r < rows-1; r++ {
			top := get_minimumArea(grid, c+1, columns-1, 0, r)
			bottom := get_minimumArea(grid, c+1, columns-1, r+1, rows-1)
			if top > 0 && bottom > 0 {
				cur := left + top + bottom
				if cur < res {
					res = cur
				}
			}
		}
	}
	//右边1个，左边上下2个
	for c := 1; c < columns; c++ {
		right := get_minimumArea(grid, c, columns-1, 0, rows-1)
		if right <= 0 {
			continue
		}
		for r := 0; r < rows-1; r++ {
			up := get_minimumArea(grid, 0, c-1, 0, r)
			bottom := get_minimumArea(grid, 0, c-1, r+1, rows-1)
			if up > 0 && bottom > 0 {
				cur := right + up + bottom
				if cur < res {
					res = cur
				}
			}
		}
	}
	return res
}

func get_minimumArea(grid [][]int, x1 int, x2 int, y1 int, y2 int) int {
	var up int = y2
	var down int = y1
	var left int = x2
	var right int = x1
	for r := y1; r <= y2; r++ {
		for c := x1; c <= x2; c++ {
			if grid[r][c] == 1 {
				if r < up {
					up = r
				}
				if r > down {
					down = r
				}
				if c < left {
					left = c
				}
				if c > right {
					right = c
				}
			}
		}
	}
	return (right - left + 1) * (down - up + 1)
}
