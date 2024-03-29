package diagram

//Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
//Output: 2
//Explanation: The maximum side length of square with sum less than 4 is 2 as shown.
func is_threshold(prefix_sum [][]int,rows int,columns int,len int,threshold int)bool{
	for i := 0;i + len <= rows;i++{
		for j := 0;j + len <= columns;j++{
			val := prefix_sum[i + len][j + len] - prefix_sum[i][j + len] - prefix_sum[i + len][j] + prefix_sum[i][j]
			if val <= threshold{
				return true
			}
		}
	}
	return false
}

func MaxSideLength(mat [][]int, threshold int) int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var prefix_sum [][]int = make([][]int,rows + 1)
	for i := 0;i <= rows;i++{
		prefix_sum[i] = make([]int,columns + 1)
	}
	var res int = 0
	for i := 1;i <= rows;i++{
		for j := 1;j <= columns;j++{
			prefix_sum[i][j] = mat[i - 1][j - 1]
			if mat[i - 1][j - 1] <= threshold{
				res = 1
			}
			if i > 0{
				prefix_sum[i][j] += prefix_sum[i - 1][j]
			}
			if j > 0{
				prefix_sum[i][j] += prefix_sum[i][j - 1]
			}
			if i > 0 && j > 0{
				prefix_sum[i][j] -= prefix_sum[i - 1][j - 1]
			}
		}
	}
	if res == 0{
		return res
	}
	var low int = 1
	var high int = rows
	if columns < rows{
		high = columns
	}
	for low <= high{
		var find bool = false
		var mid int = (low + high)/2
		if is_threshold(prefix_sum,rows,columns,mid,threshold){
			res = mid
			find = true
		}
		if find{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return high
}