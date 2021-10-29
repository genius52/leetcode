package array

//Input: mat = [[1,0,1],
//              [1,1,0],
//              [1,1,0]]
//Output: 13
//Explanation:
//There are 6 rectangles of side 1x1.
//There are 2 rectangles of side 1x2.
//There are 3 rectangles of side 2x1.
//There is 1 rectangle of side 2x2.
//There is 1 rectangle of side 3x1.
//Total number of rectangles = 6 + 2 + 3 + 1 + 1 = 13.
func count_numSubmat(mat [][]int,row int,column int)int{
	var rows int = len(mat)
	var columns int = len(mat[0])
	var cnt int = 0
	var right_bound int = columns - 1
	for i := row;i < rows;i++{
		if mat[i][column] == 0{
			break
		}
		for j := column;j <= right_bound;j++{
			if mat[i][j] == 0{
				right_bound = j - 1
				break
			}else{
				cnt++
			}
		}
	}
	return cnt
}

func NumSubmat(mat [][]int) int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if mat[i][j] == 0{
				continue
			}
			res += count_numSubmat(mat,i,j)
		}
	}
	return res
}