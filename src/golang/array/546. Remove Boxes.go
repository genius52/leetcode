package array

//Input: boxes = [1,3,2,2,2,3,4,3,1]
//Output: 23
//Explanation:
//[1, 3, 2, 2, 2, 3, 4, 3, 1]
//----> [1, 3, 3, 4, 3, 1] (3*3=9 points)
//----> [1, 3, 3, 3, 1] (1*1=1 points)
//----> [1, 1] (3*3=9 points)
//----> [] (2*2=4 points)
//func dp_removeBoxes(boxes []int,l int,r int,k int,memo [100][100][100]int)int{
//	if l > r{
//		return 0
//	}
//	if memo[l][r][k] != 0{
//		return memo[l][r][k]
//	}
//	//optimize
//	for l < r && boxes[l] == boxes[l + 1]{
//		l++
//		k++
//	}
//	var max_score int = (k + 1) * (k + 1) + dp_removeBoxes(boxes,l + 1,r,k,memo)
//	for i := l + 1;i < r;i++{
//
//		cur_score :=
//	}
//	memo[l][r][k] = max_score
//	return max_score
//}
//
//func RemoveBoxes(boxes []int) int {
//	var l int = len(boxes)
//	var memo [100][100][100]int
//	return dp_removeBoxes(boxes,0,l - 1,0,memo)
//}

func dp_removeBoxes(boxes []int)int{
	var l int = len(boxes)
	if l == 0 {
		return 0
	}
	var max_scores int = 0
	for left := 0;left < l;{
		if left > 0 && boxes[left] == boxes[left - 1]{
			left++
			continue
		}
		var right int = left
		for right < l && boxes[right] == boxes[left]{
			right++
		}
		scores := (right - left) * (right - left)
		var copy []int// = make([]int,l - (right - left + 1))
		if left > 0{
			copy = append(copy,boxes[:left]...)
		}
		if right < l{
			copy = append(copy,boxes[right:]...)
		}
		scores += dp_removeBoxes(copy)
		if scores > max_scores{
			max_scores = scores
		}
		left = right
	}
	return max_scores
}

func RemoveBoxes(boxes []int) int {
	return dp_removeBoxes(boxes)
}