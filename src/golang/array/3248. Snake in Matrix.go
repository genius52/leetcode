package array

func finalPositionOfSnake(n int, commands []string) int {
	var res int = 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == "UP" {
			res -= n
		} else if commands[i] == "RIGHT" {
			res++
		} else if commands[i] == "DOWN" {
			res += n
		} else if commands[i] == "LEFT" {
			res--
		}
	}
	return res
}
