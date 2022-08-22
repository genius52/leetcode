package number

func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var l int = len(energy)
	var cur_exp int = initialExperience
	var total_energy int = 0
	var add_exp int = 0
	for i := 0;i < l;i++{
		total_energy += energy[i]
		if experience[i] >= cur_exp{
			add_exp += experience[i] - cur_exp + 1
			cur_exp += add_exp + experience[i]
		}else{
			cur_exp += experience[i]
		}
	}
	var res int = 0
	if total_energy >= initialEnergy{
		res += total_energy - initialEnergy + 1
	}
	return res + add_exp
}