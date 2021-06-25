package diagram

func robotSim(commands []int, obstacles [][]int) int {
	var block map[int]bool = make(map[int]bool)
	for _,o := range obstacles{
		block[100000 * o[0] + o[1]] = true
	}
	//-2: turn left 90 degrees,
	//-1: turn right 90 degrees, or
	//1 <= k <= 9: move forward k units.
	var dirs [4][]int = [4][]int{{0,1},{1,0},{0,-1},{-1,0}};
	cur_dir := 0
	x := 0
	y := 0
	var res int = 0
	for _,c := range commands{
		if c == -1{
			cur_dir = (cur_dir + 1) % 4
		}else if c == -2{
			cur_dir = (cur_dir + 4 - 1) % 4
		}else{
			for c > 0{
				x += dirs[cur_dir][0]
				y += dirs[cur_dir][1]
				if _,ok := block[100000 * x + y];ok{
					x -= dirs[cur_dir][0]
					y -= dirs[cur_dir][1]
					break
				}
				c--
			}
			cur :=  x * x + y * y
			if cur > res{
				res = cur
			}
		}
	}
	return res
}