package array

func resultGrid(image [][]int, threshold int) [][]int {
	var rows int = len(image)
	var columns int = len(image[0])
	var res [][]int = make([][]int, rows)
	var cnt [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, columns)
		cnt[i] = make([]int, columns)
	}
	for i := 2; i < rows; i++ {
		for j := 2; j < columns; j++ {
			var meet bool = true
			for r := i - 2; r <= i; r++ {
				if abs_int(image[r][j-2]-image[r][j-1]) > threshold || //横向临近比较
					abs_int(image[r][j-1]-image[r][j]) > threshold {
					meet = false
					break
				}
			}
			if !meet {
				continue
			}
			for c := j - 2; c <= j; c++ {
				if abs_int(image[i-2][c]-image[i-1][c]) > threshold || //纵向临近比较
					abs_int(image[i-1][c]-image[i][c]) > threshold {
					meet = false
					break
				}
			}
			if !meet {
				continue
			}
			var cur_val int = (image[i-2][j-2] + image[i-2][j-1] + image[i-2][j] +
				image[i-1][j-2] + image[i-1][j-1] + image[i-1][j] +
				image[i][j-2] + image[i][j-1] + image[i][j]) / 9
			for m := i - 2; m <= i; m++ {
				for n := j - 2; n <= j; n++ {
					res[m][n] += cur_val
					cnt[m][n]++
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if cnt[i][j] == 0 {
				res[i][j] = image[i][j]
			} else {
				res[i][j] /= cnt[i][j]
			}
		}
	}
	return res
}
