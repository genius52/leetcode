package array

import "strconv"

//nput:  [0,2,3,4,6,8,9]
//Output: ["0","2->4","6","8->9"]
//Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
func SummaryRanges(nums []int) []string {
	l := len(nums)
	slow := 0
	fast := 0
	var res []string
	for fast < l{
		for fast < l - 1 && nums[fast] == nums[fast + 1] - 1{
			fast++
		}
		if slow == fast{
			res = append(res,strconv.Itoa(nums[slow]))
		}else{
			res = append(res,strconv.Itoa(nums[slow]) + "->" + strconv.Itoa(nums[fast]))
		}
		fast++
		slow = fast
	}
	return res
}
