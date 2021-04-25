package array

//You have a long flowerbed in which some of the plots are planted, and some are not.
//However, flowers cannot be planted in adjacent plots.
//
//Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
//and an integer n, return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.
//
//
//
//Example 1:
//
//Input: flowerbed = [1,0,0,0,1], n = 1
//Output: true
func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0{
		return true
	}
	flowerbed = append([]int{0},flowerbed...)
	flowerbed = append(flowerbed,0)
	var l int = len(flowerbed)
	var left int = 0
	//var record []int
	for left < l{
		for left < l && flowerbed[left] == 1{
			left++
		}
		if left >= l{
			break
		}
		var right int = left
		for right < l && flowerbed[right] == 0{
			right++
		}
		var cur_len int = right - left
		if cur_len > 2{
			n -= (cur_len - 1) / 2
			if n <= 0{
				return true
			}
		}
		left = right + 1
	}
	return false
}