package array

func FlipAndInvertImage(image [][]int) [][]int {
	var rows int = len(image)
	var columns int = len(image[0])
	for i := 0;i < rows;i++{
		for j := 0;j < columns/2;j++{
			if image[i][j] == image[i][columns - 1 - j]{
				image[i][j] = (image[i][j] + 1) % 2
				image[i][columns - 1 - j] = (image[i][columns - 1 - j] + 1) % 2
			}
		}
		if (columns | 1) == columns{
			image[i][columns/2] = (image[i][columns/2] + 1) % 2
		}
	}
	return image
}