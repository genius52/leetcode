package array

func is_same(mat1 [][]int,mat2 [][]int,l int)bool{
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			if mat1[i][j] != mat2[i][j]{
				return false
			}
		}
	}
	return true
}

func FindRotation(mat [][]int, target [][]int) bool {
	var l int = len(mat)
	for i := 0;i < 4;i++{
		res := is_same(mat,target,l)
		if res{
			return true
		}
		var rotate [][]int = make([][]int,l)
		for j := 0;j < l;j++{
			rotate[j] = make([]int,l)
			for k := 0;k < l;k++{
				rotate[j][k] = mat[l - 1 - k][j]
			}
		}
		mat = rotate
	}
	return false
}