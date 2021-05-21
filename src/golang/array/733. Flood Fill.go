package array

func dfs_floodFill(image [][]int,rows int,columns int,r int,c int,find_color int,newColor int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if image[r][c] != find_color || image[r][c] == newColor{
		return
	}
	image[r][c] = newColor
	dfs_floodFill(image,rows,columns,r - 1,c,find_color,newColor)
	dfs_floodFill(image,rows,columns,r + 1,c,find_color,newColor)
	dfs_floodFill(image,rows,columns,r,c - 1,find_color,newColor)
	dfs_floodFill(image,rows,columns,r,c + 1,find_color,newColor)

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var rows int = len(image)
	var columns int = len(image[0])
	dfs_floodFill(image,rows,columns,sr,sc,image[sr][sc],newColor)
	return image
}