package array

func onesMinusZeros(grid [][]int) [][]int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var horizon_zero []int = make([]int,rows) //horizon_zero[i]: 第i行0的个数
	var vertical_zero []int = make([]int,columns)//vertical_zero[i]: 第i列0的个数
	var res [][]int = make([][]int,rows)
	//horizon
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			if grid[i][j] == 0{
				horizon_zero[i]++
			}
		}
	}
	//vertical
	for j := 0;j < columns;j++{
		for i := 0;i < rows;i++{
			if grid[i][j] == 0{
				vertical_zero[j]++
			}
		}
	}
	//diff[i][j] = onesRowi + onesColj - zerosRowi - zerosColj
	//令第 i 行一的数目为 onesRowi 。
	//令第 j 列一的数目为 onesColj 。
	//令第 i 行零的数目为 zerosRowi 。
	//令第 j 列零的数目为 zerosColj 。
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			res[i][j] = columns - horizon_zero[i] +  rows - vertical_zero[j] - horizon_zero[i] - vertical_zero[j]
		}
	}
	return res
}