package diagram

func DistanceBetweenBusStops(distance []int, start int, destination int) int {
	var l int = len(distance)
	var cc int = 0
	var cw int = 0
	var visit int = start
	for visit != destination{
		cc += distance[visit]
		visit = (visit + 1) % l
	}
	visit = start
	for visit != destination{
		visit = (visit - 1 + l) % l
		cw += distance[visit]
	}
	if cc < cw{
		return cc
	}else{
		return cw
	}
}