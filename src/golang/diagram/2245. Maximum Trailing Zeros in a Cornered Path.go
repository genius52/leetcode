package diagram

const (
	left_right int = 0
	right_left int = 1
	top_bottom int = 2
	bottom_top int = 3
)

const (
	num2_idx int = 0
	num5_idx int = 1
)

func cal_maxTrailingZeros(n int)(int,int){
	var cnt2 int = 0
	var cnt5 int = 0
	for n % 5 == 0{
		cnt5++
		n /= 5
	}
	for n % 2 == 0{
		cnt2++
		n /= 2
	}
	return cnt2,cnt5
}

func MaxTrailingZeros(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var record [][][2]int = make([][][2]int,rows)
	var dp_two_cnt [][][4]int = make([][][4]int,rows)//四个方向 2的个数
	var dp_five_cnt [][][4]int = make([][][4]int,rows)//四个方向 5的个数
	for i := 0;i < rows;i++{
		record[i] = make([][2]int,columns)
		dp_two_cnt[i] = make([][4]int,columns)
		dp_five_cnt[i] = make([][4]int,columns)
		for j := 0;j < columns;j++{
			record[i][j][num2_idx],record[i][j][num5_idx] = cal_maxTrailingZeros(grid[i][j])
			if i == 0{//第一行
				dp_two_cnt[i][j][top_bottom],dp_five_cnt[i][j][top_bottom] = cal_maxTrailingZeros(grid[i][j])
			}
			if i == rows - 1{//最后一行
				dp_two_cnt[i][j][bottom_top],dp_five_cnt[i][j][bottom_top] = cal_maxTrailingZeros(grid[i][j])
			}
			if j == 0{//第一列
				dp_two_cnt[i][j][left_right],dp_five_cnt[i][j][left_right] = cal_maxTrailingZeros(grid[i][j])
			}
			if j == columns - 1{//最后一列
				dp_two_cnt[i][j][right_left],dp_five_cnt[i][j][right_left] = cal_maxTrailingZeros(grid[i][j])
			}
		}
	}
	//从左到右
	for i := 0;i < rows;i++{
		for j := 1;j < columns;j++{
			dp_two_cnt[i][j][left_right] = record[i][j][0] + dp_two_cnt[i][j - 1][left_right]
			dp_five_cnt[i][j][left_right] = record[i][j][1] + dp_five_cnt[i][j - 1][left_right]
		}
	}
	//从右到左
	for i := 0;i < rows;i++{
		for j := columns - 2;j >= 0;j--{
			dp_two_cnt[i][j][right_left] = record[i][j][0] + dp_two_cnt[i][j + 1][right_left]
			dp_five_cnt[i][j][right_left] = record[i][j][1] + dp_five_cnt[i][j + 1][right_left]
		}
	}
	//从上到下
	for j := 0;j < columns;j++{
		for i := 1;i < rows;i++{
			dp_two_cnt[i][j][top_bottom] = record[i][j][0] + dp_two_cnt[i - 1][j][top_bottom]
			dp_five_cnt[i][j][top_bottom] = record[i][j][1] + dp_five_cnt[i - 1][j][top_bottom]
		}
	}
	//从下到上
	for j := 0;j < columns;j++{
		for i := rows - 2;i >= 0;i--{
			dp_two_cnt[i][j][bottom_top] = record[i][j][0] + dp_two_cnt[i + 1][j][bottom_top]
			dp_five_cnt[i][j][bottom_top] = record[i][j][1] + dp_five_cnt[i + 1][j][bottom_top]
		}
	}
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			res = max_int(res,min_int(dp_two_cnt[i][j][left_right] + dp_two_cnt[i][j][top_bottom] - record[i][j][num2_idx],dp_five_cnt[i][j][left_right] + dp_five_cnt[i][j][top_bottom] - record[i][j][num5_idx]))
			res = max_int(res,min_int(dp_two_cnt[i][j][left_right] + dp_two_cnt[i][j][bottom_top] - record[i][j][num2_idx],dp_five_cnt[i][j][left_right] + dp_five_cnt[i][j][bottom_top] - record[i][j][num5_idx]))
			res = max_int(res,min_int(dp_two_cnt[i][j][right_left] + dp_two_cnt[i][j][top_bottom] - record[i][j][num2_idx],dp_five_cnt[i][j][right_left] + dp_five_cnt[i][j][top_bottom] - record[i][j][num5_idx]))
			res = max_int(res,min_int(dp_two_cnt[i][j][right_left] + dp_two_cnt[i][j][bottom_top] - record[i][j][num2_idx],dp_five_cnt[i][j][right_left] + dp_five_cnt[i][j][bottom_top] - record[i][j][num5_idx]))
		}
	}
	return res
}