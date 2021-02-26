package number

import "math"

//152
//Input: [-2,0,-1]
//Output: 0
//[2,3,-2,4,-2,4,6,-9,3]
func dp_maxProduct(nums []int,begin int,memo map[int][]int)(min int,max int){
	if begin < 0{
		return 1,1
	}
	if nums[begin] == 0{
		return 0,0
	}
	if _,ok := memo[begin];ok{
		return memo[begin][0],memo[begin][1]
	}
	s,b := dp_maxProduct(nums,begin - 1,memo)
	big := max_int(max_int(b * nums[begin],s * nums[begin]),nums[begin])
	small := min_int(min_int(b * nums[begin],s * nums[begin]),nums[begin])
	memo[begin] = make([]int,2)
	memo[begin][0] = small
	memo[begin][1] = big
	return small,big
}

//不容易理解
func MaxProduct(nums []int) int {
	var res int = math.MinInt32
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < len(nums) ;i++{
		_,big :=  dp_maxProduct(nums,i,record)
		if big > res{
			res = big
		}
	}
	return res
}

func MaxProduct2(nums []int) int{
	l := len(nums)
	if l == 0{
		return 0
	}
	if l == 1{
		return nums[0]
	}
	var dp_min []int = make([]int,l)//nums[0:i] =  子数组乘积的最小值
	var dp_max []int = make([]int,l)//nums[0:i] =  子数组乘积的最大值
	dp_min[0] = nums[0]
	dp_max[0] = nums[0]
	var max int = nums[0]
	for i := 1;i < l;i++{
		product1 := nums[i] * dp_min[i-1]
		product2 := nums[i] * dp_max[i-1]
		dp_min[i] = min_int_number(nums[i],product1,product2)
		dp_max[i] = max_int_number(nums[i],product1,product2)
		if dp_max[i] > max{
			max = dp_max[i]
		}
	}
	return max
}