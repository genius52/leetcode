package diagram

//Input: hats = [[3,5,1],[3,5]]
//Output: 4
//Explanation: There are 4 ways to choose hats
//(3,5), (5,3), (1,3) and (1,5)
//
func dfs_numberWays(hat_people [41][]int,cur_hat int,mask int,target int,memo *[41][1024]int)int{
	if mask == target{
		return 1
	}
	if cur_hat >= 41{
		return 0
	}
	if (*memo)[cur_hat][mask] != -1{
		return (*memo)[cur_hat][mask]
	}
	var res int = dfs_numberWays(hat_people,cur_hat + 1,mask,target,memo)
	for i := 0;i < len(hat_people[cur_hat]);i++{
		people := hat_people[cur_hat][i]
		nextmask := mask | (1 << people)
		if nextmask == mask{
			continue
		}
		res += dfs_numberWays(hat_people,cur_hat + 1,nextmask,target,memo)
		res = res % (1e9 + 7)
	}
	(*memo)[cur_hat][mask] = res
	return res
}

func NumberWays(hats [][]int) int{
	var hat_people [41][]int
	for people,hat := range hats{
		for _,h := range hat{
			hat_people[h] = append(hat_people[h],people)
		}
	}
	var memo [41][1024]int
	for i := 0;i < 41;i++{
		for j := 0;j < 1024;j++{
			memo[i][j] = -1
		}
	}
	var l int = len(hats)
	var target int = 1 << l - 1
	//mask: 1 << i 第i个人已经戴了帽子
	return dfs_numberWays(hat_people,1,0,target,&memo)
}

//TLE
//func dp_numberWays(hats [][]int,l int,pos int,mask int64,memo *[][]int)int{
//	if pos >= l{
//		return 1
//	}
//	if (*memo)[pos][mask] != 0{
//		return (*memo)[pos][mask]
//	}
//	var res int = 0
//	for i := 0;i < len(hats[pos]);i++{
//		next_mask := mask | (1 << hats[pos][i])
//		if next_mask  == mask{
//			continue
//		}
//		res += dp_numberWays(hats,l,pos + 1,next_mask,memo)
//		res = res % (1e9 + 7)
//	}
//	(*memo)[pos][mask] = res
//	return res
//}
//
//func numberWays(hats [][]int) int {
//	var l int = len(hats)
//	//var memo map[int64]int = make(map[int64]int)
//	var memo [][]int = make([][]int,l)
//	for i := 0;i < l;i++{
//		memo[i] = make([]int,1 << 41)
//	}
//	return dp_numberWays(hats,l,0,0,&memo)
//}