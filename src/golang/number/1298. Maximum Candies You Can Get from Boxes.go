package number

func MaxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var n int = len(status)
	var visited_box []bool = make([]bool,n)
	var total int = 0
	cur_keys := make([]bool,n)
	var cur_avilable_box map[int]bool = make(map[int]bool)
	for _,box := range initialBoxes{
		if status[box] == 1{
			for j := 0;j < len(keys[box]);j++{
				cur_keys[keys[box][j]] = true
			}
			total += candies[box]
			visited_box[box] = true
			for j := 0;j < len(containedBoxes[box]);j++{
				cur_avilable_box[containedBoxes[box][j]] = true
			}
		}else{
			cur_avilable_box[box] = true
		}
	}
	for true{
		var opened_box []int
		var add_box []int
		for box,_ := range cur_avilable_box{
			if visited_box[box]{
				continue
			}
			if status[box] == 1 {//箱子是打开状态
				for j := 0;j < len(keys[box]);j++{
					cur_keys[keys[box][j]] = true //取出钥匙
				}
				visited_box[box] = true
				opened_box = append(opened_box,box)
				total += candies[box]//取出糖果
				for _,next_box := range containedBoxes[box]{
					if !visited_box[next_box]{
						add_box = append(add_box,next_box)
					}
				}
			}else{
				if cur_keys[box]{//有这个箱子的钥匙
					for j := 0;j < len(keys[box]);j++{
						cur_keys[keys[box][j]] = true //取出钥匙
					}
					opened_box = append(opened_box,box)
					total += candies[box]
					for _,next_box := range containedBoxes[box]{
						if !visited_box[next_box]{
							add_box = append(add_box,next_box)
						}
					}
				}
			}
		}
		if len(opened_box) == 0 && len(add_box) == 0{
			break
		}
		for _,box := range add_box{
			cur_avilable_box[box] = true
		}
		for _,box := range opened_box{
			delete(cur_avilable_box,box) //删除已经取出糖果和钥匙的盒子
		}
	}
	return total
}