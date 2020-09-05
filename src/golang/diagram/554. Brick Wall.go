package diagram

func LeastBricks(wall [][]int) int {
	var record map[int]int = make(map[int]int)
	var rows int = len(wall)
	//var columns int = len(wall[0])
	var max_occurs int = 0
	for r := 0;r < rows;r++{
		var pos int = 0
		columns := len(wall[r])
		for c := 0;c < columns - 1;c++{
			pos += wall[r][c]
			record[pos]++
			if record[pos] > max_occurs{
				max_occurs = record[pos]
			}
		}
	}
	return rows - max_occurs
}