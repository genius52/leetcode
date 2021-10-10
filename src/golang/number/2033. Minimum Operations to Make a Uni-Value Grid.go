package number

import "sort"

func MinOperations(grid [][]int, x int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	mod := grid[0][0] % x
	var l int = rows * columns
	if l == 1{
		return 0
	}
	var record []int = make([]int,l)
	var idx int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			cur_mod := grid[i][j] % x
			if cur_mod != mod{
				return -1
			}
			record[idx] = grid[i][j]
			idx++
		}
	}
	sort.Ints(record)
	var steps int = 0
	left_idx := 0
	mid := record[l/2]
	for left_idx < l/2{
		steps += (mid - record[left_idx]) / x
		left_idx++
	}
	right_idx := l - 1
	for right_idx > l/2{
		steps += (record[right_idx] - mid) /x
		right_idx--
	}
	return steps
}