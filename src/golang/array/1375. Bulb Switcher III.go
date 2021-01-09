package array

//Input: light = [2,1,3,5,4]
//Output: 3
func NumTimesAllBlue(light []int) int {
	var l int = len(light)
	var right int = light[0]
	var res int = 0
	if light[0] == 1{
		res++
	}
	for step := 1;step < l;step++{
		if light[step] > right {
			if light[step] == step + 1{
				res++
			}
			right = light[step]
		}else{
			if step + 1 == right{
				res++
			}
		}
	}
	return res
}