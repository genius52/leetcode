package number

func NumberOfWays(corridor string) int {
	var l int = len(corridor)
	var seat_cnt int = 0
	var idx []int
	for i := 0;i < l;i++{
		if corridor[i] == 'S'{
			seat_cnt++
			idx = append(idx,i)
		}
	}
	if seat_cnt <= 1 || (seat_cnt | 1 == seat_cnt){
		return 0
	}
	if seat_cnt == 2{
		return 1
	}
	var space []int
	for i := 1;i + 1 < len(idx);i += 2{
		space = append(space,idx[i + 1] - idx[i])
	}
	var MOD int =  1e9 + 7
	var res int = 1
	for i := 0;i < len(space);i++{
		res *= space[i]
		res %= MOD
	}
	return res
}