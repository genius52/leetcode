package number

func singleNumber3(nums []int) []int {
	var val int = 0
	for i := 0;i < len(nums);i++{
		val = val ^ nums[i]
	}
	var tag int = 1
	for i := 1;i < 32;i++{
		if val | tag == val{
			break
		}else{
			tag = tag << 1
		}
	}
	var group [][]int = make([][]int,2)
	for i := 0;i < len(nums);i++{
		if nums[i] | tag == nums[i]{
			group[0] = append(group[0],nums[i])
		}else{
			group[1] = append(group[1],nums[i])
		}
	}
	var first_num,second_num int = 0,0
	for i := 0;i < len(group[0]);i++{
		first_num ^= group[0][i]
	}
	for i := 0;i < len(group[1]);i++{
		second_num ^= group[1][i]
	}

	return []int{first_num,second_num}
}