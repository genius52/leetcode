package number

func garbageCollection(garbage []string, travel []int) int {
	var last_p int = -1
	var last_g int = -1
	var last_m int = -1
	var l int = len(garbage)
	var p_cnt []int = make([]int,l)
	var g_cnt []int = make([]int,l)
	var m_cnt []int = make([]int,l)
	for i := 0;i < l;i++{
		for _,c := range garbage[i]{
			if c == 'P'{
				p_cnt[i]++
				last_p = i
			}else if c == 'G'{
				g_cnt[i]++
				last_g = i
			}else if c == 'M'{
				m_cnt[i]++
				last_m = i
			}
		}
	}
	var res int = 0
	for i := 0;i <= last_p;i++{
		res += p_cnt[i]
		if i > 0{
			res += travel[i - 1]
		}
	}
	for i := 0;i <= last_g;i++{
		res += g_cnt[i]
		if i > 0{
			res += travel[i - 1]
		}
	}
	for i := 0;i <= last_m;i++{
		res += m_cnt[i]
		if i > 0{
			res += travel[i - 1]
		}
	}
	return res
}