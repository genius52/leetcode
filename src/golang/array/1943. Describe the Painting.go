package array

//Input: segments = [[1,4,5],[4,7,7],[1,7,9]]
//Output: [[1,4,14],[4,7,16]]
func SplitPainting(segments [][]int) [][]int64 {
	var record [100002]int
	var end [100002]bool
	var l int = len(segments)
	for i := 0;i < l;i++{
		record[segments[i][0]] += segments[i][2]
		record[segments[i][1]] -= segments[i][2]
		end[segments[i][0]] = true
		end[segments[i][1]] = true
	}
	var res [][]int64
	var start int = -1
	var cur_color int = 0
	for i := 0;i <= 100001;i++{
		cur_color += record[i]
		if end[i]{
			if start == -1{
				start = i
			}else{
				res = append(res,[]int64{int64(start),int64(i),int64(cur_color - record[i])})
				if cur_color == 0{
					start = -1
				}else{
					start = i
				}
			}
		}

	}
	return res
}