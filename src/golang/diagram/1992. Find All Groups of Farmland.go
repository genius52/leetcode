package diagram

func dfs_findFarmland(land *[][]int,rows int,columns int,r int,c int)(int,int,int,int){
	if r < 0 || r >= rows || c < 0 || c >= columns || (*land)[r][c] != 1{
		return rows,columns,0,0
	}
	(*land)[r][c] = 0
	x1,y1,x11,y11 := dfs_findFarmland(land,rows,columns,r - 1,c)
	x2,y2,x22,y22 := dfs_findFarmland(land,rows,columns,r + 1,c)
	x3,y3,x33,y33 := dfs_findFarmland(land,rows,columns,r,c - 1)
	x4,y4,x44,y44 := dfs_findFarmland(land,rows,columns,r,c + 1)
	return min_int_number(r,x1,x2,x3,x4),min_int_number(c,y1,y2,y3,y4),
	max_int_number(r,x11,x22,x33,x44),max_int_number(c,y11,y22,y33,y44)
}

func findFarmland(land [][]int) [][]int {
	var rows int = len(land)
	var columns int = len(land[0])
	var res [][]int
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if land[i][j] != 1{
				continue
			}
			x1,y1,x2,y2 := dfs_findFarmland(&land,rows,columns,i,j)
			res = append(res,[]int{x1,y1,x2,y2})
		}
	}
	return res
}