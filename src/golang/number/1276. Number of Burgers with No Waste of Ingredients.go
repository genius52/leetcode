package number

//Jumbo Burger: 4 tomato slices and 1 cheese slice.
//Small Burger: 2 Tomato slices and 1 cheese slice.
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	//4 * x + 2 * y = tomatoSlices
	//2x + 2y = cheeseSlices * 2
	//3x + y = tomatoSlices - cheeseSlices
	var res []int
	diff := (tomatoSlices - cheeseSlices * 2)
	if diff % 2 != 0{
		return res
	}
	x := diff /2
	y := cheeseSlices - x
	if x < 0 || y < 0{
		return res
	}
	res = make([]int,2)
	res[0] = x
	res[1] = y
	return res
}