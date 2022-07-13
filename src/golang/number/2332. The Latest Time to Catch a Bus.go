package number

import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	var bus_len int = len(buses)
	var pass_len int = len(passengers)
	var res int = 0
	var pass_cnt int = 0
	var pass_idx int = 0
	var bus_idx int = 0
	for pass_idx < pass_len && bus_idx < bus_len{
		if passengers[pass_idx] <= buses[bus_idx]{
			pass_idx++
			pass_cnt++
			if pass_cnt == capacity{
				bus_idx++
				pass_cnt = 0
			}
		}else{
			bus_idx++
			pass_cnt = 0
		}
	}
	return res
}