package array

func dp_maxEnergyBoost(energyDrinkA []int, energyDrinkB []int, idx int, is_a bool, memo [][2]int64) int64 {
	if idx >= len(energyDrinkA) {
		return 0
	}
	var choose int = 0
	if is_a {
		if memo[idx][choose] != 0 {
			return memo[idx][choose]
		}
		memo[idx][choose] = max_int64(int64(energyDrinkA[idx])+dp_maxEnergyBoost(energyDrinkA, energyDrinkB, idx+1, true, memo),
			dp_maxEnergyBoost(energyDrinkA, energyDrinkB, idx+1, false, memo))
	} else {
		choose = 1
		if memo[idx][choose] != 0 {
			return memo[idx][choose]
		}
		memo[idx][choose] = max_int64(int64(energyDrinkB[idx])+dp_maxEnergyBoost(energyDrinkA, energyDrinkB, idx+1, false, memo),
			dp_maxEnergyBoost(energyDrinkA, energyDrinkB, idx+1, true, memo))
	}
	return memo[idx][choose]
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	var l int = len(energyDrinkA)
	var memo [][2]int64 = make([][2]int64, l)
	return max(dp_maxEnergyBoost(energyDrinkA, energyDrinkB, 0, true, memo),
		dp_maxEnergyBoost(energyDrinkA, energyDrinkB, 0, false, memo))
}
