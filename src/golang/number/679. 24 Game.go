package number

import (
	"container/list"
	"math"
)

//Input: cards = [4,1,8,7]
//Output: true
//Explanation: (8-4) * (7-1) = 24
func JudgePoint24(cards []int) bool {
	var q list.List
	var arr []float64 = []float64{float64(cards[0]),float64(cards[1]),float64(cards[2]),float64(cards[3])}
	q.PushBack(arr)
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var prev_nums []float64 = q.Front().Value.([]float64)
			q.Remove(q.Front())
			var arr_len int = len(prev_nums)
			if arr_len == 1{
				if math.Abs(prev_nums[0] - 24.0) <= 0.001{
					return true
				}
				continue
			}
			//random pick 2 number then operation on them
			for m := 0;m < arr_len;m++{
				for n := m + 1;n < arr_len;n++{
					var new_nums [5][]float64
					for j := 0;j < arr_len;j++{
						if j == m || j == n{
							continue
						}
						for k := 0;k < 5;k++{
							new_nums[k] = append(new_nums[k],float64(prev_nums[j]))
						}
					}
					new_nums[0] = append(new_nums[0],prev_nums[m] + prev_nums[n])//plus
					new_nums[1] = append(new_nums[1],math.Abs(prev_nums[m] - prev_nums[n]))//minus
					new_nums[2] = append(new_nums[2],prev_nums[m] * prev_nums[n])//multiply
					new_nums[3] = append(new_nums[3],math.Abs(prev_nums[m] /prev_nums[n]))//divide
					new_nums[4] = append(new_nums[4],math.Abs(prev_nums[n] /prev_nums[m]))//divide
					q.PushBack(new_nums[0])
					q.PushBack(new_nums[1])
					q.PushBack(new_nums[2])
					q.PushBack(new_nums[3])
					q.PushBack(new_nums[4])
				}
			}
		}
	}
	return false
}