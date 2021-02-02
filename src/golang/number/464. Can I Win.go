package number

func format_slice(s []int)int{
	var res int = 0
	for _,n := range s{
		res |= 1 << n
	}
	return res
}

func dfs_canIWin(choose_num []int, desiredTotal int,memo map[int]bool)bool{
	var l int = len(choose_num)
	if l == 0{
		return false
	}
	if choose_num[l - 1] >= desiredTotal{
		return true;
	}
	k := format_slice(choose_num)
	if ret,ok := memo[k];ok{
		return ret
	}
	for i := 0;i < l;i++{
		//player 2 roound
		var copy []int
		copy = append(copy,choose_num[:i]...)
		copy = append(copy,choose_num[i + 1:]...)
		ret := dfs_canIWin(copy,desiredTotal - choose_num[i],memo)
		if !ret{
			memo[k] = true
			return true
		}
	}
	memo[k] = false
	return false
}

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	//I can win
	if desiredTotal <= maxChoosableInteger{
		return true
	}
	//no one can win
	if ((1 + maxChoosableInteger) / 2 * maxChoosableInteger) < desiredTotal {
		return false;
	}
	var choose_num []int = make([]int,maxChoosableInteger)
	for i := 0;i < maxChoosableInteger;i++{
		choose_num[i] = i + 1
	}
	var memo map[int]bool = make(map[int]bool)
	return dfs_canIWin(choose_num,desiredTotal,memo)
}