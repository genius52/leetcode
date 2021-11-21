package array

func wateringPlants(plants []int, capacity int) int {
	var l int = len(plants)
	var steps int = 0
	var cur_cap int = capacity
	for i := 0;i < l;i++{
		if cur_cap < plants[i] {
			steps += i * 2
			cur_cap = capacity - plants[i]
		}else{
			cur_cap -= plants[i]
		}
		steps++
	}
	return steps
}