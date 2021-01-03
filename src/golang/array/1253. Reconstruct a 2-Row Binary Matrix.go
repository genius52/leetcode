package array

//Input: upper = 2, lower = 1, colsum = [1,1,1]
//Output: [[1,1,0],[0,0,1]]
//Explanation: [[1,0,1],[0,1,0]], and [[0,1,1],[1,0,0]] are also correct answers.
func ReconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var columns int = len(colsum)
	var res [][]int = make([][]int,2)
	res[0] = make([]int,columns)
	res[1] = make([]int,columns)
	for i := 0;i < columns;i++{
		if colsum[i] == 0{
			res[0][i] = 0
			res[1][i] = 0
		}else if colsum[i] == 2{
			res[0][i] = 1
			upper--
			res[1][i] = 1
			lower--
		}else if colsum[i] == 1{
			if upper >= lower{
				res[0][i] = 1
				res[1][i] = 0
				upper--
			}else{
				res[0][i] = 0
				res[1][i] = 1
				lower--
			}
		}
		if upper < 0 || lower < 0{
			return [][]int{}
		}
	}
	if upper != 0 || lower != 0{
		return [][]int{}
	}
	return res
}