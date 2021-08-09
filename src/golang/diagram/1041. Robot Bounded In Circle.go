package diagram

//"G": go straight 1 unit;
//"L": turn 90 degrees to the left;
//"R": turn 90 degress to the right.
//Input: "GL"
//Output: true
//Explanation:
//The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...
func IsRobotBounded(instructions string) bool{
	var l int = len(instructions)
	var dirs [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	var idx int = 0
	x := 0
	y := 0
	for i := 0;i < l;i++{
		if instructions[i] == 'G'{
			x += dirs[idx][0]
			y += dirs[idx][1]
		}else if instructions[i] == 'L'{
			idx = (idx + 4 - 1) % 4
		}else if instructions[i] == 'R'{
			idx = (idx + 1) % 4
		}
	}
	if x == 0 && y == 0{
		return true
	}
	return idx != 0
}
//func IsRobotBounded(instructions string) bool {
//	var l int = len(instructions)
//	var cur_x int = 0
//	var cur_y int = 0
//	var dir []int = []int{0,1}
//	for i := 0;i < l;i++{
//		if instructions[i] == 'G'{
//			cur_x += dir[0]
//			cur_y += dir[1]
//		}else if instructions[i] == 'L'{
//			if dir[0] == 0 && dir[1] == 1{
//				dir[0] = -1
//				dir[1] = 0
//			}else if dir[0] == -1 && dir[1] == 0{
//				dir[0] = 0
//				dir[1] = -1
//			}else if dir[0] == 0 && dir[1] == -1{
//				dir[0] = 1
//				dir[1] = 0
//			}else if dir[0] == 1 && dir[1] == 0{
//				dir[0] = 0
//				dir[1] = 1
//			}
//		}else if instructions[i] == 'R'{
//			if dir[0] == 0 && dir[1] == 1{
//				dir[0] = 1
//				dir[1] = 0
//			}else if dir[0] == 1 && dir[1] == 0{
//				dir[0] = 0
//				dir[1] = -1
//			}else if dir[0] == 0 && dir[1] == -1{
//				dir[0] = -1
//				dir[1] = 0
//			}else if dir[0] == -1 && dir[1] == 0{
//				dir[0] = 0
//				dir[1] = 1
//			}
//		}
//	}
//	if cur_x == 0 && cur_y == 0{
//		return true
//	}
//	for round := 0;round < 3;round++{
//		for i := 0;i < l;i++{
//			if instructions[i] == 'G'{
//				cur_x += dir[0]
//				cur_y += dir[1]
//			}else if instructions[i] == 'L'{
//				if dir[0] == 0 && dir[1] == 1{
//					dir[0] = -1
//					dir[1] = 0
//				}else if dir[0] == -1 && dir[1] == 0{
//					dir[0] = 0
//					dir[1] = -1
//				}else if dir[0] == 0 && dir[1] == -1{
//					dir[0] = 1
//					dir[1] = 0
//				}else if dir[0] == 1 && dir[1] == 0{
//					dir[0] = 0
//					dir[1] = 1
//				}
//			}else if instructions[i] == 'R'{
//				if dir[0] == 0 && dir[1] == 1{
//					dir[0] = 1
//					dir[1] = 0
//				}else if dir[0] == 1 && dir[1] == 0{
//					dir[0] = 0
//					dir[1] = -1
//				}else if dir[0] == 0 && dir[1] == -1{
//					dir[0] = -1
//					dir[1] = 0
//				}else if dir[0] == -1 && dir[1] == 0{
//					dir[0] = 0
//					dir[1] = 1
//				}
//			}
//		}
//	}
//	if cur_x == 0 && cur_y == 0{
//		return true
//	}
//	return false
//}