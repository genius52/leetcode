package array


// # : stone
// * : block
// . : empty

func RotateTheBox(box [][]byte) [][]byte {
	var rows int = len(box)
	var columns int = len(box[0])
	for i := 0;i < rows;i++{
		var empty_index int = columns - 1
		for j := columns - 1;j >= 0;j--{
			if box[i][j] == '#'{
				for empty_index >= 0 && empty_index > j{
					if box[i][empty_index] == '.'{
						box[i][empty_index] = '#'
						box[i][j] = '.'
						empty_index--
						break
					}
					empty_index--
				}
			}else if box[i][j] == '*'{
				empty_index = j - 1
			}else if box[i][j] == '.'{
				//empty_index = j//?
			}
		}
	}
	var res [][]byte = make([][]byte,columns)
	for i := 0;i < columns;i++{
		res[i] = make([]byte,rows)
		for j := 0;j < rows;j++{
			res[i][j] = box[rows - 1 - j][i]
		}
	}
	return res
}

//Input: box = [["#",".","*","."],
//              ["#","#","*","."]]
//Output: [["#","."],
//         ["#","#"],
//         ["*","*"],
//         [".","."]]