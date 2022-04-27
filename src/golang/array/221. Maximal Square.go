package array

func maximalSquare(matrix [][]byte) int {
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			matrix[i][j] -= '0'
		}
	}

	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				continue
			}
			matrix[i][j] = 1 + min_byte_number(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
	}
	var res int = 0
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			if int(matrix[i][j]) > res{
				res = int(matrix[i][j])
			}
		}
	}
	return res * res
}

func min_byte_number(nums ...byte)byte{
	var min byte = 127
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}