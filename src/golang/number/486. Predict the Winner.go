package number

import "math"

//Input: [1, 5, 2]
//Output: False
//Input: [1, 5, 233, 7]
//Output: True
func dp_PredictTheWinner(nums []int,head int,tail int,score1 int,score2 int,player1_choose bool)bool{
	if(head == tail){
		if player1_choose{
			return (score1 + nums[head]) >= score2
		}else{
			return score1 >= (score2 + nums[head])
		}
	}
	var choose_number int = 0
	if (head == tail - 1){
		if nums[head] > nums[tail]{
			choose_number = nums[head]
			head++
		}else{
			choose_number = nums[tail]
			tail--
		}
	}else{
		var choose_head int = nums[head] - int(math.Max(float64(nums[head + 1]),float64(nums[tail])))
		var choose_tail int = nums[tail] - int(math.Max(float64(nums[tail - 1]),float64(nums[head])))
		if choose_head >= choose_tail{
			choose_number = nums[head]
			head++
		}else{
			choose_number = nums[tail]
			tail--
		}
	}
	if player1_choose{
		score1 += choose_number
	}else{
		score2 += choose_number
	}
	return dp_PredictTheWinner(nums,head,tail,score1,score2,!player1_choose)
}

func PredictTheWinner(nums []int) bool {
	var head int = 0
	var tail int = len(nums) - 1
	return dp_PredictTheWinner(nums,head,tail,0,0,true)
}