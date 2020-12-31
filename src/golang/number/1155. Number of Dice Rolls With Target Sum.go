package number

//Input: d = 2, f = 6, target = 7
//Output: 6
//Explanation:
//You throw two dice, each with 6 faces.  There are 6 ways to get a sum of 7:
//1+6, 2+5, 3+4, 4+3, +52, 6+1.
func dp_numRollsToTarget(d int,f int,target int,memo map[string]int)int{
	if d == 0 {
		if target == 0{
			return 1
		}
		return 0
	}
	if d * f < target || d > target || target < 0{
		return 0
	}
	if d == 1{
		return 1
	}
	key := string(d) + "," + string(target)
	if n,ok := memo[key];ok{
		return n
	}
	var total int = 0
	for i := 1;i <= f;i++{
		if i > target{
			break
		}
		total += dp_numRollsToTarget(d - 1,f,target - i,memo)
		total %= 1000000007
	}
	memo[key] = total
	return total
}

func NumRollsToTarget(d int, f int, target int) int {
	var memo map[string]int = make(map[string]int)
	return dp_numRollsToTarget(d,f,target,memo)
}