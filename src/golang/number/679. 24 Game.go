package number

import (
	"container/list"
	"math"
	"strconv"
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

type express struct{
	val float64
	pre_exp string
}

func JudgePoint24ex(cards []int)[]string {
	var q list.List
	var arr []express = []express{express{float64(cards[0]), strconv.Itoa(cards[0]),},express{float64(cards[1]), strconv.Itoa(cards[1]),},
		express{float64(cards[2]), strconv.Itoa(cards[2]),},express{float64(cards[3]), strconv.Itoa(cards[3]),}}
	q.PushBack(arr)
	var res []string
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var prev_nums []express = q.Front().Value.([]express)
			q.Remove(q.Front())
			var arr_len int = len(prev_nums)
			if arr_len == 1{
				if math.Abs(prev_nums[0].val - 24.0) <= 0.001{
					res = append(res,prev_nums[0].pre_exp)
				}
				continue
			}
			//random pick 2 number then operation on them
			for m := 0;m < arr_len;m++{
				for n := m + 1;n < arr_len;n++{
					var new_nums [5][]express
					for j := 0;j < arr_len;j++{
						if j == m || j == n{
							continue
						}
						for k := 0;k < 5;k++{
							new_nums[k] = append(new_nums[k],prev_nums[j])
						}
					}
					new_nums[0] = append(new_nums[0],express{prev_nums[m].val + prev_nums[n].val,
						"(" + prev_nums[m].pre_exp + " + " + prev_nums[n].pre_exp + ")",})

					if prev_nums[m].val > prev_nums[n].val{
						new_nums[1] = append(new_nums[1],express{math.Abs(prev_nums[m].val - prev_nums[n].val),
							"(" + prev_nums[m].pre_exp + " - " + prev_nums[n].pre_exp + ")",})
					}else{
						new_nums[1] = append(new_nums[1],express{math.Abs(prev_nums[n].val - prev_nums[m].val),
							"(" + prev_nums[n].pre_exp + " - " + prev_nums[m].pre_exp + ")",})
					}

					new_nums[2] = append(new_nums[2],express{prev_nums[m].val * prev_nums[n].val,
				"(" + prev_nums[m].pre_exp + " * " + prev_nums[n].pre_exp + ")",})

					new_nums[3] = append(new_nums[3],express{prev_nums[m].val / prev_nums[n].val,
					"(" + prev_nums[m].pre_exp + " / " + prev_nums[n].pre_exp + ")",})
					new_nums[4] = append(new_nums[4],express{prev_nums[n].val / prev_nums[m].val,
						"(" + prev_nums[n].pre_exp + " / " + prev_nums[m].pre_exp + ")",})

					//new_nums[0] = append(new_nums[0],prev_nums[m].val + prev_nums[n].val)//plus
					//new_nums[1] = append(new_nums[1],math.Abs(prev_nums[m].val - prev_nums[n].val))//minus
					//new_nums[2] = append(new_nums[2],prev_nums[m].val * prev_nums[n].val)//multiply
					//new_nums[3] = append(new_nums[3],math.Abs(prev_nums[m].val /prev_nums[n].val))//divide
					//new_nums[4] = append(new_nums[4],math.Abs(prev_nums[n].val /prev_nums[m].val))//divide
					q.PushBack(new_nums[0])
					q.PushBack(new_nums[1])
					q.PushBack(new_nums[2])
					q.PushBack(new_nums[3])
					q.PushBack(new_nums[4])
				}
			}
		}
	}
	return res
}