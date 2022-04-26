package array

//DFS Solution
func dfs_pacificAtlantic(heights [][]int,rows int,columns int,r int,c int,pre_height int,visited [][]bool){
	if r < 0 || r >= rows || c < 0 || c >= columns || visited[r][c]{
		return
	}
	if heights[r][c] < pre_height{
		return
	}
	visited[r][c] = true
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		dfs_pacificAtlantic(heights,rows,columns,r + dir[0],c + dir[1],heights[r][c],visited)
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	var rows int = len(heights)
	var columns int = len(heights[0])
	var visited1 [][]bool = make([][]bool,rows)
	var visited2 [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		visited1[i] = make([]bool,columns)
		visited2[i] = make([]bool,columns)
	}
	for i := 0;i < rows;i++{
		//visited1[i][0] = true
		dfs_pacificAtlantic(heights,rows,columns,i,0,-1,visited1)
		dfs_pacificAtlantic(heights,rows,columns,i,columns - 1,-1,visited2)
	}
	for j := 0;j < columns;j++{
		//visited1[0][j] = true
		dfs_pacificAtlantic(heights,rows,columns,0,j,-1,visited1)
		dfs_pacificAtlantic(heights,rows,columns,rows - 1,j,-1,visited2)
	}
	var res [][]int
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if visited1[i][j] && visited2[i][j]{
				res = append(res,[]int{i,j})
			}
		}
	}
	return res
}