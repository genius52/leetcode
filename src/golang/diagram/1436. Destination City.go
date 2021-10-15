package diagram

func destCity2(paths [][]string) string {
	var dst_city map[string]bool  = make(map[string]bool)
	for _,path := range paths{
		dst_city[path[1]] = true
	}
	for _,path := range paths{
		if _,ok := dst_city[path[0]];ok{
			delete(dst_city,path[0])
		}
	}
	var res string
	for dst,_ := range dst_city{
		res = dst
	}
	return res
}

func destCity(paths [][]string) string {
	var start_city map[string]bool = make(map[string]bool)
	var dst_city map[string]bool  = make(map[string]bool)
	for _,path := range paths{
		start_city[path[0]] = true
		dst_city[path[1]] = true
	}
	for start,_ := range start_city{
		if _,ok := dst_city[start];ok{
			delete(dst_city,start)
		}
	}
	var res string
	for dst,_ := range dst_city{
		res = dst
	}
	return res
}