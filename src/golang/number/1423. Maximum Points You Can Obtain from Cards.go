package number

//Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
//Output: 202
func MaxScore2(cardPoints []int, k int) int{
	l := len(cardPoints)
	var sum int = 0
	for i := l - k;i < l;i++{
		sum += cardPoints[i]
	}
	var res int = sum
	for i := l - k;i < l ;i++{
		cur := sum - cardPoints[i % l] + cardPoints[(i + k) % l]
		if cur > res{
			res = cur
		}
		sum = cur
	}
	return res
}

//func MaxScore2(cardPoints []int, k int) int {
//	l := len(cardPoints)
//	begin := l - k
//	end := l - 1
//	var max int = 0
//	for i := begin;i <= end;i++{
//		max += cardPoints[i]
//	}
//	var cur_sum int = max
//	for i := begin;i < l;i++{
//		cur_sum = cur_sum - cardPoints[i] + cardPoints[(i + k) % l]
//		if cur_sum > max{
//			max = cur_sum
//		}
//	}
//	return max
//}