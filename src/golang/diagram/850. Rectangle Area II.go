package diagram

import "fmt"

//Input: rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
//Output: 6
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0]<rec2[2] && rec2[0] < rec1[2] && rec1[3] > rec2[1] && rec1[1] < rec2[3];
}

func add_to_no_overlap(no_overlap *[][]int,rect []int){
	var to_check [][]int = [][]int{rect}
	var l int = len(*no_overlap)
	for i := 0;i < l;i++{
		var check_len int = len(to_check)
		var new_add [][]int
		for j := 0;j < check_len;j++{
			if !isRectangleOverlap((*no_overlap)[i],to_check[j]){
				continue
			}
			left1 := to_check[j][0]
			bottom1 := to_check[j][1]
			right1 := to_check[j][2]
			top1 := to_check[j][3]

			left2 := (*no_overlap)[i][0]
			bottom2 := (*no_overlap)[i][1]
			right2 := (*no_overlap)[i][2]
			top2 := (*no_overlap)[i][3]
			if left1 < left2{
				var add []int = make([]int,4)
				add[0] = left1
				add[1] = bottom1
				//add[2] = min_int_number(right1,left2)
				add[2] = left2
				add[3] = top1
				new_add = append(new_add,add)
			}
			if right1 > right2{
				var add []int = make([]int,4)
				//add[0] = max_int_number(left1,right2)
				add[0] = right2
				add[1] = bottom1
				add[2] = right1
				add[3] = top1
				new_add = append(new_add,add)
			}
			if bottom1 < bottom2{
				var add []int = make([]int,4)
				add[0] = max_int_number(left1,left2)
				add[1] = bottom1
				add[2] = min_int_number(right1,right2)
				add[3] = bottom2
				new_add = append(new_add,add)
			}
			if top1 > top2{
				var add []int = make([]int,4)
				add[0] = max_int_number(left1,left2)
				add[1] = top2
				add[2] = min_int_number(right1,right2)
				add[3] = top1
				new_add = append(new_add,add)
			}
			to_check = append(to_check[:j],to_check[j + 1:]...)
			check_len = len(to_check)
		}
		to_check = append(to_check,new_add...)
	}
	if len(*no_overlap) == 0{
		(*no_overlap) = append((*no_overlap),to_check...)
	}else{
		var cur_len int = len(*no_overlap)
		for i := 0;i < len(to_check);i++{
			var overlap bool = false
			for j := 0;j < cur_len;j++{
				ret := isRectangleOverlap((*no_overlap)[j],to_check[i])
				if ret{
					overlap = true
					break
				}
			}
			if !overlap{
				(*no_overlap) = append((*no_overlap),to_check[i])
			}else{
				fmt.Println("fail")
			}
		}
	}
	//(*no_overlap) = append((*no_overlap),to_check...)
}

func RectangleArea(rectangles [][]int) int {
	var l int = len(rectangles)
	var no_overlap [][]int
	for i := 0;i < l;i++{
		add_to_no_overlap(&no_overlap,rectangles[i])
	}
	var res int = 0
	for i := 0;i < len(no_overlap);i++{
		res += (no_overlap[i][2] - no_overlap[i][0]) * (no_overlap[i][3] - no_overlap[i][1])
		res %= 1000000007
	}
	for i := 0;i < len(no_overlap);i++{
		for j := i + 1;j < len(no_overlap);j++{
			ret := isRectangleOverlap(no_overlap[i],no_overlap[j])
			if ret{
				fmt.Println(ret)
			}
		}
	}
	return res
}