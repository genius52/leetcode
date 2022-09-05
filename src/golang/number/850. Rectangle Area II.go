package number

import "sort"

func rectangleArea(rectangles [][]int) int {
	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][0] == rectangles[j][0]{
			return rectangles[i][1] < rectangles[j][1]
		}
		return rectangles[i][0] < rectangles[j][0]
	})
	var l int = len(rectangles)
	var res int = 0
	for i := 0;i < l;i++{

	}
	return res
}