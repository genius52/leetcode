package number

//让我们找到一个最小的数字，在卡的背面，且要求其他卡正面上均没有这个数字。
//简而言之，就是要在backs数组找一个最小数字，使其不在fronts数组中

//Input: fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
//Output: 2
//Explanation: If we flip the second card, the fronts are [1,3,4,4,7] and the backs are [1,2,4,1,3].
//We choose the second card, which has number 2 on the back, and it isn't on the front of any card, so 2 is good.
func flipgame(fronts []int, backs []int) int {
	var same map[int]bool = make(map[int]bool)
	var l int = len(fronts)
	for i := 0;i < l;i++{
		if fronts[i] == backs[i]{
			same[fronts[i]] = true
		}
	}
	var min_num int = 2147483647
	var find bool = false
	for i := 0;i < l;i++{
		if fronts[i] == backs[i] {
			continue
		}
		if _,ok := same[fronts[i]];!ok{
			if fronts[i] < min_num{
				min_num = fronts[i]
			}
			find = true
		}
		if _,ok := same[backs[i]];!ok{
			if backs[i] < min_num{
				min_num = backs[i]
			}
			find = true
		}
	}
	if !find{
		return 0
	}
	return min_num
}