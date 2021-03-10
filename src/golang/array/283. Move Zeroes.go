package array

//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]
func MoveZeroes2(nums []int){
	var l int = len(nums)
	var visit int = 0
	for i := 0;i < l;i++{
		if nums[i] != 0{
			nums[visit] = nums[i]
			visit++
		}
	}
	for visit < l{
		nums[visit] = 0
		visit++
	}
}

func MoveZeroes(nums []int)  {
	var l int = len(nums)
	var zero_index int = 0;
	for zero_index < l {
		for zero_index < l && nums[zero_index] != 0{
			zero_index++
		}
		if zero_index >= l{
			break
		}
		nozero := zero_index
		for nozero < l{
			if nums[nozero] == 0{
				nozero++
			}else{
				break
			}
		}
		if nozero >= l{
			break
		}
		nums[zero_index] = nums[nozero]
		nums[nozero] = 0
		zero_index++
	}
}