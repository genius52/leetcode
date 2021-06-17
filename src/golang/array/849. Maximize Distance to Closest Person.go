package array

//Input: seats = [1,0,0,0,1,0,1]
//Output: 2
func maxDistToClosest(seats []int) int {
	var l int = len(seats)
	var left int = 0
	for left < l && seats[left] != 1{
		left++
	}
	var res int = left
	var right int = l - 1
	for right >= 0 && seats[right] != 1{
		right--
	}
	if  (l - right - 1) > res{
		res =  l - right - 1
	}
	var head int = left
	for head < right {
		for head < right && seats[head] == 0{
			head++
		}
		var tail int = head + 1
		for tail < right && seats[tail] == 0{
			tail++
		}
		dis := (tail - head)/2
		if dis > res{
			res = dis
		}
		head = tail
	}
	return res
}