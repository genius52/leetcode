package array

//Input: targetGrid = [[1,1,1,1],[1,1,3,3],[1,1,3,4],[5,5,1,4]]
//Output: true

//topologic sort
//数字大的颜色指向包围他的颜色
func isPrintable(targetGrid [][]int) bool {
	var rows int = len(targetGrid)
	var columns int = len(targetGrid[0])
	var max_val int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if targetGrid[i][j] > max_val{
				max_val = targetGrid[i][j]
			}
		}
	}
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		visited[i] = make([]bool,columns)
	}
	var find_value int = max_val
	for find_value > 0{
		//var find bool = false
		for i := 0;i < rows;i++{
			for j := 0;j < columns;j++{
				if targetGrid[i][j] == find_value{

				}
			}
		}
		find_value--
	}
	//return dfs_isPrintable(targetGrid,rows,columns,)
	return true
}