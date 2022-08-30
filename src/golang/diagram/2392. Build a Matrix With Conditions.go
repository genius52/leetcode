package diagram

import "container/list"

func BuildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	var up_down [][]int = make([][]int,k + 1)
	var indegree_up_down []int = make([]int,k + 1)
	for _,row := range rowConditions{
		up_down[row[0]] = append(up_down[row[0]],row[1])
		indegree_up_down[row[1]]++
	}
	var left_right [][]int = make([][]int,k + 1)
	var indegree_left_right []int = make([]int,k + 1)
	for _,col := range colConditions{
		left_right[col[0]] = append(left_right[col[0]],col[1])
		indegree_left_right[col[1]]++
	}
	var q1 list.List
	var horizon []int
	for i := 1;i <= k;i++{
		if indegree_left_right[i] == 0{
			horizon = append(horizon,i)
			q1.PushBack(i)
		}
	}
	for q1.Len() > 0{
		var cur_len int = q1.Len()
		for i := 0;i < cur_len;i++{
			var left int = q1.Front().Value.(int)
			q1.Remove(q1.Front())
			for _,right := range left_right[left]{
				indegree_left_right[right]--
				if indegree_left_right[right] == 0{
					q1.PushBack(right)
					horizon = append(horizon,right)
				}
			}
		}
	}
	var res [][]int
	if len(horizon) != k{
		return res
	}
	var q2 list.List
	var vertical []int
	for i := 1;i <= k;i++{
		if indegree_up_down[i] == 0{
			vertical = append(vertical,i)
			q2.PushBack(i)
		}
	}
	for q2.Len() > 0{
		var cur_len int = q2.Len()
		for i := 0;i < cur_len;i++{
			var up int = q2.Front().Value.(int)
			q2.Remove(q2.Front())
			for _,down := range up_down[up]{
				indegree_up_down[down]--
				if indegree_up_down[down] == 0{
					vertical = append(vertical,down)
					q2.PushBack(down)
				}
			}

		}
	}
	if len(vertical) != k{
		return res
	}
	res = make([][]int,k)
	for i := 0;i < k;i++{
		res[i] = make([]int,k)
	}
	for i := 0;i < k;i++{
		var col int = -1
		for j := 0;j < k;j++{
			if horizon[j] == vertical[i]{
				col = j
				break
			}
		}
		res[i][col] = vertical[i]
	}
	return res
}