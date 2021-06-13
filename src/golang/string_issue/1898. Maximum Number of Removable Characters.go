package string_issue

func check_maximumRemovals(s string,p string,forbbiden map[int]bool)bool{
	var len_s int = len(s)
	var len_p int = len(p)
	var i int = 0
	var j int = 0
	for i < len_s && j < len_p{
		for i < len_s {
			if _,ok := forbbiden[i];ok{
				i++
				continue
			}
			if s[i] == p[j]{
				break
			}
			i++
		}
		if i == len_s{
			return false
		}
		i++
		j++
	}
	if j == len_p{
		return true
	}
	return false
}

func MaximumRemovals(s string, p string, removable []int) int {
	var len_k int = len(removable)
	var low int = 0
	var high int = len_k
	var res int = 0
	for low <= high{
		mid := (low + high)/2
		removable2 := removable[:min_int(len_k,mid + 1)]
		var forbbiden map[int]bool = make(map[int]bool)
		for _,idx := range removable2{
			forbbiden[idx] = true
		}
		ret := check_maximumRemovals(s,p,forbbiden)
		if ret{
			if mid == len_k{
				res = mid
				break
			}
			res = mid + 1
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return res
}