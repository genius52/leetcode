package number

//如果玩家移除石子后，导致 所有已移除石子 的价值 总和 可以被 3 整除，那么该玩家就 输掉游戏 。
//如果不满足上一条，且移除后没有任何剩余的石子，那么 Bob 将会直接获胜（即便是在 Alice 的回合）。
func check_StoneGameIX(mod0_cnt,mod1_cnt,mod2_cnt int,sum int)bool{
	if mod1_cnt < 0 || mod2_cnt < 0{
		return false
	}
	var reverse bool = false
	if mod0_cnt | 1 == mod0_cnt{
		reverse = true
	}
	var is_alice bool = false
	for mod1_cnt > 0 || mod2_cnt > 0{
		if sum == 1{
			if mod1_cnt > 0{
				mod1_cnt--
			}else{
				if reverse{
					return is_alice
				}else{
					return !is_alice
				}
			}
			sum = 2
		}else if sum == 2{
			if mod2_cnt > 0{
				mod2_cnt--
			}else{
				if reverse{
					return is_alice
				}else{
					return !is_alice
				}
			}
			sum = 1
		}
		is_alice = !is_alice
	}
	return false
}

func StoneGameIX(stones []int) bool {
	var mod0_cnt,mod1_cnt,mod2_cnt int = 0,0,0
	for _,n := range stones{
		mod := n % 3
		if mod == 0{
			mod0_cnt++
		}else if mod == 1{
			mod1_cnt++
		}else{
			mod2_cnt++
		}
	}
	//1,1,2,1,2,1
	//2,2,1,2,1,2
	return check_StoneGameIX(mod0_cnt,mod1_cnt - 1,mod2_cnt,1) ||
		check_StoneGameIX(mod0_cnt,mod1_cnt,mod2_cnt - 1,2)
}