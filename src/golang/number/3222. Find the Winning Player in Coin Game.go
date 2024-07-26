package number

func losingPlayer(x int, y int) string {
	var is_alice bool = false
	for x >= 1 && y >= 4 {
		x--
		y -= 4
		is_alice = !is_alice
	}
	if is_alice {
		return "Alice"
	} else {
		return "Bob"
	}
}
