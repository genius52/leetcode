package number

func dfs_maximumBobPoints(totalArrows int,numArrows int,aliceArrows []int,idx int,cur_status int,max_score *int,res *[]int){
	if idx < 0 || numArrows == 0{
		score := 0
		for i := 0;i <= 11;i++{
			if (cur_status | (1 << i)) == cur_status{
				score += i
			}
		}
		if score > *max_score{
			*max_score = score
			for i := 11;i >= 1;i--{
				if (cur_status | (1 << i)) == cur_status{
					if aliceArrows[i] == 0{
						(*res)[i] = 1
					}else{
						(*res)[i] = aliceArrows[i] + 1
					}
					totalArrows -= (*res)[i]
				}else{
					(*res)[i] = 0
				}
			}
			(*res)[0] = totalArrows
		}
		return
	}
	if aliceArrows[idx] == 0{
		dfs_maximumBobPoints(totalArrows,numArrows - 1,aliceArrows,idx - 1,cur_status | (1 << idx),max_score,res)
	}else{
		if numArrows > aliceArrows[idx]{
			dfs_maximumBobPoints(totalArrows,numArrows - aliceArrows[idx] - 1,aliceArrows,idx - 1,cur_status | (1 << idx),max_score,res)
		}
		dfs_maximumBobPoints(totalArrows,numArrows,aliceArrows,idx - 1,cur_status,max_score,res)
	}
}

func MaximumBobPoints(numArrows int, aliceArrows []int) []int {
	var max_score int = 0
	var res []int = make([]int,12)
	dfs_maximumBobPoints(numArrows,numArrows,aliceArrows,11,0,&max_score,&res)
	return res
}