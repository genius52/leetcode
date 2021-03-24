package array

func dp_canCross(stones []int,l int,cur_pos int,steps int,first_step bool,allowed map[int]bool,memo map[int]map[int]bool)bool{
	if cur_pos > stones[l - 1]{
		return false
	}
	if cur_pos == stones[l - 1]{
		return true
	}
	if _,ok := allowed[cur_pos];!ok{
		return false
	}
	if _,ok := memo[cur_pos];ok{
		if _,ok2 := memo[cur_pos][steps];ok2{
			return false
		}
	}
	var res bool = false
	if steps > 1{
		res = dp_canCross(stones,l,cur_pos + steps - 1,steps - 1,false,allowed,memo)
	}
	if res{
		return true
	}
	res = dp_canCross(stones,l,cur_pos + steps,steps,false,allowed,memo)
	if res{
		return true
	}
	if !first_step{
		res = dp_canCross(stones,l,cur_pos + steps + 1,steps + 1,false,allowed,memo)
		if res{
			return true
		}
	}
	memo[cur_pos] = make(map[int]bool)
	memo[cur_pos][steps] = false
	return false
}

func CanCross(stones []int) bool {
	var l int = len(stones)
	var allowed map[int]bool = make(map[int]bool)
	for i := 0;i < l;i++{
		allowed[stones[i]] = true
	}
	var memo map[int]map[int]bool = make(map[int]map[int]bool)
	return dp_canCross(stones,l,0,1,true,allowed,memo)
}

func CanCross2(stones []int) bool {
	var l int = len(stones)
	var allowed_steps []map[int]bool = make([]map[int]bool,l)
	allowed_steps[0] = make(map[int]bool)
	allowed_steps[0][1] = true
	for i := 1;i < l;i++{
		allowed_steps[i] = make(map[int]bool)
		for j := i - 1;j >= 0;j--{
			for step,_ := range allowed_steps[j]{
				if stones[j] + step == stones[i]{
					if step > 1 && i != 1{
						allowed_steps[i][step - 1] = true
					}
					allowed_steps[i][step] = true
					allowed_steps[i][step + 1] = true
				}
			}
		}
	}
	if len(allowed_steps[l - 1]) > 0{
		return true
	}
	return false
}