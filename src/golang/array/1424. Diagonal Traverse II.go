package array

//Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [1,4,2,7,5,3,8,6,9]
func FindDiagonalOrder(nums [][]int) []int {
	rows := len(nums)
	var column_record []int = make([]int,rows)
	var max_column int = 0
	var total int = 0
	for i := 0;i < rows;i++{
		column_record[i] = len(nums[i])
		total += column_record[i]
		if column_record[i] > max_column{
			max_column = column_record[i]
		}
	}
	var res []int = make([]int,total)
	var cnt int = 0
	for i := 0;i < rows;i++{
		col := 0
		for j := i;j >= 0;j--{
			if col < column_record[j]{
				res[cnt] = nums[j][col]
				cnt++
			}
			col++;
		}
	}
	for j := 1;j < max_column;j++{
		col := j
		for i := rows - 1;i >=0;i--{
			if col < column_record[i]{
				res[cnt] = nums[i][col]
				cnt++
			}
			col++
		}
	}
	return res
}