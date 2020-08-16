package array

func check_route(nums []int,pos int,forward bool,trace map[int]bool)bool{
	var l int = len(nums)
	if nums[pos] % l == 0{
		trace[pos] = true
		return false
	}
	var next_pos int
	if forward{
		if nums[pos] < 0{
			trace[pos] = true
			return false
		}
		next_pos = (pos + nums[pos]) % l
	}else{
		if nums[pos] > 0{
			trace[pos] = true
			return false
		}
		next_pos = (pos + nums[pos] % l + l) % l
	}
	if _,ok := trace[pos];ok{
		return true
	}
	trace[pos] = true
	return check_route(nums,next_pos,forward,trace)
}

func CircularArrayLoop(nums []int) bool {
	var l int = len(nums)
	var bad_pos map[int]bool = make(map[int]bool)
	for begin := 0;begin < l;begin++{
		if nums[begin] % l == 0{
			bad_pos[begin] = true
			continue
		}
		if _,ok := bad_pos[begin];ok{
			continue
		}
		var trace map[int]bool = make(map[int]bool)
		var forward bool = true
		if(nums[begin] < 0){
			forward = false
		}
		if check_route(nums,begin,forward,trace){
			return true
		}
		for pos,_ := range trace{
			bad_pos[pos] = true
		}
	}
	return false
}