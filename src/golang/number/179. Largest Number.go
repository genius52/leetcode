package number

import (
	"strconv"
	"strings"
)

//179
func compare_pre_digit(a int,b int)bool{
	var s1 string = strconv.Itoa(a)
	var s2 string = strconv.Itoa(b)
	return strings.Compare(s1 + s2,s2 + s1) <= 0
}

func LargestNumber(nums []int) string {
	for i := 0;i < len(nums);i++{
		for j := 1;j < len(nums) - i;j++{
			if compare_pre_digit(nums[j-1],nums[j]){
				nums[j-1],nums[j] = nums[j],nums[j-1]
			}
		}
	}
	var res string
	for _,v := range nums{
		res += strconv.Itoa(v)
	}
	res = strings.TrimLeft(res,"0")
	if res == ""{
		return "0"
	}
	return res
}
