package array

//offset > 0,move row to bottom
//offset < 0,move row to top
func cnt_largestOverlap(img1 [][]int, img2 [][]int,offset_row int,offset_col int)int{
	var l int = len(img1)
	var res int = 0
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			if (i + offset_row) >= 0 && (i + offset_row) < l && (j + offset_col) >= 0 && (j + offset_col) < l{
				if (img1[i + offset_row][j + offset_col] == img2[i][j]) && img1[i + offset_row][j + offset_col] == 1{
					res++
				}
			}
		}
	}
	return res
}

func LargestOverlap(img1 [][]int, img2 [][]int) int {
	var l int = len(img1)
	var res int = 0
	for i := 1 - l;i < l;i++{
		for j := 1 - l;j < l;j++{
			cur := cnt_largestOverlap(img1,img2,i,j)
			if cur > res{
				res = cur
				if res == (l * l){
					return res
				}
			}
		}
	}
	return res
}