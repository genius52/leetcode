package array

import "sort"

func GetBiggestThree(grid [][]int) []int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var record map[int]bool = make(map[int]bool)
	//center point
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			record[grid[i][j]] = true
			var max_edge int = min_int_number(rows-1-i, i, j, columns-1-j)
			for l := 1; l <= max_edge; l++ { //半径
				var sum int = 0
				sum += grid[i-l][j] //上
				sum += grid[i+l][j] //下
				sum += grid[i][j-l] //左
				sum += grid[i][j+l] //右
				for k := 1; k < l; k++ {
					sum += grid[i-l+k][j+k]
					sum += grid[i+l-k][j-k]
					sum += grid[i-k][j-l+k]
					sum += grid[i+k][j+l-k]
				}
				record[sum] = true
			}
		}
	}
	var data []int
	for k, _ := range record {
		data = append(data, k)
	}
	sort.Ints(data)
	var res_len int = len(data)
	var res []int
	if res_len > 3 {
		res = []int{data[res_len-1], data[res_len-2], data[res_len-3]}
	} else {
		for i := res_len - 1; i >= 0; i-- {
			res = append(res, data[i])
		}
	}
	return res
}
