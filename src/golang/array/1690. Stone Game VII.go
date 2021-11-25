package array

//输入：stones = [5,3,1,4,2]
//输出：6
//解释：
//- 爱丽丝移除 2 ，得分 5 + 3 + 1 + 4 = 13 。游戏情况：爱丽丝 = 13 ，鲍勃 = 0 ，石子 = [5,3,1,4] 。
//- 鲍勃移除 5 ，得分 3 + 1 + 4 = 8 。游戏情况：爱丽丝 = 13 ，鲍勃 = 8 ，石子 = [3,1,4] 。
//- 爱丽丝移除 3 ，得分 1 + 4 = 5 。游戏情况：爱丽丝 = 18 ，鲍勃 = 8 ，石子 = [1,4] 。
//- 鲍勃移除 1 ，得分 4 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [4] 。
//- 爱丽丝移除 4 ，得分 0 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [] 。
//得分的差值 18 - 12 = 6 。
func dp_stoneGameVII(stones []int,l int,left int,right int,alice int,prefix []int,memo [][]int)int{
	if left == right{
		return alice
	}
	if memo[left][right] != 0{
		return memo[left][right]
	}
	//删除最左边后，本次的值 + 下次的值
	if alice != 0{
		del_left_diff := alice - (prefix[right + 1] - prefix[left + 1])
		next1 := dp_stoneGameVII(stones,l,left + 1,right,0,prefix,memo)
		del_right_diff :=alice - (prefix[right] - prefix[left])
		next2 := dp_stoneGameVII(stones,l,left,right - 1,0,prefix,memo)
		memo[left][right] = min_int(del_left_diff + next1,del_right_diff + next2)
		return memo[left][right]
	}else{
		next1 := dp_stoneGameVII(stones,l,left + 1,right,prefix[right + 1] - prefix[left + 1],prefix,memo)
		next2 := dp_stoneGameVII(stones,l,left,right - 1,prefix[right] - prefix[left],prefix,memo)
		memo[left][right] = max_int(next1,next2)
		return memo[left][right]
	}
}

func StoneGameVII(stones []int) int {
	var l int = len(stones)
	var prefix []int = make([]int,l + 1)
	for i := 0;i < l;i++{
		prefix[i + 1] = stones[i] + prefix[i]
	}
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,l)
	}
	return dp_stoneGameVII(stones,l,0,l - 1,0,prefix,memo)
}