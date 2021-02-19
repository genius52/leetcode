package array

func sortColors(nums []int)  {
	zero_cnt := 0
	one_cnt := 0
	two_cnt := 0
	for i:= 0;i < len(nums);i++{
		if(nums[i] == 0){
			zero_cnt++
		}else if(nums[i] == 1){
			one_cnt++
		}else{
			two_cnt++
		}
	}
	var index int = 0
	for index < zero_cnt{
		nums[index] = 0
		index++
	}
	for index < zero_cnt + one_cnt{
		nums[index] = 1
		index++
	}
	for index < zero_cnt + one_cnt + two_cnt{
		nums[index] = 2
		index++
	}
}