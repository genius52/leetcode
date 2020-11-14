package number

func dfs_canIWin(maxChoosableInteger int, desiredTotal int,unchosen_nums []bool,is_my_round bool,memo map[int]bool)bool{
	if desiredTotal <= 0{
		return false
	}
	//if r,ok := memo[desiredTotal];ok{
	//	return r
	//}
	for i := 1;i <= maxChoosableInteger;i++{
		if unchosen_nums[i]{
			continue
		}
		unchosen_nums[i] = true
		if !dfs_canIWin(maxChoosableInteger,desiredTotal - i,unchosen_nums,!is_my_round,memo){
			//memo[desiredTotal] = true
			unchosen_nums[i] = false
			return true
		}
		unchosen_nums[i] = false
	}
	//memo[desiredTotal] = false
	return false
}

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= maxChoosableInteger{
		return true
	}
	if ((1 + maxChoosableInteger) / 2 * maxChoosableInteger) < desiredTotal {
		return false;
	}
	var unchosen_nums []bool = make([]bool,maxChoosableInteger + 1)
	var memo map[int]bool = make(map[int]bool)
	return dfs_canIWin(maxChoosableInteger,desiredTotal,unchosen_nums,true,memo)
}